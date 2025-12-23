# Makefiles need real tabs before commands, not expanded tabs, i.e. whitespace/spaces
run:
	@go run cmd/main.go

watch:
	@wgo run cmd/main.go

# May leave behind a process in the background, bring up with `fg` then `Ctrl-C` to kill
regen:
	@go tool templ generate -watch --proxy "http://localhost:3000" --cmd "go run cmd/main.go"
