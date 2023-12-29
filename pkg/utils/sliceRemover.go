package utils

import "reflect"

func Remove(slice []any, is bool, dst interface{}) []any {
	for _, element := range slice {
		tp := reflect.TypeOf(element)
		tp.Name()
	}
	if is {
		j := 0
		for _, val := range slice {
			if val == 0 {
				slice[j] = val
				j++
			}
		}
		return slice[:j]
	}
	return slice
}
