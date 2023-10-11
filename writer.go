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
	case bool:
		return writeData(w, Bool, v)
	case int:
		return writeData(w, Int, int32(v))
	case int8:
		return writeData(w, Int8, v)
	case int16:
		return writeData(w, Int16, v)
	case int32:
		return writeData(w, Int32, v)
	case int64:
		return writeData(w, Int64, v)
	case uint:
		return writeData(w, Uint, uint32(v))
	case uint8:
		return writeData(w, Uint8, v)
	case uint16:
		return writeData(w, Uint16, v)
	case uint32:
		return writeData(w, Uint32, v)
	case uint64:
		return writeData(w, Uint64, v)
	case float32:
		return writeData(w, Float32, v)
	case float64:
		return writeData(w, Float64, v)
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
	if err := writeData(w, Slice); err != nil {
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
	if err := writeData(w, Map); err != nil {
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

	return writeData(w, Break)
}

func writeString(w io.Writer, v string) error {
	if err := writeData(w, String); err != nil {
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
