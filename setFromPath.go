package mpy

import (
	"errors"
	"fmt"
	"strings"
)

func (mv *Map) SetValueFromPath(path, setValue any) (any, error) {
	var fullPath Path
	switch v := path.(type) {
	case []string:
		fullPath = v
	case string:
		fullPath = strings.Split(v, charToSplitPath)
	}
	err := mv.setValueForPath(fullPath, setValue)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (mv *Map) setValueForPath(path []string, value any) error {
	if len(path) > 1 {
		mvTemp := *mv
		newMv, ok := mvTemp[path[0]]
		if !ok {
			return errors.New(fmt.Sprint(path[0]) + " not in " + fmt.Sprint(mvTemp))
		}
		newMp, ok := newMv.(map[string]any)
		if !ok {
			return errors.New("couldn't convert newMv to map")
		}
		mkMap := Map(newMp)
		mpPointer := &mkMap
		err := mpPointer.setValueForPath(path[1:], value)
		if err != nil {
			return err
		}
	} else {
		err := mv.SetValue(path[0], value)
		if err != nil {
			return err
		}
	}
	return nil
}
