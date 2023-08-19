package mpy

import (
	"fmt"
	"strings"
)

// Path can either a single string seperated by a period, or a string slice
// Example for string this.is.a.path would give you the value this|is|a|path
// Example of string slice ["this","is","a","path"] would give the same result
// This will return anything
func (mv Map) GetValueFromPath(path any) (any, error) {
	var fullPath Path
	switch v := path.(type) {
	case []string:
		fullPath = v
	case string:
		fullPath = strings.Split(v, charToSplitPath)
	}

	var currentLevel map[string]any
	var returnValue any
	for i, key := range fullPath {
		if i == 0 {
			value, ok := searchLevel(mv, key)
			if !ok {
				return nil, getPathError(mv, fullPath, i)
			}
			if i == len(fullPath)-1 {
				returnValue = value
			} else {
				currentLevel, ok = value.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("%s cannot is not a map, can't go further down [%s]", value, fullPath)
				}
			}
		} else {
			value, ok := searchLevel(currentLevel, key)
			if !ok {
				return nil, getPathError(currentLevel, fullPath, i)
			}
			if i == len(fullPath)-1 {
				returnValue = value
			} else {
				currentLevel, ok = value.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("%s is not a map, can't go further down [%s]", value, fullPath)
				}
			}
		}
	}

	return returnValue, nil
}
