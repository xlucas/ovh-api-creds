package commands

import (
	"github.com/spf13/cobra"
	"github.com/xlucas/ovh-api-creds/commands/application"
	"github.com/xlucas/ovh-api-creds/commands/consumerkey"
)

var Root = &cobra.Command{
	Short: "A helper for OVH API credentials",
}

func init() {
	Root.AddCommand(
		application.NewCommand(),
		consumerkey.NewCommand(),
	)
}
