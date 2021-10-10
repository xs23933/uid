package uid

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

const (
	size   = 12
	digits = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	// Nil Null
	Nil = UID{}
)

// UID UID
type UID [size]byte

// Bytes out bytes
func (u UID) Bytes() []byte {
	return u[:]
}

// IsEmpty check is Empty
func (u UID) IsEmpty() bool {
	return bytes.Equal(u[:], []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
}

// String out string
func (u UID) String() string {
	if bytes.Equal(u[:], []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
		return ""
	}
	return string(u[:])
}

// Value driver.sql
func (u UID) Value() (driver.Value, error) {
	return u.String(), nil
}

// MarshalText MarshalText interface
func (u UID) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

// UnmarshalText UnmarshalText interface
func (u *UID) UnmarshalText(text []byte) error {
	if len(text) != size {
		return fmt.Errorf("uid: UID must be exactly 12 bytes long, got %d bytes", len(text))
	}
	copy(u[:], []byte(text))
	return nil
}

// MarshalBinary MarshalBinary interface
func (u UID) MarshalBinary() ([]byte, error) {
	return u.Bytes(), nil
}

// UnmarshalBinary UnmarshalBinary interface
func (u *UID) UnmarshalBinary(data []byte) error {
	if len(data) != size {
		return fmt.Errorf("uid: UID must be exactly 12 bytes long, got %d bytes", len(data))
	}
	copy(u[:], data)

	return nil
}

// GormDataType schema.Field DataType
func (UID) GormDataType() string {
	return "char(12)"
}

// Scan driver.Scaner
func (u *UID) Scan(src interface{}) error {
	switch src := src.(type) {
	case UID:
		*u = src
		return nil
	case []byte:
		if len(src) == size {
			return u.UnmarshalBinary(src)
		}
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}
	return fmt.Errorf("uid: cannot convert %T to UID", src)
}

// generate build uid
func (u *UID) generate() {
	buf := make([]byte, 10)
	rand.Read(buf)
	seed := binary.BigEndian.Uint64(buf[:8])
	pre := uint64(binary.BigEndian.Uint16(buf[8:]))
	src := (uint64(time.Now().UnixNano()) - seed) / 3
	var a [64 + 1]byte
	i := len(a)
	l := uint64(len(digits))
	for src >= l {
		i--
		q := src / l
		a[i] = digits[uint((src - q*l))]
		src = q
	}
	i--
	a[i] = digits[uint(src)]
	for pre >= l {
		i--
		q := pre / l
		a[i] = digits[uint((pre - q*l))]
		pre = q
	}
	i--
	a[i] = digits[uint(pre)]
	copy(u[:], a[i:])
}

// FromString parse string
func FromString(input string) (UID, error) {
	u := UID{}
	if len(input) != size {
		return Nil, fmt.Errorf("uid: UID must be exactly 12 bytes long, got %d bytes", len(input))
	}
	err := u.UnmarshalText([]byte(strings.ToUpper(input)))
	return u, err
}

func FromBytes(data []byte) (UID, error) {
	u := UID{}
	err := u.UnmarshalBinary(data)
	return u, err
}

// New build uid
func New() UID {
	u := new(UID)
	u.generate()
	return *u
}
