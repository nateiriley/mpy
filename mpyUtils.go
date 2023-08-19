package mpy

import (
	"encoding/json"
	"fmt"
	"strings"
)

// The basic search for this will be a []string in order for better error handling
type Path []string

// Be able to change the value that splits the path
var charToSplitPath string = "."

// Change the value you want to split paths
func ChangePathSplitChar(newChar string) {
	charToSplitPath = newChar
	fmt.Println("New path split of ", newChar)
}

func (mv *Map) SetValue(key string, value any) {
	mvTemp := *mv
	mvTemp[key] = value
}
func getPath(path any) ([]string, error) {
	var fullPath Path
	switch v := path.(type) {
	case []string:
		fullPath = v
	case string:
		fullPath = strings.Split(v, charToSplitPath)
	default:
		return nil, fmt.Errorf("Wrong type for path, got: %s", v)
	}
	return fullPath, nil
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
	return fmt.Errorf("[%s] -> key = '%s' at spot %d not in %v", errorString, path[level], level, currentMap)
}

func (mv Map) Json() ([]byte, error) {
	returnBytes, err := json.Marshal(mv)
	if err != nil {
		return nil, err
	}
	return returnBytes, nil
}
