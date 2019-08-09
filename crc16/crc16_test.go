package crc16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func packIndicatorMessage(rs485addr byte, data []byte) (msg []byte) {
	msg = []byte{rs485addr}
	msg = append(msg, data...)
	crc := CRC16MODBUS(msg)
	msg = append(msg, byte(crc), byte(crc>>8))
	return
}

func TestLog1(t *testing.T) {

	d1 := []byte{0x01, 0x05, 0x00, 0x00, 0xFF, 0x00}
	d2 := []byte{0x01, 0x05, 0x00, 0x00, 0xFF, 0x00, 0x8C, 0x3A}

	// assert equality
	assert.Equal(t, CRC16MODBUS([]byte{1, 4, 2, 0, 1}), uint16(0x3a8c), "they should be equal")
	assert.Equal(t, CRC16MODBUS(d2), uint16(0), "they should be equal")
	assert.NotEqual(t, packIndicatorMessage(5, d1[1:]), d2, "they should be equal")

	// assert inequality
	//assert.NotEqual(t, 123, 456, "they should not be equal")

}
