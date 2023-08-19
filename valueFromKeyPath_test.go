package mpy_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/nateiriley/mpy"
)

func TestValueFromKeyPath(t *testing.T) {
	byter, err := os.ReadFile("test2.json")
	if err != nil {
		t.Error(err)
	}
	myThing := mpy.New()
	err = json.Unmarshal(byter, &myThing)
	if err != nil {
		t.Error(err)
	}
	myValue, err := myThing.GetValueFromKeyPath("map3.map3._id")
	if err != nil {
		t.Error(err)
	}
	if myValue != "64e0c5000ddcf0aaf2c56b86" {
		t.Error(myValue, "!=64e0c5000ddcf0aaf2c56b86")
	}
}

func TestValuesFromKeyPath(t *testing.T) {
	byter, err := os.ReadFile("test2.json")
	if err != nil {
		t.Error(err)
	}
	myThing := mpy.New()
	err = json.Unmarshal(byter, &myThing)
	if err != nil {
		t.Error(err)
	}
	myValues, err := myThing.GetValuesFromKeyPath("map3.favoriteFruit")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(myValues)
}
