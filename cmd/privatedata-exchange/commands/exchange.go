package commands

import (
	"os"

	"github.com/monetha/go-verifiable-data/cmd"
)

// ExchangeCommand is a root command used to parse command line arguments
type ExchangeCommand struct {
	Version func()         `short:"v" long:"version" description:"Print the version of tool and exit"`
	Propose ProposeCommand `command:"propose"        description:"Propose private data exchange"`
	Status  StatusCommand  `command:"status"         description:"Get the status of private data exchange"`
	Accept  AcceptCommand  `command:"accept"         description:"Accept private data exchange"`
	Finish  FinishCommand  `command:"finish"         description:"Finish private data exchange"`
	Timeout TimeoutCommand `command:"timeout"        description:"Close private data exchange after proposal timeout"`
	Dispute DisputeCommand `command:"dispute"        description:"Open private data exchange dispute"`
	Read    ReadCommand    `command:"read"           description:"Reads private data after private data exchange accepted"`
}

// Exchange is a prepared command to be used in command line arguments parsing
var Exchange ExchangeCommand

func init() {
	Exchange.Version = func() {
		cmd.PrintVersion()
		os.Exit(0)
	}
}
