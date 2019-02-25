package smartos

import (
	"github.com/caius/goadm/imgadm"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceImageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceImageRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*Client)
	var foo imgadm.Client

	// client.Imgadm.GetByName(d.Get("name"))

	d.Set("name", "test1")

	return nil
}
