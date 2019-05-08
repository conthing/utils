package converter

import (
	"testing"
)

// TODO func TestGb18030ToUtf8(t *testing.T)

func TestUint32ToIPString(t *testing.T) {
	addr := Uint32ToIPString(0x12345678)

	if addr != "18.52.86.120" {
		t.Errorf("Expected '18.52.86.120' but got: %s", addr)
	}
}

func TestIPStringToUint32_1(t *testing.T) {
	addr,err := IPStringToUint32("192.168.0.98")

	if addr != 0xc0a80062 || err != nil {
		t.Errorf("Expected '0xc0a80062' but got: 0x%x", addr)
	}
}

func TestIPStringToUint32_2(t *testing.T) {
	addr,err := IPStringToUint32("192.256.1.98")

	if addr != 0 || err != ErrInvalidIPString {
		t.Errorf("Expected error '%v' but got error '%v'", ErrInvalidIPString, err)
	}
}

func TestIPStringToUint32_3(t *testing.T) {
	addr,err := IPStringToUint32("192.25.98")

	if addr != 0 || err != ErrInvalidIPString {
		t.Errorf("Expected error '%v' but got error '%v'", ErrInvalidIPString, err)
	}
}

func TestIPStringToUint32_4(t *testing.T) {
	addr,err := IPStringToUint32("192.25.98.22.")

	if addr != 0 || err != ErrInvalidIPString {
		t.Errorf("Expected error '%v' but got error '%v'", ErrInvalidIPString, err)
	}
}

func TestIPStringToUint32_5(t *testing.T) {
	addr,err := IPStringToUint32("192.25.98.-1")

	if addr != 0 || err != ErrInvalidIPString {
		t.Errorf("Expected error '%v' but got error '%v'", ErrInvalidIPString, err)
	}
}