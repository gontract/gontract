package main

import "github.com/gontract/gontract"

func exampleFunc() string {

	gontract.Require(true, "trivially true")

	gontract.Ensure(true, "trivially true")
	return "OK"
}

func main() {
	message := exampleFunc()

	println(message)

}
