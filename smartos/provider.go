package smartos

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		// DataSourcesMap: map[string]*schema.Resource{
		//   "smartos_image": dataImage(),
		// },
		// ResourcesMap: map[string]*schema.Resource{
		//   "smartos_zone": resourceZone(),
		// },
	}
}

type Config struct {
	Host string
}

func (c Config) validate() error {
	if c.Host == "" {
		return errors.New("Host must be configured for the SmartOS provider")
	}

	return nil
}

func (c Config) newClient() (*Client, error) {
	return &Client{
		Host: c.Host,
	}, nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Host: d.Get("host").(string),
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	client, err := config.newClient()
	if err != nil {
		return nil, err
	}

	return client, nil
}
