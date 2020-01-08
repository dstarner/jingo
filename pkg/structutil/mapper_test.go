package structutil_test

import (
	"testing"

	"github.com/dstarner/jingo/pkg/structutil"
	"github.com/stretchr/testify/require"
)

type SampleStruct struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type EmbededStruct struct {
	FieldStruct `json:"field"`
	Hello       string `json:"hello"`
}

type FieldStruct struct {
	OnePoint string        `json:"one_point"`
	Sample   *SampleStruct `json:"sample"`
}

func TestStructToMap_Normal(t *testing.T) {
	sample := SampleStruct{
		Name: "John Doe",
		ID:   "12121",
	}
	expected := map[string]interface{}{
		"name": "John Doe",
		"id":   "12121",
	}

	res := structutil.ToMap(sample, "json")
	require.NotNil(t, res)
	require.Equal(t, expected, res)

}
func TestStructToMap_FieldStruct(t *testing.T) {

	sample := &SampleStruct{
		Name: "John Doe",
		ID:   "12121",
	}
	field := FieldStruct{
		Sample:   sample,
		OnePoint: "yuhuhuu",
	}

	expected := map[string]interface{}{
		"one_point": "yuhuhuu",
		"sample":    sample,
	}

	res := structutil.ToMap(field, "json")
	require.NotNil(t, res)
	require.Equal(t, expected, res)
}

func TestStructToMap_EmbeddedStruct(t *testing.T) {

	sample := &SampleStruct{
		Name: "John Doe",
		ID:   "12121",
	}
	field := FieldStruct{
		Sample:   sample,
		OnePoint: "yuhuhuu",
	}

	embed := EmbededStruct{
		FieldStruct: field,
		Hello:       "WORLD!!!!",
	}

	expected := map[string]interface{}{
		"hello": "WORLD!!!!",
		"field": map[string]interface{}{
			"one_point": "yuhuhuu",
			"sample":    sample,
		},
	}

	res := structutil.ToMap(embed, "json")
	require.NotNil(t, res)
	require.Equal(t, expected, res)
}
