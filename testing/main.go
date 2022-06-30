package main

import "regexp"

func IsValidEmail(addr string) bool {
	compiled, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("Failed to compile RegExp")
	} else {
		return compiled.Match([]byte(addr))
	}
}

func main() {

}
