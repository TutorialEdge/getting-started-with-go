package main

import (
	"fmt"
	"getting-started-with-go/internal/api"
)

func main() {
	fmt.Println("I'm officially a gopher!")
	myAPI := api.New()

	if err := myAPI.Start(); err != nil {
		panic(err)
	}
}
