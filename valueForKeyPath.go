package mpy

// Given a key path (doesn't have to be from the top level) will give the closest key path
// from the top level
func (mv Map) GetValueFromKeyPath(keyPath any) (any, error) {
	newKeyPath, err := getPath(keyPath)
	if err != nil {
		return nil, err
	}
	var allPaths [][]string
	_, allPaths = getAllPaths(mv, nil, allPaths)
	var shortestKey []string
	var shortestDistance int = 100
	for _, s := range allPaths {
		distance, inside := sliceInSlice(s, newKeyPath)
		if inside {
			if shortestDistance > distance {
				shortestDistance = distance
				shortestKey = s[:distance+len(newKeyPath)]
			}
		}
	}
	finalValue, err := mv.GetValueFromPath(shortestKey)
	if err != nil {
		return nil, err
	}
	return finalValue, nil
}

func (mv Map) GetValuesFromKeyPath(keyPath any) ([]any, error) {
	newKeyPath, err := getPath(keyPath)
	if err != nil {
		return nil, err
	}
	var allPaths [][]string
	_, allPaths = getAllPaths(mv, nil, allPaths)
	var allKeys [][]string
	for _, s := range allPaths {
		distance, inside := sliceInSlice(s, newKeyPath)
		if inside {
			allKeys = append(allKeys, s[:distance+len(newKeyPath)])
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
