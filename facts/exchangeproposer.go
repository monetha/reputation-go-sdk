package facts

import (
	"context"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/crypto/ecies"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/pkg/errors"
)

const proposeGasLimit = 500000

// ExchangeProposer allows to propose data exchange
type ExchangeProposer eth.Session

// NewExchangeProposer converts session to ExchangeProposer
func NewExchangeProposer(s *eth.Session) *ExchangeProposer {
	return (*ExchangeProposer)(s)
}

// ProposePrivateDataExchangeResult holds result of calling ProposePrivateDataExchange
type ProposePrivateDataExchangeResult struct {
	ExchangeIdx     *big.Int    // exchange index
	ExchangeKey     [32]byte    // exchange key is a key to be XORed with the secret key of private data (should be kept in secret by the data requester)
	ExchangeKeyHash common.Hash // hash of exchange key hash
}

// ProposePrivateDataExchange creates private data exchange proposition
func (e *ExchangeProposer) ProposePrivateDataExchange(
	ctx context.Context,
	passportAddress common.Address,
	factProviderAddress common.Address,
	factKey [32]byte,
	exchangeStakeWei *big.Int,
	rand io.Reader,
) (*ProposePrivateDataExchangeResult, error) {
	noWaitResult, err := e.ProposePrivateDataExchangeNoWait(ctx, passportAddress, factProviderAddress, factKey, exchangeStakeWei, rand)
	if err != nil {
		return nil, err
	}

	return e.GetProposePrivateDataExchangeResult(ctx, noWaitResult)
}

// ProposePrivateDataExchangeResultNoWait holds result of calling ProposePrivateDataExchangeNoWait
type ProposePrivateDataExchangeResultNoWait struct {
	ExchangeKey     [32]byte
	ExchangeKeyHash common.Hash
	TxHash          common.Hash
}

// ProposePrivateDataExchangeNoWait creates private data exchange proposition.
// This method does not wait for the transaction to be mined. Use the method without the NoWait suffix if you need to make
// sure that the transaction was successfully mined.
func (e *ExchangeProposer) ProposePrivateDataExchangeNoWait(
	ctx context.Context,
	passportAddress common.Address,
	factProviderAddress common.Address,
	factKey [32]byte,
	exchangeStakeWei *big.Int,
	rand io.Reader,
) (*ProposePrivateDataExchangeResultNoWait, error) {
	backend := e.Backend
	auth := e.TransactOpts
	auth.Context = ctx
	auth.Value = exchangeStakeWei
	auth.GasLimit = proposeGasLimit

	// reading passport owner key
	passportOwnerPublicKey, err := NewReader(e.Eth).ReadOwnerPublicKey(ctx, passportAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read passport owner public key")
	}

	// creating exchange key shared with passport owner
	ec, err := ecies.NewGenerate(passportOwnerPublicKey.Curve, rand)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ECIES instance")
	}

	_, exchangeKey, exchangeKeyHash, err := deriveSecretKeyringMaterial(ec, passportOwnerPublicKey, passportAddress, factProviderAddress, factKey)
	if err != nil {
		return nil, err
	}
	encryptedExchangeKey := ecies.MarshalPublicKey(ec.PublicKey())

	// propose private data exchange
	c := contracts.InitPassportLogicContract(passportAddress, backend)

	e.Log("Proposing private data exchange", "fact_provider", factProviderAddress, "fact_key", factKey, "encrypted_key", encryptedExchangeKey, "key_hash", exchangeKeyHash)
	tx, err := c.ProposePrivateDataExchange(&auth, factProviderAddress, factKey, encryptedExchangeKey, exchangeKeyHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to propose private data exchange")
	}

	res := &ProposePrivateDataExchangeResultNoWait{
		ExchangeKeyHash: exchangeKeyHash,
		TxHash:          tx.Hash(),
	}
	copy(res.ExchangeKey[:], exchangeKey)
	return res, nil
}

// GetProposePrivateDataExchangeResult waits for transaction to be mined if it's not mined yet and retrieves proposed private data exchange result.
func (e *ExchangeProposer) GetProposePrivateDataExchangeResult(ctx context.Context, noWaitResult *ProposePrivateDataExchangeResultNoWait) (*ProposePrivateDataExchangeResult, error) {
	auth := e.TransactOpts

	txr, err := e.WaitForTxReceipt(ctx, noWaitResult.TxHash)
	if err != nil {
		return nil, err
	}

	// filtering PrivateDataExchangeProposed event to get exchangeIdx
	evs, err := contracts.FilterPrivateDataExchangeProposed(ctx, txr.Logs, nil, []common.Address{auth.From}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "filtering PrivateDataExchangeProposed event")
	}
	if len(evs) != 1 {
		return nil, errors.New("expected exactly one PrivateDataExchangeProposed event")
	}
	ev := evs[0]
	exchangeIdx := ev.ExchangeIdx

	log.Warn("PrivateDataExchangeProposed", "exchange_index", ev.ExchangeIdx.String(), "data_requester", ev.DataRequester.String(), "passport_owner", ev.PassportOwner.String())

	return &ProposePrivateDataExchangeResult{
		ExchangeIdx:     exchangeIdx,
		ExchangeKey:     noWaitResult.ExchangeKey,
		ExchangeKeyHash: noWaitResult.ExchangeKeyHash,
	}, nil
}
