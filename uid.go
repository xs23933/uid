package uid

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"time"
)

const (
	size   = 12
	digits = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	Nil = UID{}
)

type UID [size]byte

func (u UID) Bytes() []byte {
	return u[:]
}

func (u UID) IsEmpty() bool {
	return bytes.Equal(u[:], []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
}

func (u UID) String() string {
	if bytes.Equal(u[:], []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
		return ""
	}
	return string(u[:])
}

func (u UID) Value() (driver.Value, error) {
	return u.String(), nil
}

func Must(u UID, err error) UID {
	if err != nil {
		panic(err)
	}
	return u
}

func (u UID) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *UID) UnmarshalText(text []byte) error {
	copy(u[:], []byte(text))
	return nil
}

func (u UID) MarshalBinary() ([]byte, error) {
	return u.Bytes(), nil
}

func (u *UID) UnmarshalBinary(data []byte) error {
	if len(data) != size {
		return fmt.Errorf("uuid: UUID must be exactly 16 bytes long, got %d bytes", len(data))
	}
	copy(u[:], data)

	return nil
}

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

func (u *UID) generate() {
	buf := make([]byte, 10)
	rand.Read(buf)
	seed := binary.BigEndian.Uint64(buf[:8])
	pre := uint64(binary.BigEndian.Uint16(buf[8:]))
	src := (uint64(time.Now().UnixNano()) - seed) / 5
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

func FromString(input string) (UID, error) {
	u := UID{}
	err := u.UnmarshalText([]byte(input))
	return u, err
}

func New() UID {
	u := new(UID)
	u.generate()
	return *u
}
