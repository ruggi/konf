package konf_test

import (
	"os"
	"testing"

	"github.com/ruggi/konf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Config struct {
	User User `json:"user" yaml:"user" toml:"user"`
}

type User struct {
	Name     string   `json:"name" yaml:"name" toml:"name" env:"name"`
	Age      int      `json:"age" yaml:"age" toml:"age" env:"age"`
	Likes    []string `json:"likes" yaml:"likes" toml:"likes"`
	Children []Child  `json:"children" yaml:"children" toml:"children"`
}

type Child struct {
	Name string `json:"name" yaml:"name" toml:"name"`
}

var expected = Config{
	User: User{
		Name: "homer",
		Age:  38,
		Likes: []string{
			"donuts",
			"beer",
			"marge",
		},
		Children: []Child{
			Child{Name: "bart"},
			Child{Name: "lisa"},
			Child{Name: "maggie"},
		},
	},
}

func TestLoadFile(t *testing.T) {
	testCases := []struct {
		name        string
		path        string
		shouldError bool
	}{
		{
			name: "good json",
			path: "examples/user.json",
		},
		{
			name: "good yaml",
			path: "examples/user.yaml",
		},
		{
			name: "good toml",
			path: "examples/user.toml",
		},
		{
			name:        "non existing file",
			path:        "notfound",
			shouldError: true,
		},
		{
			name:        "unsupported extension",
			path:        "examples/unsupported.bad",
			shouldError: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var config Config
			err := konf.LoadFile(tc.path, &config)
			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, config)
			}
		})
	}
}

func TestLoadWithEnv(t *testing.T) {
	var config Config
	os.Setenv("name", "homer jay simpson")
	err := konf.LoadFile("examples/user.json", &config)
	require.NoError(t, err)
	assert.Equal(t, "homer jay simpson", config.User.Name)
}
