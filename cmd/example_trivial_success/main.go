package main

import "github.com/gontract/gontract"

func exampleFunc() string {

	gontract.Condition(true, gontract.KindPre, "trivially true")

	gontract.Condition(true, gontract.KindPost, "trivially true")
	return "OK"
}

func main() {
	message := exampleFunc()

	println(message)

}
