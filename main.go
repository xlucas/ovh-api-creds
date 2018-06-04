package main

import (
	"fmt"
	"os"

	"github.com/xlucas/ovh-api-creds/commands"
)

func main() {
	if err := commands.Root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
