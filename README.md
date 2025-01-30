# Detect Global Environment Variable Declarations in Golang

This script scans a Golang project directory for global environment variable declarations using `os.Getenv` and reports them.

## Features
- Recursively scans all `.go` files in a specified folder.
- Parses each file and detects global variables assigned using `os.Getenv`.
- Prints the detected global environment variable names along with the file path.

## Prerequisites
- Install [Go](https://golang.org/doc/install) (Golang 1.16 or later recommended).

## Installation
1. Clone this repository.
2. Navigate to the folder where the script is located.

## Usage
Run the script with the directory path of your Go project:

```sh
go run main.go /path/to/your/go/project
```

## Example Output
```
Scanning directory: /Users/ahmet/projects/backend
Global env variable detected: readUser in file: /Users/ahmet/projects/backend/config/db.go
Global env variable detected: readPassword in file: /Users/ahmet/projects/backend/config/db.go
Global env variable detected: readHost in file: /Users/ahmet/projects/backend/config/db.go
```

## How It Works
1. **Scans the provided directory** recursively for `.go` files.
2. **Parses each file** using Go's `go/parser` and `go/token` packages.
3. **Checks for global `var` declarations** that use `os.Getenv`.
4. **Outputs detected variables** and their respective file paths.

## Notes
- This script does not detect environment variables declared inside functions.
- Ensure your Go files are properly formatted and parseable.

