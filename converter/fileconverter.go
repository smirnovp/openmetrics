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
	return convertFromFile(c.FileName)
}

func convertFromFile(fileName string) (string, error) {
	type YamlData struct {
		Currencies []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"currencies"`
	}

	f, err := os.Open(path.Clean(fileName))
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(f)
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
