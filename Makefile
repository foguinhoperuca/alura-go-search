build:
	@clear
	@date
	time go build .
	@date

run:
	@clear
	@date
	time go run cmd/main.go
	@date
