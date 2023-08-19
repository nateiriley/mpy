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
		return nil, fmt.Errorf("wrong type for path, got: %s", v)
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

func getAllPaths(m map[string]any, previousKeys []string, allPaths [][]string) ([]string, [][]string) {
	for k, v := range m {
		if vMap, ok := v.(map[string]any); ok {
			previousKeys = append(previousKeys, k)
			previousKeys, allPaths = getAllPaths(vMap, previousKeys, allPaths)
		} else {

			previousKeys = append(previousKeys, k)
			thing := getSliceCopy(previousKeys)
			allPaths = append(allPaths, thing)
			previousKeys = previousKeys[:len(previousKeys)-1]
		}
	}
	return previousKeys, allPaths
}

func getSliceCopy(input []string) []string {
	var newThing []string
	newThing = append(newThing, input...)
	return newThing
}

func compareSlice(a, b []string) error {
	for i, aa := range a {
		bb := b[i]
		if aa != bb {
			return fmt.Errorf("%s != %s", aa, bb)
		}
	}
	return nil
}

func sliceInSlice(stringSlice, subStringSlice []string) (distance int, inside bool) {
	if len(subStringSlice) > len(stringSlice) {
		return 0, false
	}
	if err := compareSlice(stringSlice, subStringSlice); err == nil {
		return 0, true
	}
	for n := 0; n < len(stringSlice)-len(subStringSlice)+1; n++ {
		if stringSlice[n] == subStringSlice[0] {
			for i := 0; i < len(subStringSlice); i++ {
				if stringSlice[n+i] != subStringSlice[i] {
					break
				}
				if i == len(subStringSlice)-1 {
					return n, true
				}
			}
		}
	}
	return 0, false
}
