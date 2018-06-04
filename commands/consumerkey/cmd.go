package consumerkey

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "consumer-key",
		Short: "Manage consumer keys",
	}

	cmd.AddCommand(
		NewCreateCommand(),
	)

	return cmd
}
