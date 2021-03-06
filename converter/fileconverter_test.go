package converter

import (
	"fmt"
	"openmetrics/apiserver"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BadReader struct{}

func (b BadReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("Ошибка от BadReader")
}

// TestInterface канареечный тест интерфейса
func TestInterface(t *testing.T) {
	var _ apiserver.IConverter = &FileConverter{}
}

func TestFileConverter(t *testing.T) {
	fc := NewFileConverter("testdata/currencies.yaml")
	metrics, err := fc.GetMetrics()
	assert.Equal(t, nil, err, "ошибка должна быть nil")
	assert.Equal(t, "# TYPE currencies gauge\ncurrencies{name:\"usd\"} 70\ncurrencies{name:\"eur\"} 80\n", metrics, "должны быть одинаковы")

	fc = NewFileConverter("testdata/notexisting.yaml")
	metrics, err = fc.GetMetrics()
	assert.NotEqual(t, nil, err, "Должна быть ошибка")
	assert.Equal(t, "", metrics, "Должно быть пустое значение")

	fc = NewFileConverter("testdata/baddata.yaml")
	metrics, err = fc.GetMetrics()
	assert.NotEqual(t, nil, err, "Должна быть ошибка")
	assert.Equal(t, "", metrics, "Должно быть пустое значение")

	var br BadReader
	mtype := ""
	format := "some format"
	metrics, err = convertFromReader(br, mtype, format)
	assert.NotEqual(t, nil, err, "Должна быть ошибка")
	assert.Equal(t, "", metrics, "Должно быть пустое значение")
}
