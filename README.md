# MiniWallet

## Overview

MiniWallet is a Go-based project that provides a RESTful API for managing users, categories, products, and carts. The project uses the Gin framework for handling HTTP requests and GORM for database interactions.

## Project Structure

```
MiniWallet/
├── cmd/
│   └── main.go
├── internal/
│   ├── delivery/
│   │   └── http/
│   │       ├── user.go
│   │       ├── wallet.go
│   │       ├── transaction.go
│   │       
│   ├── repository/
│   │   ├── user_repository.go
│   │   ├── wallet_repository.go
│   │   ├── transaction_repository.go
│   │   
│   └── usecase/
│       ├── user_usecase.go
│       ├── wallet_usecase.go
│       ├── transaction_usecase.go
│       
├── middleware/
│   └── auth_middleware.go
├── pkg/
│   └── utils/
│       ├── jwt.go
│       └── response.go
└── go.mod
```

## Getting Started

### Prerequisites

- Go 1.16+
- A running instance of a database (e.g., PostgreSQL)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/MiniWallet.git
   cd MiniWallet
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Set up your database and update the connection details in `pkg/utils/db.go`.

4. Run the application:
   ```sh
   go run cmd/main.go
   ```

## API Documentation

### Authentication

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.