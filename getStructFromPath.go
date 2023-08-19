package mpy

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CustomStruct interface {
	MpyGet()
}

// Given a custom struct and a path, this will retrieve that struct from the path
// Your struct has to implement MpyGet()
// Example Method
// func (mS *myStruct) MpyGet() {}
// Can only return a single struct
func (mv Map) GetStuctFromPath(path any, customStruct CustomStruct) error {
	var pathToGet Path
	switch v := path.(type) {
	case []string:
		pathToGet = v
	case string:
		pathToGet = strings.Split(v, charToSplitPath)
	}

	var currentLevel map[string]any
	for i, key := range pathToGet {
		if i == 0 {
			value, ok := searchLevel(mv, key)
			if !ok {
				return getPathError(mv, pathToGet, i)
			}
			if i == len(pathToGet)-1 {
				tempBytes, err := json.Marshal(value)
				if err != nil {
					return err
				}
				err = json.Unmarshal(tempBytes, customStruct)
				if err != nil {
					return err
				}
			} else {
				currentLevel, ok = value.(map[string]any)
				if !ok {
					return fmt.Errorf("%s cannot is not a map, can't go further down [%s]", value, pathToGet)
				}
			}
		} else {
			value, ok := searchLevel(currentLevel, key)
			if !ok {
				return getPathError(currentLevel, pathToGet, i)
			}
			if i == len(pathToGet)-1 {
				tempBytes, err := json.Marshal(value)
				if err != nil {
					return err
				}
				err = json.Unmarshal(tempBytes, customStruct)
				if err != nil {
					return err
				}
			} else {
				currentLevel, ok = value.(map[string]any)
				if !ok {
					return fmt.Errorf("%s cannot is not a map, can't go further down [%s]", value, pathToGet)
				}
			}
		}
	}
	return nil
}
