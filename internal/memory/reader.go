package memory

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
)

type Reader struct {
	r *bytes.Reader
}

func NewReader(data []byte) *Reader {
	return &Reader{
		r: bytes.NewReader(data),
	}
}

func (r *Reader) Read(dest any) error {
	val := reflect.ValueOf(dest)

	if val.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer")
	}

	val = val.Elem()

	switch val.Kind() {
	case reflect.String:
		return r.ReadString(dest.(*string))
	case reflect.Struct:
		for i := range val.NumField() {
			field := val.Field(i)

			if !field.CanSet() {
				continue
			}

			fieldAddr := field.Addr().Interface()
			if err := r.Read(fieldAddr); err != nil {
				return err
			}
		}
		return nil
	default:
		return r.ReadRaw(dest)
	}
}

func (r *Reader) ReadRaw(dest any) error {
	return binary.Read(r.r, binary.LittleEndian, dest)
}

func (r *Reader) ReadString(dest *string) error {
	var size uint16
	if err := r.ReadRaw(&size); err != nil {
		return err
	}

	buf := make([]byte, size)
	if err := r.ReadRaw(buf); err != nil {
		return err
	}

	*dest = string(buf)
	return nil
}

func (r *Reader) ReadEncryptedString(dest *string, cryptID int, secretCode string) error {
	var size uint16
	if err := r.ReadRaw(&size); err != nil {
		return err
	}

	buf := make([]byte, size)
	if err := r.ReadRaw(buf); err != nil {
		return err
	}

	secretLen := len(secretCode)

	for i := range buf {
		buf[i] ^= secretCode[(i+cryptID)%secretLen]
	}

	*dest = string(buf)
	return nil
}
