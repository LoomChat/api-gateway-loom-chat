package config

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	RegexEnvVarNameInsideString = "\\$[a-zA-Z0-9_]+" // TODO: Write a grouping regex!
)

// Function ReplaceEnvVarsIn identifies env variables in a given text string value and replaces each
// of those env variable occurrences with their value found from a given map of env variables.
//
// An env variable occurrence inside a string value is considered to be any substring value that is
// preceded by '$' character and contains one or more alphanumeric and '_' (underscore) characters.
func ReplaceEnvVarsIn(text string, envVars map[string]string) (string, error) {
	re := regexp.MustCompile(RegexEnvVarNameInsideString)
	envVarNames := re.FindAll([]byte(text), -1)

	for _, envVarName := range envVarNames {
		envVarNameStr := string(envVarName)
		if envVarVal, ok := envVars[envVarNameStr[1:]]; ok {
			text = strings.ReplaceAll(text, envVarNameStr, envVarVal)
		} else {
			return "", fmt.Errorf("Environment variable '%s' is not defined!", envVarNameStr)
		}
	}

	return text, nil
}

func ReplaceEnvVarsInConfigs(configs any, envVars map[string]string) error {
	configsReflVal := reflect.ValueOf(configs).Elem()

	if configsReflVal.Kind() == reflect.String {
		withEnvVarsReplaced, err := ReplaceEnvVarsIn(configsReflVal.String(), envVars)
		if err != nil {
			return err
		}

		configsReflVal.SetString(withEnvVarsReplaced)
	}

	if configsReflVal.Kind() == reflect.Struct {
		for i := 0; i < configsReflVal.NumField(); i++ {
			field := configsReflVal.Field(i)
			if field.Kind() == reflect.String {
				withEnvVarsReplaced, err := ReplaceEnvVarsIn(field.String(), envVars)
				if err != nil {
					return err
				}

				field.SetString(withEnvVarsReplaced)
			}
		}
	} else {
		return fmt.Errorf("Non struct type parameter is passed!")
	}

	return nil
}
