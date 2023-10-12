package tisp

import (
	"encoding/binary"
	"fmt"
	"io"
)

func Read(r io.Reader) ([]interface{}, error) {
	return readSlice(r)
}

func readValue(r io.Reader) (interface{}, error) {
	t, err := readData[_type](r)
	if err != nil {
		return nil, err
	}

	switch t {
	case _nil:
		return nil, nil
	case _bool:
		return readData[bool](r)
	case _int:
		v, err := readData[int32](r)
		return int(v), err
	case _int8:
		return readData[int8](r)
	case _int16:
		return readData[int16](r)
	case _int32:
		return readData[int32](r)
	case _int64:
		return readData[int64](r)
	case _uint:
		v, err := readData[uint32](r)
		return uint(v), err
	case _uint8:
		return readData[uint8](r)
	case _uint16:
		return readData[uint16](r)
	case _uint32:
		return readData[uint32](r)
	case _uint64:
		return readData[uint64](r)
	case _float32:
		return readData[float32](r)
	case _float64:
		return readData[float64](r)
	case _string:
		return readString(r)
	case _map:
		return readMap(r)
	case _slice:
		return readSlice(r)
	default:
		return nil, fmt.Errorf("expected type, found: %s", string(t))
	}
}

func readSlice(r io.Reader) ([]interface{}, error) {
	size, err := readData[uint32](r)
	if err != nil {
		return nil, err
	}

	arr := make([]interface{}, size)

	for i := 0; i < int(size); i++ {
		v, err := readValue(r)
		if err != nil {
			return nil, err
		}
		arr[i] = v
	}

	return arr, err
}

func readMap(r io.Reader) (map[string]interface{}, error) {
	mp := make(map[string]interface{})

	for {
		t, err := readData[_type](r)
		if err != nil {
			return nil, err
		}

		if t == _break {
			break
		} else if t != _string {
			return nil, fmt.Errorf("expected string, found: %s", string(t))
		}

		k, err := readString(r)
		if err != nil {
			return nil, err
		}

		v, err := readValue(r)
		if err != nil {
			return nil, err
		}

		mp[k] = v
	}

	return mp, nil
}

func readString(r io.Reader) (string, error) {
	buf, err := readBytes(r)
	if err != nil {
		return "", err
	}
	return string(buf[:]), nil
}

func readBytes(r io.Reader) ([]uint8, error) {
	size, err := readData[uint32](r)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, size)
	err = binary.Read(r, binary.LittleEndian, &buf)
	return buf, err
}

func readData[T interface{}](r io.Reader) (T, error) {
	var v T
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}
