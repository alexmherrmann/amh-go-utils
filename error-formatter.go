package amh_go_utils

import (
	"errors"
	"strings"
)

//DeadSimpleFormatter unwraps and formats a chain of errors
func DeadSimpleFormatter(input error) string {
	built := strings.Builder{}
	idx := 1
	for current := input; idx < 32 && current != nil; current, idx = errors.Unwrap(current), idx+1 {
		if idx > 1 {
			built.WriteString("\n")
		}
		// Write the stars
		built.WriteString(strings.Repeat("*", idx))
		built.WriteString(" ")

		// Write the message
		//built.WriteString(reflect.TypeOf(input).Name())
		built.WriteString(current.Error())

	}
	if idx == 32 {
		panic("Too many loop runs!")
	}

	return built.String()
}
