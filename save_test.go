package konf_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ruggi/konf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	cfg := Config{
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
	testCases := []struct {
		path     string
		expected string
	}{
		{
			path:     "homer.json",
			expected: "{\"user\":{\"name\":\"homer\",\"age\":38,\"likes\":[\"donuts\",\"beer\",\"marge\"],\"children\":[{\"name\":\"bart\"},{\"name\":\"lisa\"},{\"name\":\"maggie\"}]}}",
		},
		{
			path:     "homer.yaml",
			expected: "user:\n  name: homer\n  age: 38\n  likes:\n  - donuts\n  - beer\n  - marge\n  children:\n  - name: bart\n  - name: lisa\n  - name: maggie",
		},
		{
			path:     "homer.toml",
			expected: "[user]\n  name = \"homer\"\n  age = 38\n  likes = [\"donuts\", \"beer\", \"marge\"]\n\n  [[user.children]]\n    name = \"bart\"\n\n  [[user.children]]\n    name = \"lisa\"\n\n  [[user.children]]\n    name = \"maggie\"",
		},
	}

	for _, tc := range testCases {
		path := filepath.Join(os.TempDir(), tc.path)
		err := konf.Save(path, cfg)
		require.NoError(t, err)
		content, err := ioutil.ReadFile(path)
		require.NoError(t, err)
		assert.Equal(t, tc.expected, strings.TrimSpace(string(content)))
	}
}
