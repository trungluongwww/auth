gen:
	rm ./internal/model/*
	gentool -c ./gen.tool

run:
	go run cmd/main.go