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
can export from selftest/*.postman.json for postman collection

### self test
#### positive case
init wallet
<img width="1333" alt="image" src="https://github.com/user-attachments/assets/21e8eacf-f12d-44e5-92de-8e057f738d44" />
enable wallet
<img width="1352" alt="image" src="https://github.com/user-attachments/assets/8823b96f-2e17-46b2-a4b8-c82bbb82a809" />
get ballance
<img width="1356" alt="image" src="https://github.com/user-attachments/assets/7199213b-4bde-40f2-a493-52298f9f3207" />
all transactions
<img width="1333" alt="image" src="https://github.com/user-attachments/assets/dca0578c-9e5f-4eb9-91a3-e110e213bf0b" />
deposit
<img width="1346" alt="image" src="https://github.com/user-attachments/assets/fae76781-2b4b-4728-905a-6c05893e7150" />
withdraw
<img width="1362" alt="image" src="https://github.com/user-attachments/assets/5769480c-a5fb-4509-a8ac-5437f29a43d8" />
diable wallet
<img width="1352" alt="image" src="https://github.com/user-attachments/assets/b64f34d7-6975-4d17-9def-4c56209a987f" />



#### negative case
wallet diabled
<img width="1354" alt="image" src="https://github.com/user-attachments/assets/d4c8550b-6a19-462d-b7e0-4e2f67b79f3b" />
unauth
<img width="1337" alt="image" src="https://github.com/user-attachments/assets/3eb58c08-69b6-4653-b369-bbffc1c9571f" />






### Authentication

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
