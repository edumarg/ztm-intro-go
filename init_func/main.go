package main

import (
	"fmt"
	"regexp"
)

var EmailExpr *regexp.Regexp

// will run before main func
// each package can have it's own init func
func init() {
	compiled, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("Failed to compile RegExp")
	}
	EmailExpr = compiled
	fmt.Println("RegExp compile successfully")

}
func isValidEmail(addr string) bool {
	return EmailExpr.Match([]byte(addr))
}

func main() {
	fmt.Println(isValidEmail("invalid"))
	fmt.Println(isValidEmail("valid@email.com"))
	fmt.Println(isValidEmail("invalid@mail"))
}
