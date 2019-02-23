package smartos

import (
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
