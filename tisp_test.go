package tisp_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/andriimwks/tisp"
)

var testCase = []interface{}{
	// nil,
	true,
	123,
	int8(10),
	int16(89),
	int32(5),
	int64(1044),
	uint(164),
	uint8(51),
	uint16(754),
	uint32(156),
	uint64(224),
	float32(554.655),
	float64(1537.895),
	"hello",
	map[string]interface{}{"hello": "world"},
	[]interface{}{123, false},
}

func TestTISP(t *testing.T) {
	buf := new(bytes.Buffer)

	err := tisp.Write(buf, testCase...)
	if err != nil {
		t.Fatal(err)
	}

	vs, err := tisp.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	if len(vs) != len(testCase) {
		t.FailNow()
	}

	for i, v := range vs {
		vv := reflect.ValueOf(v)
		tv := reflect.ValueOf(testCase[i])

		if vv.Kind() != tv.Kind() ||
			!reflect.DeepEqual(vv.Interface(), tv.Interface()) {
			t.FailNow()
		}
	}
}
