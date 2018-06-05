package application

import (
	"net/url"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/xlucas/ovh-api-creds/endpoint"
)

type createOpts struct {
	aio     bool
	service string
	zone    string
}

func NewCreateCommand() *cobra.Command {
	opt := new(createOpts)
	cmd := &cobra.Command{
		Use:          "create",
		Short:        "Create an application",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(opt.aio, opt.service, opt.zone)
		},
	}

	pflags := cmd.PersistentFlags()
	pflags.BoolVar(&opt.aio, "all-in-one", false, "Create all keys in a single step")
	pflags.StringVar(&opt.service, "service", "ovh", "The service to create credentials for")
	pflags.StringVar(&opt.zone, "zone", "eu", "The zone where credentials will be used")

	return cmd
}

func runCreate(aio bool, service, zone string) error {
	apiURL, err := endpoint.Get(service, zone)
	if err != nil {
		return err
	}

	endpointURL, _ := url.Parse(apiURL)
	if aio {
		endpointURL.Path = "/createToken"
	} else {
		endpointURL.Path = "/createApp"
	}

	return open.Run(endpointURL.String())
}
