package main

import (
	"os"

	"github.com/xlucas/ovh-api-creds/commands"
)

func main() {
	if err := commands.Root.Execute(); err != nil {
		os.Exit(1)
	}
}
