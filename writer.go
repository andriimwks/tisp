package tisp

import (
	"encoding/binary"
	"fmt"
	"io"
)

func Write(w io.Writer, vs ...interface{}) error {
	if err := writeData(w, uint32(len(vs))); err != nil {
		return err
	}

	for _, v := range vs {
		if err := writeValue(w, v); err != nil {
			return err
		}
	}

	return nil
}

func writeValue(w io.Writer, v interface{}) error {
	switch v := v.(type) {
	case nil:
		return writeData(w, _nil)
	case bool:
		return writeData(w, _bool, v)
	case int:
		return writeData(w, _int, int32(v))
	case int8:
		return writeData(w, _int8, v)
	case int16:
		return writeData(w, _int16, v)
	case int32:
		return writeData(w, _int32, v)
	case int64:
		return writeData(w, _int64, v)
	case uint:
		return writeData(w, _uint, uint32(v))
	case uint8:
		return writeData(w, _uint8, v)
	case uint16:
		return writeData(w, _uint16, v)
	case uint32:
		return writeData(w, _uint32, v)
	case uint64:
		return writeData(w, _uint64, v)
	case float32:
		return writeData(w, _float32, v)
	case float64:
		return writeData(w, _float64, v)
	case string:
		return writeString(w, v)
	case map[string]interface{}:
		return writeMap(w, v)
	case []interface{}:
		return writeSlice(w, v)
	}
	return fmt.Errorf("unsupported type: %T", v)
}

func writeSlice(w io.Writer, arr []interface{}) error {
	if err := writeData(w, _slice); err != nil {
		return err
	}

	if err := writeData(w, uint32(len(arr))); err != nil {
		return err
	}

	for _, v := range arr {
		if err := writeValue(w, v); err != nil {
			return err
		}
	}

	return nil
}

func writeMap(w io.Writer, m map[string]interface{}) error {
	if err := writeData(w, _map); err != nil {
		return err
	}

	for k, v := range m {
		if err := writeString(w, k); err != nil {
			return err
		}

		if v, ok := v.(map[string]interface{}); ok {
			if err := writeMap(w, v); err != nil {
				return err
			}
			continue
		}

		if err := writeValue(w, v); err != nil {
			return err
		}
	}

	return writeData(w, _break)
}

func writeString(w io.Writer, v string) error {
	if err := writeData(w, _string); err != nil {
		return err
	}
	return writeByteSlice(w, []byte(v))
}

func writeByteSlice(w io.Writer, data []byte) error {
	if err := writeData(w, uint32(len(data))); err != nil {
		return err
	}
	return writeData(w, data)
}

func writeData(w io.Writer, vs ...interface{}) error {
	for _, v := range vs {
		if err := binary.Write(w, binary.LittleEndian, v); err != nil {
			return err
		}
	}
	return nil
}
