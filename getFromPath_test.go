package mpy_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nateiriley/mpy"
)

type getTest3 struct {
	A string
	B string
	C int
	D map[string]any
}
type getTest2 struct {
	A string
	B string
	C getTest3
}
type getTest struct {
	A string
	B string
	C getTest2
}

func TestGetFromPath(t *testing.T) {
	var testMap = getTest{A: "a", B: "b", C: getTest2{A: "a2", B: "b2",
		C: getTest3{A: "a3", B: "b3", C: 1, D: map[string]any{"key1": "value1", "key2": "value2", "key3": 12}}}}
	testBytes, err := json.Marshal(testMap)
	if err != nil {
		t.Error(err)
	}
	myThing := mpy.New()
	err = json.Unmarshal(testBytes, &myThing)
	if err != nil {
		t.Error(err)
	}
	myValue, err := myThing.GetValueFromPath([]string{"C", "C", "D", "key1"})
	if err != nil {
		t.Error(err)
	}
	if myValue != "value1" {
		t.Error(myValue, "!= value1")
	}
	fmt.Println(myValue)
}
