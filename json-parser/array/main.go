package main

import (
	"fmt"
	"encoding/json"
)

type Bird struct {
	Name string
  // Description string
}

func main() {
	birdJson := `[{"name":"pigeon"},{"name":"kes"}]`
	var birds []Bird
	err := json.Unmarshal([]byte(birdJson), &birds)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Birds : %+v", birds)
}