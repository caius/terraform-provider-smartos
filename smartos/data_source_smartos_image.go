package smartos

import (
	"fmt"
	"log"

	"github.com/caius/goadm"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceSmartosImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSmartosImageRead,
		Schema: map[string]*schema.Schema{

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},

			// computed attributes
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSmartosImageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goadm.Client)
	log.Printf("Got client %s@%s:%d", client.User, client.Host, client.Port)

	images, err := client.Imgadm().ListImages()
	if err != nil {
		d.SetId("")
		return err
	}

	image, err := findImageByNameVersion(images, d.Get("name").(string), d.Get("version").(string))

	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(image.Uuid)
	d.Set("name", image.Name)
	d.Set("version", image.Version)

	return nil
}

// images.select { |i| i.name == name && i.version == version }
func findImageByNameVersion(images []goadm.Image, name string, version string) (*goadm.Image, error) {
	results := make([]goadm.Image, 0)
	for _, img := range images {
		if img.Name == name && img.Version == version {
			results = append(results, img)
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no image found with name %s, version %s", name, version)
	}

	if len(results) >= 2 {
		return nil, fmt.Errorf("too many images found with name %s, version %s (found %d, expected 1)", name, version, len(results))
	}

	return &results[0], nil
}
