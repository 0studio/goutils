package goutils

import (
	"bytes"
	"encoding/gob"
)

func GobEncode(i interface{}) (data []byte) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode(i)
	data = network.Bytes()
	return
}
func GobDecode(data []byte, i interface{}) bool {
	var network bytes.Buffer
	network.Write(data)
	dec := gob.NewDecoder(&network)
	err := dec.Decode(i)
	if err != nil {
		return false
	}

	return true
}
