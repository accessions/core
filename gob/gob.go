package gob

import (
	"bytes"
	"encoding/gob"
)


type Sgob struct {}

// EnCoder
func (*Sgob) EnCoder(data interface{}) ([]byte, error) {

	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&data)
	if err != nil {
		return nil, nil
	}
	return buffer.Bytes(), nil

}

// DeCoder
func (*Sgob) DeCoder(data []byte) (interface{}, error){
	decoder := gob.NewDecoder(bytes.NewReader(data))
	var reply interface{}
	err := decoder.Decode(&reply)
	if err != nil {
		return nil, nil
	}
	return reply, nil
}
