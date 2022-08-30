package amh_go_utils

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name   string
		array  []bool
		result bool
	}{
		{"All true", []bool{true, true, true}, true},
		{"All false", []bool{false, false, false}, false},
		{"Last false", []bool{true, true, false}, false},
		{"None", []bool{}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := All(test.array, func(idx int, i bool) bool { return i })
			if result != test.result {
				t.Errorf("%v != %v", result, test.result)
			}
		})
	}
}

func TestSliceUtils(t *testing.T) {

	t.Run("Increasing idx", func(t *testing.T) {
		toMap := []int{1, 2, 4, 8, 16}
		result := []int{0, 1, 2, 3, 4}

		actualResult := MapSlice(toMap, func(i int, result int) int {
			return i
		})

		if !reflect.DeepEqual(result, actualResult) {
			t.Errorf("%v and %v not equal", result, actualResult)
		}
	})

	t.Run("Integer to string mapping", func(t *testing.T) {
		toMap := []int{1, 2, 4, 8, 16}
		result := []string{"1", "2", "4", "8", "16"}

		actualResult := MapSlice(toMap, func(i int, result int) string {
			return strconv.Itoa(result)
		})

		if !reflect.DeepEqual(result, actualResult) {
			t.Errorf("%v and %v not equal", result, actualResult)
		}

	})

	t.Run("Zip up two things", func(t *testing.T) {
		// Going to check numbers and their square roots
		first := []int{1, 2, 3, 4, 5}
		second := []int{1, 2 * 2, 3 * 3, 4 * 4, 5 * 5}

		zipped, zipErr := ZipTwoMap(first, second, func(i int, t1 int, t2 int) bool {
			return t1*t1 == t2
		})

		if zipErr != nil {
			t.Fatalf("Had an error %v", zipErr.Error())
		}

		for idx, boo := range zipped {
			if !boo {
				t.Fatalf("%v to %v was not computed correctly", first[idx], second[idx])
			}
		}
	})
}
