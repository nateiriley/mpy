package mpy_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/nateiriley/mpy"
)

type structTest3 struct {
	A string
	B string
	C int
	D map[string]any
}
type structTest2 struct {
	A string
	B string
	C structTest3
}
type structTest struct {
	A string
	B string
	C structTest2
}

func (t *structTest3) MpyGet() {}

func TestGetStuctFromPath(t *testing.T) {
	var testMap = structTest{A: "a", B: "b", C: structTest2{A: "a2", B: "b2",
		C: structTest3{A: "a3", B: "b3", C: 1, D: map[string]any{"key1": "value1", "key2": "value2", "key3": 12}}}}
	testBytes, err := json.Marshal(testMap)
	if err != nil {
		t.Error(err)
	}

	myThing := mpy.New()
	err = json.Unmarshal(testBytes, &myThing)
	if err != nil {
		t.Error(err)
	}

	var fillTest structTest3
	err = myThing.GetStuctFromPath([]string{"C", "C"}, &fillTest)
	if err != nil {
		t.Error(err)
	}
	var correctStruct = structTest3{A: "a3", B: "b3", C: 1, D: map[string]any{"key1": "value1", "key2": "value2", "key3": 12}}
	if reflect.DeepEqual(correctStruct, fillTest) {
		t.Error(fillTest, "!=", correctStruct)
	}
}
