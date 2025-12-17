package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/loomchat/api-gateway-loom-chat/pkg/log"
)

type Environment struct {
	Variables map[string]string `json:"variables"`
}

func (e Environment) String() string {
	return fmt.Sprintf("{variables: %s}", e.Variables)
}

func GetEnv() *Environment {
	env, err := parseEnv()
	if err != nil {
		log.Fatal("Failed to get the environment data due to: %s", err)
		log.Fatal("Exiting...")
		os.Exit(1)
	}

	return env
}

const envFilename = "envs/env_dev.json"

func parseEnv() (*Environment, error) {
	envContent, err := os.ReadFile(envFilename)
	if err != nil {
		log.Error("Failed to read from the env file '%s' due to: %s", envFilename, err)
		return nil, err
	}

	var env Environment
	err = json.Unmarshal(envContent, &env)
	if err != nil {
		log.Error("Failed to json unmarshal the env file '%s' content", envFilename)
		return nil, err
	}

	return &env, nil
}
