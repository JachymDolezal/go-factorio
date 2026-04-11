package main

import (
	"context"
	"factorio-ai/src/provider"
	"flag"
	"fmt"
	// "github.com/gorcon/rcon"
)

var model = flag.String("model", "gemini-flash-lite-latest", "the model name, e.g. gemini-flash-lite-latest")

func main() {
	ctx := context.Background()
	// flag.Parse()

	llm := provider.NewLLMProvider(*model, ctx)

	response := llm.GenerateText("Hello")
	fmt.Println(response)
	// fmt.Println("Attempting to connect to factorio...")

	// conn, err := rcon.Dial("127.0.0.1:27015", "secretpassword")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Factorio: %v\n(Make sure the game is running and a save is loaded!)", err)
	// }
	// defer conn.Close()
	// fmt.Println("Success! Connected to Factorio RCON.")

	// 2. Define our Factorio Lua command
	// The "/c" tells Factorio to execute this as a God Mode cheat.
	// "rcon.print()" is REQUIRED in Factorio to send text back to the Go program.

	// 4. Print the response we got back from the game engine
}
