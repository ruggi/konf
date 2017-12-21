package konf

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-yaml/yaml"

	"github.com/BurntSushi/toml"
)

// Format is a supported configuration format.
type Format int32

const (
	_ Format = iota
	// JSON is the JSON format: https://www.json.org
	JSON
	// YAML is the YAML format: http://yaml.org
	YAML
	// TOML is the TOML format: https://github.com/toml-lang/toml
	TOML
)

var formatters = map[Format]formatter{
	JSON: jsonFormatter{},
	TOML: tomlFormatter{},
	YAML: yamlFormatter{},
}

type formatter interface {
	Encode(w io.Writer, v interface{}) error
	Decode(r io.Reader, v interface{}) error
	extensions() []string
}

type jsonFormatter struct{}

func (f jsonFormatter) Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func (f jsonFormatter) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (f jsonFormatter) extensions() []string {
	return []string{".json"}
}

type tomlFormatter struct{}

func (f tomlFormatter) Encode(w io.Writer, v interface{}) error {
	return toml.NewEncoder(w).Encode(v)
}

func (f tomlFormatter) Decode(r io.Reader, v interface{}) error {
	_, err := toml.DecodeReader(r, v)
	return err
}

func (f tomlFormatter) extensions() []string {
	return []string{".toml"}
}

type yamlFormatter struct{}

func (f yamlFormatter) Encode(w io.Writer, v interface{}) error {
	out, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(out)
	return err
}

func (f yamlFormatter) Decode(r io.Reader, v interface{}) error {
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(in, v)
}

func (f yamlFormatter) extensions() []string {
	return []string{".yaml", ".yml"}
}
