package propertiesReader

import (
	"bufio"
	"bytes"
	"os"
)

// ReadProptice .
func ReadProptice(file *os.File) (map[string]interface{}, error) {
	reader := bufio.NewScanner(file)
	properties := make(map[string]interface{})
	for reader.Scan() {
		line := bytes.TrimSpace(reader.Bytes()[:])
		if len(line) <= 1 || bytes.HasPrefix(line, []byte{'#'}) {
			continue
		}
		if split := bytes.IndexByte(line, '='); split != -1 {
			properties[string(line[:split])] = string(line[split:])
		}
	}
	return properties, nil
}
