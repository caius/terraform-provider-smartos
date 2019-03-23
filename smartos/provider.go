package smartos

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The SmartOS host to connect to",
			},

			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user to connect to the host as",
			},

			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The port to connect to the host on",
			},
		},

		// DataSourcesMap: map[string]*schema.Resource{
		//   "smartos_image": dataSourceSmartosImage(),
		// },

		ResourcesMap: map[string]*schema.Resource{
			"smartos_image": resourceSmartosImage(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	wankbucket("providerConfigure called")

	config := Config{
		Host: d.Get("host").(string),
		User: d.Get("user").(string),
		Port: d.Get("port").(int),
	}

	wankbucket(fmt.Sprintf("Configured provider with config %+v", config))

	return config.Client()
}

func wankbucket(message string) {
	log.Printf("WANKBUCKET %s", message)
}
