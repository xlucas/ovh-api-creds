package consumerkey

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"

	"github.com/ovh/go-ovh/ovh"
	"github.com/spf13/cobra"
	"github.com/xlucas/ovh-api-creds/endpoint"
)

type createOpts struct {
	applicationKey    string
	applicationSecret string
	path              string
	readOnly          bool
	service           string
	zone              string
}

func NewCreateCommand() *cobra.Command {
	opt := new(createOpts)
	cmd := &cobra.Command{
		Use:          "create",
		Short:        "Create a consumer key",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(
				opt.path,
				opt.applicationKey,
				opt.applicationSecret,
				opt.service,
				opt.zone,
				opt.readOnly,
			)
		},
	}

	pFlags := cmd.PersistentFlags()
	pFlags.StringVar(&opt.applicationKey, "app-key", "", "The application key")
	pFlags.StringVar(&opt.applicationSecret, "app-secret", "", "The application secret")
	pFlags.StringVar(&opt.path, "path", "/*", "Only allow this path to be reached with the consumer key")
	pFlags.BoolVar(&opt.readOnly, "ro", false, "Only allow GET request to be made with the consumer key")
	pFlags.StringVar(&opt.service, "service", "ovh", "The service to create a consumer key for")
	pFlags.StringVar(&opt.zone, "zone", "eu", "The zone where the consumer key will be used")

	cmd.MarkPersistentFlagRequired("app-key")
	cmd.MarkPersistentFlagRequired("app-secret")

	return cmd
}

func runCreate(path, appKey, appSecret, service, zone string, readOnly bool) error {
	apiURL, err := endpoint.Get(service, zone)
	if err != nil {
		return err
	}

	client, err := ovh.NewClient(apiURL, appKey, appSecret, "")
	if err != nil {
		return err
	}

	req := client.NewCkRequest()
	if readOnly {
		req.AddRules(ovh.ReadOnly, path)
	} else {
		req.AddRules(ovh.ReadWrite, path)
	}

	response, err := req.Do()
	if err != nil {
		return err
	}

	fmt.Println("Consumer key:", response.ConsumerKey)

	return open.Run(response.ValidationURL)
}
