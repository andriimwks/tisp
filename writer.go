package tisp

import (
	"encoding/binary"
	"fmt"
	"io"
)

func Write(w io.Writer, vs ...interface{}) error {
	return writeSlice(w, vs, false)
}

func writeValue(w io.Writer, v interface{}) error {
	switch v := v.(type) {
	case nil:
		return writeData(w, flagNil)
	case bool:
		return writeData(w, flagBool, v)
	case int:
		return writeData(w, flagBool, int32(v))
	case int8:
		return writeData(w, flagInt8, v)
	case int16:
		return writeData(w, flagInt16, v)
	case int32:
		return writeData(w, flagInt32, v)
	case int64:
		return writeData(w, flagInt64, v)
	case uint:
		return writeData(w, flagUint, uint32(v))
	case uint8:
		return writeData(w, flagUint8, v)
	case uint16:
		return writeData(w, flagUint16, v)
	case uint32:
		return writeData(w, flagUint32, v)
	case uint64:
		return writeData(w, flagUint64, v)
	case float32:
		return writeData(w, flagFloat32, v)
	case float64:
		return writeData(w, flagFloat64, v)
	case string:
		return writeString(w, v)
	case map[string]interface{}:
		return writeMap(w, v)
	case []interface{}:
		return writeSlice(w, v)
	}
	return fmt.Errorf("unsupported type: %T", v)
}

func writeSlice(w io.Writer, arr []interface{}, flag ...bool) error {
	if len(flag) <= 0 || flag[0] {
		if err := writeData(w, flagSlice); err != nil {
			return err
		}
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
	if err := writeData(w, flagMap); err != nil {
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

	return writeData(w, flagBreak)
}

func writeString(w io.Writer, v string) error {
	if err := writeData(w, flagString); err != nil {
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
