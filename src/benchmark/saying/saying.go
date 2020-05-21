package saying

import "fmt"

// Greet welcomes you.
func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
}
