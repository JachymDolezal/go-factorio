# GO factorio

hobby project to learn go + create a functional AI agent that attempts to play Factorio


## Setup

To start the headless server
```bash
./bin/x64/factorio --start-server ./clean_world.zip --bind 0.0.0.0 --rcon-port 27015 --rcon-password secretpassword
```


## TODO

- [ ] create llm provider
- [ ] add mcp support
- [ ] basic agentic loop
- [ ] create a basic mcp server with 'hardcoded' lua scripts
- [ ] download the whole rcon API json, parse to go struct, create tool or utility for semantic search over the API
- [ ] create advanced tools with better agentic loop
