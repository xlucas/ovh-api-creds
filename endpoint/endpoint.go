package endpoint

import (
	"fmt"
	"strings"

	"github.com/ovh/go-ovh/ovh"
)

func Get(service, zone string) (string, error) {
	endpointKey := fmt.Sprintf("%s-%s", strings.ToLower(service), strings.ToLower(zone))
	endpoint, ok := ovh.Endpoints[endpointKey]
	if !ok {
		return "", fmt.Errorf("invalid combination of service and zone")
	}
	return endpoint, nil
}
