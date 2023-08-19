package mpy_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nateiriley/mpy"
)

type setTest3 struct {
	A string
	B string
	C int
	D map[string]any
}
type setTest2 struct {
	A string
	B string
	C setTest3
}
type setTest struct {
	A string
	B string
	C setTest2
}

func TestSetFromPath(t *testing.T) {
	var testMap = setTest{A: "a", B: "b", C: setTest2{A: "a2", B: "b2",
		C: setTest3{A: "a3", B: "b3", C: 1, D: map[string]any{"key1": "value1", "key2": "value2", "key3": 12}}}}
	testBytes, err := json.Marshal(testMap)
	if err != nil {
		t.Error(err)
	}
	myThing := mpy.New()
	err = json.Unmarshal(testBytes, &myThing)
	if err != nil {
		t.Error(err)
	}
	myValue, err := myThing.SetValueFromPath([]string{"C", "C", "D", "key4"}, "here")
	if err != nil {
		t.Error(err)
	}
	if myValue != "value1" {
		t.Error(myValue, "!= value1")
	}
	fmt.Println(myValue)
}
