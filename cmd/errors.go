package cmd

import "errors"

// Errors emitted by the command framework
var (
	ErrSubcommandRequired = errors.New("Command can not be run directly, refer to Available Commands")
)
