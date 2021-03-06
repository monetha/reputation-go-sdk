package commands

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/internal/cmdutils"
	"github.com/monetha/go-verifiable-data/cmd/internal/flag"
	"github.com/monetha/go-verifiable-data/facts"
)

// DisputeCommand handles dispute command
type DisputeCommand struct {
	cmdutils.QuorumPrivateTxIOCommand
	PassportAddress  flag.EthereumAddress         `long:"passportaddr" required:"true" description:"Ethereum address of passport contract"`
	ExchangeIndex    flag.ExchangeIndex           `long:"exchidx"      required:"true" description:"private data exchange index"`
	ExchangeKey      flag.ExchangeKeyFromFile     `long:"exchangekey"  required:"true" description:"exchange key filename"`
	DataRequesterKey flag.ECDSAPrivateKeyFromFile `long:"requesterkey" required:"true" description:"data requester private key filename"`
	BackendURL       string                       `long:"backendurl"   required:"true" description:"Ethereum backend URL"`
	Verbosity        int                          `long:"verbosity"                    description:"log verbosity (0-9)" default:"2"`
	VModule          string                       `long:"vmodule"                      description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *DisputeCommand) Execute(args []string) error {
	cmdutils.InitLogging(log.Lvl(c.Verbosity), c.VModule)
	ctx := cmdutils.CreateCtrlCContext()

	e, err := cmdutils.NewEth(ctx, c.BackendURL, c.QuorumEnclave, c.QuorumPrivateFor.AsStringArr())
	if err != nil {
		return err
	}

	dispRes, err := facts.NewExchangeDisputer(
		e.NewSession(c.DataRequesterKey.AsECDSAPrivateKey()),
		nil,
	).DisputePrivateDataExchange(
		ctx,
		c.PassportAddress.AsCommonAddress(),
		c.ExchangeIndex.AsBigInt(),
		c.ExchangeKey,
	)
	if err != nil {
		return err
	}

	log.Warn("Dispute result", "successful", dispRes.Successful, "cheater_address", dispRes.Cheater.String())

	return nil
}
