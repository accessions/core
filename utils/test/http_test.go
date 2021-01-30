package test

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"testing"
)

type Demo struct {
	Value int `json:"value"`
	Name string `json:"name"`
}
func TestName(t *testing.T) {

	rspStr := ``
	var mod Demo
	_, _ = jsonparser.ArrayEach([]byte(rspStr), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		values, _ := jsonparser.GetString(value, "value")
		fmt.Println(values)
		_ = json.Unmarshal(value, &mod)

	},"data", "list")
}
