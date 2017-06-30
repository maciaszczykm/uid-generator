package main

import (
	"bytes"
	"encoding/binary"
	"time"
	"net"
)

// Generator struct contains unique ID generator parameters.
type Generator struct {
	counter int16
	mac     []byte
}

// NewGenerator constructs new  unique ID generator object.
func NewGenerator() Generator {
	return Generator {
		counter: 0,
		mac:     getMacAddress(),
	}
}

// TODO Fill with random data if empty.
func getMacAddress() net.HardwareAddr {
	mac := make([]byte, 6)
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		if len(i.HardwareAddr) >= 6 {
			copy(mac, i.HardwareAddr)
		}
	}
	return mac
}

// Generate unique ID.
func (g *Generator) Generate() UID {
	var buffer bytes.Buffer

	now := make([]byte, 8)
	binary.BigEndian.PutUint64(now, uint64(time.Now().UnixNano()))
	buffer.Write(now) // 8 bytes to be able to sort by date and make IDs almost unique.

	g.counter++
	buffer.WriteByte(byte(g.counter))
	buffer.WriteByte(byte(g.counter >> 8)) // 2 bytes to differ IDs from same time.

	buffer.Write(g.mac) // 6 bytes to differ IDs from different instances.

	return buffer.Bytes()
}