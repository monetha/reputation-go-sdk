package quorum

import (
	"context"
	"crypto/ecdsa"
	"net"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/monetha/go-verifiable-data/eth/backend/ethclient"
)

// PrivateTxClient is a client to work with Quorum's private transactions
type PrivateTxClient struct {
	*ethclient.Client
	c             *rpc.Client
	hc            *http.Client
	privateFor    []string
	quorumEnclave string
}

type privateForParams struct {
	PrivateFor []string `json:"privateFor"`
}

// DialHTTP connects a client to the given URL using provided context.
func DialHTTP(rawurl string, privateFor []string, quorumEnclave string) (*PrivateTxClient, error) {
	hc := &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}}

	return DialHTTPWithClient(rawurl, privateFor, quorumEnclave, hc)
}

// DialHTTPWithClient connects a client to the given URL using provided context and uses given HTTP client for
// internal http requests.
func DialHTTPWithClient(rawurl string, privateFor []string, quorumEnclave string, hc *http.Client) (*PrivateTxClient, error) {
	c, err := rpc.DialHTTP(rawurl)
	if err != nil {
		return nil, err
	}

	return NewClient(c, hc, privateFor, quorumEnclave), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, hc *http.Client, privateFor []string, quorumEnclave string) *PrivateTxClient {
	return &PrivateTxClient{
		c:             c,
		hc:            hc,
		Client:        ethclient.NewClient(c),
		privateFor:    privateFor,
		quorumEnclave: quorumEnclave,
	}
}

// SendTransaction injects a signed private transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec *PrivateTxClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}

	privateFor := &privateForParams{
		PrivateFor: ec.privateFor,
	}

	return ec.c.CallContext(ctx, nil, "eth_sendRawPrivateTransaction", hexutil.Encode(data), privateFor)
}

// GetSenderPublicKey retrieves public key of sender from transaction.
func (ec *PrivateTxClient) GetSenderPublicKey(t *types.Transaction) (*ecdsa.PublicKey, error) {
	return ec.Client.GetSenderPublicKey(RestoreTxToSignedForm(t))
}

// NewKeyedTransactor is a utility method to easily create a transaction signer
// from a single private key.
func (ec *PrivateTxClient) NewKeyedTransactor(key *ecdsa.PrivateKey) *bind.TransactOpts {
	return NewPrivateTransactor(ec.hc, key, ec.quorumEnclave)
}

// DecryptPayload decrypts transaction payload.
func (ec *PrivateTxClient) DecryptPayload(ctx context.Context, tx *types.Transaction) (bs []byte, err error) {
	return DecryptPayload(ctx, tx, ec.c)
}
