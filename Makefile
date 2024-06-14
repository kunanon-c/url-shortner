compile:
	go mod tidy -compat=1.22
	mkdir -p out/
	go build -o out/app cmd/main.go

run:
	go run cmd/main.go