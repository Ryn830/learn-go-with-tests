package hello

import (
	"fmt"
)

func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return fmt.Sprintf("Hello, %s", name)
}
