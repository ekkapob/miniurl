package utils

import (
	"errors"
	"strconv"
)

func GetIntFromMap(m map[string]string, key string) (int, error) {
	if v, ok := m[key]; ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i, nil
		}
	}
	return 0, errors.New("unable to find value")
}

func GetInt(value string, defaultValue int) int {
	i, err := strconv.Atoi(value)
	if err == nil && i >= 0 {
		return i
	}
	return defaultValue
}
