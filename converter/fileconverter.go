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
	mtype := "# TYPE currencies gauge\n"
	format := "currencies{name:\"%s\"} %s\n"
	f, err := os.Open(path.Clean(c.FileName))
	if err != nil {
		return "", err
	}
	defer f.Close()
	return convertFromReader(f, mtype, format)
}

// YamlData ...
type YamlData struct {
	Currencies []struct {
		Name  string `yaml:"name"`
		Value string `yaml:"value"`
	} `yaml:"currencies"`
}

func convertFromReader(r io.Reader, mtype, format string) (string, error) {

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

	_, err = buf.WriteString(mtype)
	if err != nil {
		return "", err
	}

	for _, cur := range data.Currencies {
		_, err := buf.WriteString(fmt.Sprintf(format, cur.Name, cur.Value))
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
