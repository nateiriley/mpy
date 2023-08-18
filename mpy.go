package mpy

type Map map[string]any

// Make a new map
func New() Map {
	m := make(map[string]any, 0)
	return m
}

// Make a copy of a map
func (mv Map) Copy() Map {
	var newMap Map
	for k, v := range mv {
		newMap[k] = v
	}
	return newMap
}
