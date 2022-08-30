package amh_go_utils

import "errors"

func FindAllMatching[T any](tosearch []T, searcher func(int, T) bool) []T {
	found := []T{}

	for idx, tocheck := range tosearch {
		if searcher(idx, tocheck) {
			found = append(found, tocheck)
		}
	}

	return found
}

func MapSlice[T any, R any](toMap []T, mapper func(int, T) R) []R {
	newMap := make([]R, len(toMap))

	for idx, valueToMap := range toMap {
		newMap[idx] = mapper(idx, valueToMap)
	}

	return newMap
}

var DifferentLengths = errors.New("The lists are different lengths")

func ZipTwoMap[T1, T2, R any](first []T1, second []T2, mapper func(int, T1, T2) R) ([]R, error) {
	if len(first) != len(second) {
		return nil, DifferentLengths
	}

	newMap := make([]R, len(first))

	for i := 0; i < len(first); i++ {
		firstElement := first[i]
		secondElement := second[i]

		newMap[i] = mapper(i, firstElement, secondElement)
	}

	return newMap, nil

}

// All checks that ALL things in the array match the tester
func All[T any](toTest []T, test func(idx int, input T) bool) bool {
	if len(toTest) == 0 {
		return false
	}
	for i, t := range toTest {
		if !test(i, t) {
			return false
		}
	}

	return true
}

//func DeleteAllMatching
