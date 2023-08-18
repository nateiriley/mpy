package mpy

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Path []string

// Be able to change the value that splits the path
var charToSplitPath string = "."

// Change the value you want to split paths
func ChangePathSplitChar(newChar string) {
	charToSplitPath = newChar
	fmt.Println("New path split of ", newChar)
}

// Path can either a single string seperated by a period, or a string slice
// Example for string this.is.a.path would give you the value this|is|a|path
// Example of string slice ["this","is","a","path"] would give the same result
func (mv Map) GetValueForPath(path any) (any, error) {
	var pathToGet Path
	switch v := path.(type) {
	case []string:
		pathToGet = v
	case string:
		pathToGet = strings.Split(v, charToSplitPath)
	}

	var currentLevel map[string]any
	var returnValue any
	for i, key := range pathToGet {
		if i == 0 {
			value, ok := searchLevel(mv, key)
			if !ok {
				return nil, getPathError(mv, pathToGet, i)
			}
			if i == len(pathToGet) {
				returnValue = value
			} else {
				currentLevel, ok = value.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("%s cannot is not a map, can't go further down [%s]", value, pathToGet)
				}
			}
		} else {
			value, ok := searchLevel(currentLevel, key)
			if !ok {
				return nil, getPathError(currentLevel, pathToGet, i)
			}
			if i == len(pathToGet) {
				returnValue = value
			} else {
				currentLevel, ok = value.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("%s cannot is not a map, can't go further down [%s]", value, pathToGet)
				}
			}
		}
	}

	return returnValue, nil
}

type CustomStruct interface {
	MpyGet()
}

// Given a custom struct and a path, this will retrieve that struct from the path
// Your struct has to implement MpyGet()
// Example Method
// func (mS *myStruct) MpyGet() {}
// 	***********************************************
// Warning can't get struct within structs as of now
// 	***********************************************
func (mv Map) GetStuctFromPath(customStruct CustomStruct, path any) error {
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
				json.Unmarshal(tempBytes, customStruct)
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
				json.Unmarshal(tempBytes, customStruct)
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

func searchLevel(currentMap map[string]any, key string) (any, bool) {
	value, ok := currentMap[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func getPathError(currentMap map[string]any, path []string, level int) error {
	var errorString string
	for j := 0; j < level; j++ {
		if j == 0 {
			errorString = path[j]
		} else {
			errorString = errorString + "." + path[j]
		}
	}
	return fmt.Errorf("[%s] -> key = '%s' not in %v", errorString, path[level], currentMap)
}
