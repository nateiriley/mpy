package mpy

import (
	"errors"
	"fmt"
)

func (mv *Map) SetValueFromPath(path, setValue any) error {
	fullPath, err := getPath(path)
	if err != nil {
		return err
	}
	err = mv.setValueForPath(fullPath, setValue)
	if err != nil {
		return err
	}

	return nil
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
		mv.SetValue(path[0], value)
	}
	return nil
}
