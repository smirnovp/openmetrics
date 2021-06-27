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
	format := "currencies{name:\"%s\"} %s\n"
	f, err := os.Open(path.Clean(c.FileName))
	if err != nil {
		return "", err
	}
	defer f.Close()
	return convertFromReader(f, format)
}

// YamlData ...
type YamlData struct {
	Currencies []struct {
		Name  string `yaml:"name"`
		Value string `yaml:"value"`
	} `yaml:"currencies"`
}

func convertFromReader(r io.Reader, format string) (string, error) {

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
		_, err := buf.WriteString(fmt.Sprintf(format, cur.Name, cur.Value))
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
