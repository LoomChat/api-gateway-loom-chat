package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/loomchat/api-gateway-loom-chat/pkg/log"
)

type Configs struct {
	Port            int           `json:"port"`
	Timeout         time.Duration `json:"timeout"`         // in milliseconds
	RouteConfigPath string        `json:"routeConfigPath"` // path to the global config file that includes all the
	// other config files.
}

func (c *Configs) String() string {
	return fmt.Sprintf("{port: %d, timeout: %d, routeConfigPath: %s}", c.Port, c.Timeout, c.RouteConfigPath)
}

func parseConfigs(defaultConfigs *Configs) error {
	configBytes, err := os.ReadFile("internal/config/configs.json")
	if err != nil {
		log.Error("Failed to read the configs.json file because: %s", err)
		return err
	}

	err = json.Unmarshal(configBytes, defaultConfigs)
	if err != nil {
		log.Error("Failed to json unmarshal the content of the configs file")
		return err
	}
	defaultConfigs.Timeout = defaultConfigs.Timeout * time.Millisecond

	return nil
}

func GetConfigs() *Configs {
	configs := Configs{
		Port:    8080,
		Timeout: 10 * time.Millisecond,
	}
	log.Debug("Default server configs: %v", &configs)

	err := parseConfigs(&configs)
	if err != nil {
		log.Error("Failed to get the configs due to: %s", err)
	}
	log.Debug("Server configs are set to: %v", &configs)

	return &configs
}
