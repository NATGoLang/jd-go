# Project Name

Welcome to the NAT project! This README provides essential information to get you started.

## Getting Started

1. create a config file in `configs/private_config.yaml` with following content
   ```
   # Database credentials
   database:
     dsn: "<username>:<password>@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
   ```
2. run `go run main.go` to start the server
