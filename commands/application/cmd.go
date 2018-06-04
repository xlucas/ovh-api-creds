package application

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "application",
		Short: "Manage applications",
	}

	cmd.AddCommand(NewCreateCommand())

	return cmd
}
