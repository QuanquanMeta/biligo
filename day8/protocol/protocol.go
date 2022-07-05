// socket sticky package
package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode
func Encode(msg string) ([]byte, error) {
	// read size the convert to int32 (4 bytes)
	length := int32(len(msg))
	pkg := new(bytes.Buffer)

	// write msg head
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// write msg body
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode
func Decode(reader *bufio.Reader) (string, error) {
	// read msg length
	lengthByte, _ := reader.Peek(4) // reads first 4 char
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// buffered return
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// read the real msg

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
