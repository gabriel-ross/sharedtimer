server-up:
	go run ./cmd/main.go

tauri-up:
	cargo tauri dev
	
npm-install:
	npm install
