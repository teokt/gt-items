package memory

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"reflect"
)

type Reader struct {
	data   []byte
	offset int
	order  binary.ByteOrder
}

func NewReader(data []byte) *Reader {
	return &Reader{
		data:   data,
		offset: 0,
		order:  binary.LittleEndian,
	}
}

func (r *Reader) ReadBytes(n int) ([]byte, error) {
	if r.offset+n > len(r.data) {
		return nil, errors.New("read out of bounds")
	}
	b := r.data[r.offset : r.offset+n]
	r.offset += n
	return b, nil
}

func (r *Reader) ReadUint8(out *uint8) error {
	b, err := r.ReadBytes(1)
	if err != nil {
		return err
	}
	*out = b[0]
	return nil
}

func (r *Reader) ReadUint16(out *uint16) error {
	b, err := r.ReadBytes(2)
	if err != nil {
		return err
	}
	*out = r.order.Uint16(b)
	return nil
}

func (r *Reader) ReadUint32(out *uint32) error {
	b, err := r.ReadBytes(4)
	if err != nil {
		return err
	}
	*out = r.order.Uint32(b)
	return nil
}

func (r *Reader) ReadUint64(out *uint64) error {
	b, err := r.ReadBytes(8)
	if err != nil {
		return err
	}
	*out = r.order.Uint64(b)
	return nil
}

func (r *Reader) ReadInt8(out *int8) error {
	var n uint8
	if err := r.ReadUint8(&n); err != nil {
		return err
	}
	*out = int8(n)
	return nil
}

func (r *Reader) ReadInt16(out *int16) error {
	var n uint16
	if err := r.ReadUint16(&n); err != nil {
		return err
	}
	*out = int16(n)
	return nil
}

func (r *Reader) ReadInt32(out *int32) error {
	var n uint32
	if err := r.ReadUint32(&n); err != nil {
		return err
	}
	*out = int32(n)
	return nil
}

func (r *Reader) ReadInt64(out *int64) error {
	var n uint64
	if err := r.ReadUint64(&n); err != nil {
		return err
	}
	*out = int64(n)
	return nil
}

func (r *Reader) ReadFloat32(out *float32) error {
	var bits uint32
	if err := r.ReadUint32(&bits); err != nil {
		return err
	}
	*out = math.Float32frombits(bits)
	return nil
}

func (r *Reader) ReadFloat64(out *float64) error {
	var bits uint64
	if err := r.ReadUint64(&bits); err != nil {
		return err
	}
	*out = math.Float64frombits(bits)
	return nil
}

func (r *Reader) ReadString(out *string) error {
	var size uint16
	if err := r.ReadUint16(&size); err != nil {
		return err
	}

	b, err := r.ReadBytes(int(size))
	if err != nil {
		return err
	}

	*out = string(b)
	return nil
}

func (r *Reader) ReadEncryptedString(out *string, cryptID int, secretCode string) error {
	var size uint16
	if err := r.ReadUint16(&size); err != nil {
		return err
	}

	b, err := r.ReadBytes(int(size))
	if err != nil {
		return err
	}

	secretBytes := []byte(secretCode)
	secretLen := len(secretBytes)

	for i := range b {
		b[i] ^= secretBytes[(i+cryptID)%secretLen]
	}

	*out = string(b)
	return nil
}

func (r *Reader) ReadStaticArray(arrayPtr any) error {
	val := reflect.ValueOf(arrayPtr)
	if val.Kind() != reflect.Ptr {
		return errors.New("must pass array pointer")
	}

	arr := val.Elem()
	if arr.Kind() != reflect.Array {
		return errors.New("must point to an array")
	}

	size := arr.Len()
	elemType := arr.Type().Elem()

	for i := range size {
		elem := reflect.New(elemType).Elem()
		if err := r.Read(elem.Addr().Interface()); err != nil {
			return fmt.Errorf("array element %d: %w", i, err)
		}
		arr.Index(i).Set(elem)
	}

	return nil
}

func (r *Reader) Read(v any) error {
	switch v := v.(type) {
	case *uint8:
		return r.ReadUint8(v)
	case *uint16:
		return r.ReadUint16(v)
	case *uint32:
		return r.ReadUint32(v)
	case *uint64:
		return r.ReadUint64(v)
	case *int8:
		return r.ReadInt8(v)
	case *int16:
		return r.ReadInt16(v)
	case *int32:
		return r.ReadInt32(v)
	case *int64:
		return r.ReadInt64(v)
	case *float32:
		return r.ReadFloat32(v)
	case *float64:
		return r.ReadFloat64(v)
	case *string:
		return r.ReadString(v)
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer")
	}

	switch rv.Elem().Kind() {
	case reflect.Array:
		return r.ReadStaticArray(v)
	}

	return fmt.Errorf("unsupported type: %T", v)
}
