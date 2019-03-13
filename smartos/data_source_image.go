package smartos

import (
	"github.com/caius/goadm"
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

	// client.Imgadm.GetByName(d.Get("name"))

	client := goadm.NewClient("127.0.0.1", "root", 2022)

	images, err := client.Imgadm().ListImages()
	if err != nil {
		return err
	}
	image := images[0]

	d.Set("id", image.Uuid)
	d.Set("name", image.Name)

	return nil
}
