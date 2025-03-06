# fiber-generic
A simple Generic Golang App using Go-Fiber Framework

## Features
- CRUD operations for User model
- Uses GORM for database interactions
- RESTful API with Go-Fiber

## Requirements
- Go 1.22.3 or higher
- PostgreSQL

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/fiber-generic.git
    ```
2. Navigate to the project directory:
    ```sh
    cd fiber-generic
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Configuration
Update the database connection string in `:
[main.go](http://_vscodecontentref_/1)``go
dsn := "host=localhost user=postgres password=yourpassword dbname=gin_generic port=5432 sslmode=disable TimeZone=Asia/Jakarta"