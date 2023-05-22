npm-dl:
	cd ./client/ui
	npm install

ui-run:
	cd ./client/ui
	npm start

server-run:
	go run ./cmd/main.go

tauri-run:
	cd client/sharedtimer-client
	npm run tauri dev