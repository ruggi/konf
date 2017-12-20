# konf

A simple and straightforward configuration files loader.

## Description

konf loads configurations into structs, supporting multiple formats.

It supports direct files loading, or from `[]byte` raw data, and also injects environment variables into the structs, overriding the values in the configuration file.

## Supported formats

* JSON
* YAML
* TOML

## Usage

Get the package with `go get github.com/ruggi/konf`, then use it like this:

```go
type Config struct {
	Address string `json:"address" yaml:"address" env:"ADDRESS"`
	Port    int    `json:"port" yaml:"port" env:"PORT"`
}

func main() {
	var config Config
	err := konf.LoadFile("path/to/config.yaml", &config)
	if err != nil {
		// ...
	}
}
```

### Environment variables

Looking at the example above, if your config looks like
```yaml
address: "127.0.0.1"
port: 8080
```

and you have an environment variable `PORT=1234`, the final configuration struct will be
```go
Config{
    Address: "127.0.0.1",
    Port:    1234,
}
```

because the environment variable `PORT` (as specified in the struct's tag) overrides the value `8080` in the loaded configuration file.