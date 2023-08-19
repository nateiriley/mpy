package mpy

// index of all paths
var allPaths [][]string

// Will search the map and return the first value
func (mv Map) GetValueFromKey(key string) (any, error) {
	getAllPaths(mv, nil)
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
	getAllPaths(mv, nil)
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

func getAllPaths(m map[string]any, previousKeys []string) []string {
	for k, v := range m {
		if vMap, ok := v.(map[string]any); ok {
			previousKeys = append(previousKeys, k)
			previousKeys = getAllPaths(vMap, previousKeys)
		} else {

			previousKeys = append(previousKeys, k)
			thing := getSliceCopy(previousKeys)
			allPaths = append(allPaths, thing)
			previousKeys = previousKeys[:len(previousKeys)-1]
		}
	}
	return previousKeys
}

func getSliceCopy(input []string) []string {
	var newThing []string
	newThing = append(newThing, input...)
	return newThing
}
