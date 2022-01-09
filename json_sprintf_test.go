package json_sprintf_test

import (
	json_sprintf "json-sprintf"
	"testing"
)

type TestStruct struct {
	String                 string            `json:"string"`
	Int                    int               `json:"int"`
	Float                  float64           `json:"float"`
	TestChildStruct        TestChildStruct   `json:"test_child_struct"`
	TestChildStructs       []TestChildStruct `json:"test_child_structs"`
	TestChildStructPointer *TestChildStruct  `json:"test_child_struct_pointer"`
	TestEmbedStruct
}

type TestChildStruct struct {
	StringPointer *string  `json:"string_pointer"`
	IntPointer    *int     `json:"int_pointer"`
	FloatPointer  *float64 `json:"float_pointer"`
}

type TestEmbedStruct struct {
	Uint                 uint                 `json:"uint"`
	Bool                 bool                 `json:"bool"`
	TestEmbedChildStruct TestEmbedChildStruct `json:"test_embed_child_struct"`
}

type TestEmbedChildStruct struct {
	UintPointer *uint `json:"uint_pointer"`
	BoolPointer *bool `json:"bool_pointer"`
}

func TestExec(t *testing.T) {
	s := TestStruct{
		String:          "",
		Int:             0,
		Float:           0,
		TestChildStruct: TestChildStruct{},
		TestChildStructs: []TestChildStruct{{
			StringPointer: nil,
			IntPointer:    nil,
			FloatPointer:  nil,
		}},
		TestChildStructPointer: &TestChildStruct{},
		TestEmbedStruct:        TestEmbedStruct{},
	}

	got := json_sprintf.Exec(s)

	want := "fmt.Sprintf(`{\"string\":\"%s\",\"int\":%d,\"float\":%d,\"test_child_struct\":{\"string_pointer\":%v,\"int_pointer\":%v,\"float_pointer\":%v},\"test_child_structs\":[{\"string_pointer\":%v,\"int_pointer\":%v,\"float_pointer\":%v}],\"test_child_struct_pointer\":{\"string_pointer\":%v,\"int_pointer\":%v,\"float_pointer\":%v},\"uint\":%d,\"bool\":%t,\"test_embed_child_struct\":{\"uint_pointer\":%v,\"bool_pointer\":%v}}`,\nTestStruct.String,//TestStruct.String\nTestStruct.Int,//TestStruct.Int\nTestStruct.Float,//TestStruct.Float\nTestChildStruct.StringPointer,//TestChildStruct.StringPointer\nTestChildStruct.IntPointer,//TestChildStruct.IntPointer\nTestChildStruct.FloatPointer,//TestChildStruct.FloatPointer\nTestChildStruct.StringPointer,//TestChildStruct.StringPointer\nTestChildStruct.IntPointer,//TestChildStruct.IntPointer\nTestChildStruct.FloatPointer,//TestChildStruct.FloatPointer\nTestChildStruct.StringPointer,//TestChildStruct.StringPointer\nTestChildStruct.IntPointer,//TestChildStruct.IntPointer\nTestChildStruct.FloatPointer,//TestChildStruct.FloatPointer\nTestEmbedStruct.Uint,//TestEmbedStruct.Uint\nTestEmbedStruct.Bool,//TestEmbedStruct.Bool\nTestEmbedChildStruct.UintPointer,//TestEmbedChildStruct.UintPointer\nTestEmbedChildStruct.BoolPointer,//TestEmbedChildStruct.BoolPointer\n)"

	if got != want {
		t.Errorf("期待する結果じゃない \nwant: %s \ngot: %s", want, got)
	}

}
