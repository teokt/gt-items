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

func (r *Reader) Read(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer")
	}

	elem := rv.Elem()

	switch elem.Kind() {
	case reflect.String:
		return r.ReadString(v.(*string))
	case reflect.Struct:
		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
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
		return r.ReadRaw(v)
	}
}

func (r *Reader) ReadRaw(v any) error {
	return binary.Read(r.r, binary.LittleEndian, v)
}

func (r *Reader) ReadString(out *string) error {
	var size uint16
	if err := r.ReadRaw(&size); err != nil {
		return err
	}

	buf := make([]byte, size)
	if err := r.ReadRaw(buf); err != nil {
		return err
	}

	*out = string(buf)
	return nil
}

func (r *Reader) ReadEncryptedString(out *string, cryptID int, secretCode string) error {
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

	*out = string(buf)
	return nil
}
