test:
	oj test --command "go run main.go" --directory tests --tle 2 --jobs 10

fmt:
	gostrip -w ./...
	deadcode -json . | rdeadcode
	go fmt
