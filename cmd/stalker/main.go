package main

import (
	"encoding/json"
	"fmt"

	"github.com/HMasataka/stalker"
)

func main() {
	frame := stalker.NewFrame()

	frame = stalker.Wrap(frame)

	f, _ := json.Marshal(frame)
	fmt.Println("JSON Frame: ", string(f))
}
