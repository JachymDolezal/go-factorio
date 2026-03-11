package main

import (
	"fmt"
	"log"
	"github.com/gorcon/rcon"
)

func main() {

	fmt.Println("Attempting to connect to factorio...")
	conn, err := rcon.Dial("127.0.0.1:27015", "secretpassword")
		if err != nil {
			log.Fatalf("Failed to connect to Factorio: %v\n(Make sure the game is running and a save is loaded!)", err)
		}
		defer conn.Close()
		fmt.Println("Success! Connected to Factorio RCON.")

		// 2. Define our Factorio Lua command
		// The "/c" tells Factorio to execute this as a God Mode cheat.
		// "rcon.print()" is REQUIRED in Factorio to send text back to the Go program.
		luaCmd := `/c rcon.print("Hello from Golang! The factory must grow.")`

		fmt.Printf("\nSending command: %s\n", luaCmd)

		// 3. Execute the command and wait for the response
		response, err := conn.Execute(luaCmd)
		if err != nil {
			log.Fatalf("Failed to execute command: %v", err)
		}

		// 4. Print the response we got back from the game engine
		fmt.Printf("Factorio Server Responded: %s\n", response)


		luaCmd = `/c local p = game.players[1]; if p then p.surface.create_entity{name="stone-furnace", position={x=p.position.x, y=p.position.y - 2}, force=p.force}; rcon.print("Spawned furnace on " .. p.name .. " at " .. math.floor(p.position.x) .. "," .. math.floor(p.position.y)) else rcon.print("Player not found") end`

		fmt.Println("Sending spawn command...")
		response, err = conn.Execute(luaCmd)
		if err != nil {
			log.Fatalf("Command failed: %v", err)
		}

		fmt.Printf("Factorio Server responded: %s\n", response)

}
