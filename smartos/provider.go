package smartos

import (
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
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
