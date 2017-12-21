package konf

import (
	"fmt"
	"path/filepath"
)

func parsePath(path string) (formatter, error) {
	for _, formatter := range formatters {
		ext := filepath.Ext(path)
		for _, fe := range formatter.extensions() {
			if ext == fe {
				return formatter, nil
			}
		}
	}
	return nil, fmt.Errorf("format not recognized")
}
