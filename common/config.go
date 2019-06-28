package common

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// LoadConfig 从confName文件中，默认为"configuration.toml"加载配置到configPtr中
func LoadConfig(confName string, configPtr interface{}) (err error) {
	if len(confName) == 0 {
		confName = "configuration.toml"
	}

	// As the toml package can panic if TOML is invalid,
	// or elements are found that don't match members of
	// the given struct, use a defered func to recover
	// from the panic and output a useful error.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("could not load configuration file; invalid TOML (%s)", confName)
		}
	}()

	contents, err := ioutil.ReadFile(confName)
	if err != nil {
		return fmt.Errorf("could not load configuration file (%s): %v", confName, err.Error())
	}

	// Decode the configuration from TOML
	//
	// TODO: invalid input can cause a SIGSEGV fatal error (INVESTIGATE)!!!
	//       - test missing keys, keys with wrong type, ...
	err = toml.Unmarshal(contents, configPtr)
	if err != nil {
		return fmt.Errorf("unable to parse configuration file (%s): %v", confName, err.Error())
	}

	return nil
}
