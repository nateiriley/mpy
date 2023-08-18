package mpy

import (
	"strings"
)

type Path []string

var CharToSplitPath string = "."

// Path can either a single string seperated by a period, or a string slice
// Example for string this.is.a.path would give you the value this|is|a|path
// Example of string slice ["this","is","a","path"] would give the same result
func (mv Map) GetValueForPath(path any) (any, error) {
	var pathToGet Path
	switch v := path.(type) {
	case []string:
		pathToGet = v
	case string:
		pathToGet = strings.Split(v, ".")
	}

	var currentLevel map[string]any
	for i, key := range pathToGet {
		holdValue, ok := mv[key]
		if !ok {
			return nil, getPathError(pathToGet, i)
		}

	}

	return nil, nil
}

func getPathError(path []string, level int) error {
	var errorString string
	for j := 0; j < level; j++ {
		if j == 0 {
			errorString = path[j]
		} else {
			errorString = errorString + "." + path[j]
		}
	}
}
