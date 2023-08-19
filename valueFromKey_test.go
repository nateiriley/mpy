package mpy_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/nateiriley/mpy"
)

func TestValueFromKey(t *testing.T) {
	byter, err := os.ReadFile("test.json")
	if err != nil {
		t.Error(err)
	}
	myThing := mpy.New()
	err = json.Unmarshal(byter, &myThing)
	if err != nil {
		t.Error(err)
	}
	myValue, err := myThing.GetValueFromKey("latitude")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(myValue)
}

func TestValuesFromKey(t *testing.T) {
	byter, err := os.ReadFile("test.json")
	if err != nil {
		t.Error(err)
	}
	myThing := mpy.New()
	err = json.Unmarshal(byter, &myThing)
	if err != nil {
		t.Error(err)
	}
	myValue, err := myThing.GetValuesFromKey("latitude")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(myValue)
}
