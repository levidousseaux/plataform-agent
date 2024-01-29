run:
	- go run cmd/server/main.go & npm run dev

build:
	- npm install
	- npm run build
	- mkdir -p bin/server/public
	- cp -r web/public/* bin/server/public
	- go build -o bin/server cmd/server/main.go