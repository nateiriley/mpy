package mpy

// Will search the map and return the first value
func (mv Map) GetValueFromKey(key string) (any, error) {
	var allPaths [][]string
	_, allPaths = getAllPaths(mv, nil, allPaths)
	var shortestKey []string
	var shortestDistance int = 100
	for _, s := range allPaths {
		for n, s2 := range s {
			if s2 == key {
				if n < shortestDistance {
					shortestDistance = n
					shortestKey = s
				}
			}
		}
	}
	finalValue, err := mv.GetValueFromPath(shortestKey)
	if err != nil {
		return nil, err
	}
	return finalValue, nil
}

// Given a key this will return all values that match with that key
func (mv Map) GetValuesFromKey(key string) ([]any, error) {
	var allPaths [][]string
	_, allPaths = getAllPaths(mv, nil, allPaths)
	var allKeys [][]string
	for _, s := range allPaths {
		for _, s2 := range s {
			if s2 == key {
				allKeys = append(allKeys, s)
			}
		}
	}
	var allValues []any
	for _, path := range allKeys {
		newValue, err := mv.GetValueFromPath(path)
		if err != nil {
			return nil, err
		}
		allValues = append(allValues, newValue)
	}
	return allValues, nil
}
