package converter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// FileConverter ...
type FileConverter struct {
	FileName string
}

// NewFileConverter ...
func NewFileConverter(f string) *FileConverter {
	return &FileConverter{
		FileName: f,
	}
}

// GetMetrics ...
func (c *FileConverter) GetMetrics() (string, error) {
	f, err := os.Open(path.Clean(c.FileName))
	if err != nil {
		return "", err
	}
	defer f.Close()
	return convertFromReader(f)
}

func convertFromReader(r io.Reader) (string, error) {
	type YamlData struct {
		Currencies []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"currencies"`
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	var data YamlData
	err = yaml.Unmarshal(b, &data)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	for _, cur := range data.Currencies {
		_, err := buf.WriteString(fmt.Sprintf("currencies{name:%s} %s\n", cur.Name, cur.Value))
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
