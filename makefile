.PHONY: all clean test

# Default target
all: Compile

# Build target
Compile:
	go build -o Cleaner --trimpath -ldflags "-s -w" main.go
	strip -s ./Cleaner
	upx -9 ./Cleaner
	upx -t ./Cleaner

windows:
	GOOS=windows GOARCH=amd64 go build -o Cleaner.exe --trimpath -ldflags "-s -w" main.go
	strip -s ./Cleaner.exe
	upx -9 ./Cleaner.exe
	upx -t ./Cleaner.exe

# Test target
test:
	go run main.go

# Clean target
clean:
	rm -f Cleaner

# Test branch target
test-branch: clean
	git checkout -b test-branch
	make all
