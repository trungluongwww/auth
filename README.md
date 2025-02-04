# Base API

## Steps to Run the Project
1. Setup environment variables by copying `.env.example` and configuring the necessary values.
2. Install dependencies using following command:\
   ```go mod tidy```
3. Start the project
   - Using Go: `go run cmd/main.go`
   - Using Makefile: `make run`

## Steps to Generate Database Structures
1. Configure `gen.tool` file with the correct database environment settings
2. Install **Gentool**\
   ```go install gorm.io/gen/tools/gentool@latest```
3. Apply database changes:
   - Add a sql file with new changes in `mysql/ddl` directory
   - Run project first to apply new changes
4. Gen database structures:
   - Using Gentool: `gentool -c gen.tool`
   - Using Makefile: `make gen` if makefile is available
5. Verify the generated structures in `internal/model` directory
