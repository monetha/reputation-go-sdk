package ethclient

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/monetha/go-verifiable-data/eth/backend"
)

// Client defines typed wrappers for the Ethereum RPC API, it also implements additional methods like GetSenderPublicKey.
type Client struct {
	*ethclient.Client
}

// GetSenderPublicKey retrieves public key of sender from transaction.
func (c *Client) GetSenderPublicKey(t *types.Transaction) (*ecdsa.PublicKey, error) {
	return (*backend.Transaction)(t).GetSenderPublicKey()
}

// NewKeyedTransactor is a utility method to easily create a transaction signer
// from a single private key.
func (c *Client) NewKeyedTransactor(key *ecdsa.PrivateKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key)
}

// DecryptPayload decrypts transaction payload.
func (c *Client) DecryptPayload(ctx context.Context, tx *types.Transaction) ([]byte, error) {
	return tx.Data(), nil
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

// DialContext connects a client to the given URL using provided context.
func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := ethclient.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return &Client{Client: c}, nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *Client {
	return &Client{Client: ethclient.NewClient(c)}
}
