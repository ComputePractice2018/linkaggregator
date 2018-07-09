package config

import (
	/*"github.com/BurntSushi/toml"
	"log"*/
)

type DatabaseConfig struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *DatabaseConfig) Read() {
	/*if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}*/

}
