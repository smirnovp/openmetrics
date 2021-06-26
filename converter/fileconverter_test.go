package converter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BadReader struct{}

func (b BadReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("Ошибка от BadReader")
}

func TestFileConverter(t *testing.T) {
	fc := NewFileConverter("testdata/currencies.yaml")
	metrics, err := fc.GetMetrics()
	assert.Equal(t, nil, err, "ошибка должна быть nil")
	assert.Equal(t, "currencies{name:usd} 70\ncurrencies{name:eur} 80\n", metrics, "должны быть одинаковы")

	fc = NewFileConverter("testdata/notexisting.yaml")
	metrics, err = fc.GetMetrics()
	assert.NotEqual(t, nil, err, "Должна быть ошибка")
	assert.Equal(t, "", metrics, "Должно быть пустое значение")

	fc = NewFileConverter("testdata/baddata.yaml")
	metrics, err = fc.GetMetrics()
	assert.NotEqual(t, nil, err, "Должна быть ошибка")
	assert.Equal(t, "", metrics, "Должно быть пустое значение")

	var br BadReader
	metrics, err = convertFromReader(br)
	assert.NotEqual(t, nil, err, "Должна быть ошибка")
	assert.Equal(t, "", metrics, "Должно быть пустое значение")
}
