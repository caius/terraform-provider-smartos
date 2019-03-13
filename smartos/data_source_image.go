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
				Computed: false,
				Required: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
		},
	}
}

func dataSourceImageRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*Client)
	client := goadm.NewClient("127.0.0.1", "root", 2022)

	images, err := client.Imgadm().ListImages()
	if err != nil {
		return err
	}

	var foundImage goadm.Image

	for _, image := range images {
		if d.Get("name") == image.Name && d.Get("version") == image.Version {
			foundImage = image
		}
	}

	if foundImage.Name != "" {
		d.Set("id", foundImage.Uuid)
	}

	return nil
}
