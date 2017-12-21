package konf

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/ruggi/env"
)

// EnvTag is the tag used for injecting environment variables into the configuration during Load.
var EnvTag = env.DefaultTag

// Load loads the configuration in the file specified at path into the target interface (which must be a pointer).
// It also injects environment variables into it, if any match is found. See github.com/ruggi/env for more.
func Load(path string, target interface{}) error {
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

func load(formatter formatter, data []byte, target interface{}) error {
	r := bytes.NewReader(data)
	if err := formatter.Decode(r, target); err != nil {
		return err
	}
	env.ParseInto(target)
	return nil
}

// Save saves the given interface v to the file at the given path, using the format obtained from the path's extension,
// if supported.
func Save(path string, v interface{}) error {
	formatter, err := parsePath(path)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return formatter.Encode(file, v)
}
