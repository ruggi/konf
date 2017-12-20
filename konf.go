package konf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/ruggi/env"

	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
)

type Format string

const (
	JSON Format = "json|jsn"
	TOML Format = "toml"
	YAML Format = "yaml|yml"
)

type loader func([]byte, interface{}) error

var formats = map[Format]loader{
	JSON: loadJSON,
	TOML: loadTOML,
	YAML: loadYAML,
}

// EnvTag is the tag used for injecting environment variables into the configuration during Load.
var EnvTag = env.DefaultTag

// LoadFile loads the configuration in the file specified at path into the target interface (which must be a pointer).
// It also injects environment variables into it, if any match is found. See github.com/ruggi/env for more.
func LoadFile(path string, target interface{}) error {
	for format, loader := range formats {
		re := regexp.MustCompile(fmt.Sprintf("\\.(%s)$", format))
		if !re.MatchString(path) {
			continue
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		return load(loader, data, target)
	}
	return fmt.Errorf("format not recognized")
}

// LoadData loads the configuration passed as raw data into the target interface (which must be a pointer).
// It also injects environment variables into it, if any match is found. See github.com/ruggi/env for more.
func LoadData(format Format, data []byte, target interface{}) error {
	loader, ok := formats[format]
	if !ok {
		return fmt.Errorf("invalid format")
	}
	return load(loader, data, target)
}

func load(loader loader, data []byte, target interface{}) error {
	if err := loader(data, target); err != nil {
		return err
	}
	env.ParseInto(target)
	return nil
}

func loadJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func loadTOML(data []byte, v interface{}) error {
	_, err := toml.Decode(string(data), v)
	return err
}

func loadYAML(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
