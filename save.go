package konf

import "os"

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
