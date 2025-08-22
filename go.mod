module cmd/api/main.go

go 1.24.6

replace internal/interfaces/router => ./internal/interfaces/router

require internal/interfaces/router v1.0.0

require github.com/joho/godotenv v1.5.1 // indirect
