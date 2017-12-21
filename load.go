package konf

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/ruggi/env"
)

// EnvTag is the tag used for injecting environment variables into the configuration during Load.
var EnvTag = env.DefaultTag

// LoadFile loads the configuration in the file specified at path into the target interface (which must be a pointer).
// It also injects environment variables into it, if any match is found. See github.com/ruggi/env for more.
func LoadFile(path string, target interface{}) error {
	formatter, err := parsePath(path)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return load(formatter, data, target)
}

// LoadData loads the configuration passed as raw data into the target interface (which must be a pointer).
// It also injects environment variables into it, if any match is found. See github.com/ruggi/env for more.
func LoadData(format Format, data []byte, target interface{}) error {
	formatter, ok := formatters[format]
	if !ok {
		return fmt.Errorf("invalid format")
	}
	return load(formatter, data, target)
}

func load(formatter formatter, data []byte, target interface{}) error {
	r := bytes.NewReader(data)
	if err := formatter.Decode(r, target); err != nil {
		return err
	}
	env.ParseInto(target)
	return nil
}
