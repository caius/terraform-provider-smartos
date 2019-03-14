package smartos

import (
	"errors"
	"log"

	"github.com/caius/goadm"
)

type Config struct {
	Host string
	User string
	Port int
}

func (c Config) Client() (*goadm.Client, error) {
	err := c.validate()
	if err != nil {
		return nil, err
	}

	client := goadm.NewClient(c.Host, c.User, c.Port)

	log.Printf("[INFO] SmartOS Provider configured for %s@%s:%d", c.Host, c.User, c.Port)

	return &client, nil
}

func (c Config) validate() error {
	if c.Host == "" {
		return errors.New("Host must be configured for the SmartOS provider")
	}

	if c.User == "" {
		return errors.New("User must be configured for the SmartOS provider")
	}

	if c.Port == 0 {
		return errors.New("Port must be configured for the SmartOS provider")
	}

	return nil
}
