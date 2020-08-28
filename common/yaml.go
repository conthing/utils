package common

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadConfig 从confName文件中，默认为"config.yaml"加载配置到configPtr中
func LoadYaml(confName string, configPtr interface{}) (err error) {
	if len(confName) == 0 {
		confName = "config.yaml"
	}

	contents, err := ioutil.ReadFile(confName)
	if err != nil {
		return fmt.Errorf("could not load configuration file (%s): %v", confName, err.Error())
	}

	// Decode the configuration from YAML
	err = yaml.Unmarshal(contents, configPtr)
	if err != nil {
		return fmt.Errorf("unable to parse configuration file (%s): %v", confName, err.Error())
	}

	return nil
}
