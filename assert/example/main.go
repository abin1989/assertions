package main

import (
	"fmt"

	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	exampleUsage(assert.So(1, should.Equal, 1)) // pass
	exampleUsage(assert.So(1, should.Equal, 2)) // fail
}

func exampleUsage(result *assert.Result) {
	if result.Failed() {
		fmt.Println("The assertion failed.")
	} else if result.Passed() {
		fmt.Println("The assertion passed.")
	}

	fmt.Print("\nAbout to see result.Err()...\n\n")

	if err := result.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Print("\nAbout to see result.Println()...\n\n")

	result.Println()

	fmt.Print("\nAbout to see result.Log()...\n\n")

	result.Log()

	fmt.Print("\nAbout to see result.Panic()...\n\n")

	defer func() {
		recover()

		fmt.Print("\nAbout to see result.Fatal()...\n\n")

		result.Fatal()

		fmt.Print("---------------------------------------------------------------\n\n")
	}()

	result.Panic()
}
