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
	t, err := readData[flag](r)
	if err != nil {
		return nil, err
	}

	switch t {
	case flagNil:
		return nil, nil
	case flagBool:
		return readData[bool](r)
	case flagInt:
		v, err := readData[int32](r)
		return int(v), err
	case flagInt8:
		return readData[int8](r)
	case flagInt16:
		return readData[int16](r)
	case flagInt32:
		return readData[int32](r)
	case flagInt64:
		return readData[int64](r)
	case flagUint:
		v, err := readData[uint32](r)
		return uint(v), err
	case flagUint8:
		return readData[uint8](r)
	case flagUint16:
		return readData[uint16](r)
	case flagUint32:
		return readData[uint32](r)
	case flagUint64:
		return readData[uint64](r)
	case flagFloat32:
		return readData[float32](r)
	case flagFloat64:
		return readData[float64](r)
	case flagString:
		return readString(r)
	case flagMap:
		return readMap(r)
	case flagSlice:
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
		t, err := readData[flag](r)
		if err != nil {
			return nil, err
		}

		if t == flagBreak {
			break
		} else if t != flagString {
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
