package openrtb2

import (
	"bytes"
	"encoding/json"
	"strings"
	"sync"
)

var bufferPool = &sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// JSONMarshal can marshal object without HTML-chars escaping
// https://stackoverflow.com/a/28596225/5160055
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	str := strings.Trim(buffer.String(), "\n") // encoder.Encode always adds "\n" in the end but it may be changed any moment

	bufferPool.Put(buffer)

	return []byte(str), err
}
