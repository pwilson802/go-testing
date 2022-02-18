package messages

import (
	"fmt"
	"strings"
)

func Greet(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	return fmt.Sprintf("Hello, %v\n", strings.Join(names, ", "))
}

func Depart(name string) string {
	return fmt.Sprintf("Goodbye, %v\n", name)
}
