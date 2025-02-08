# E-Procurement System

This is a procurement system built with Go, Gin, and GORM.

## Prerequisites

- Go 1.22 or higher
- PostgreSQL
- Git

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/febriaricandra/e-procurement-golang.git
cd e-procurement-golang
```
## Setting up ENV

```sh
cp .env .env.example
```
## Install dependecies
```sh
go mod tidy
```

## Run with makefile
```
make run
```

## Project Structure
```
my-procurement-system/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── migrations.go
│   ├── controllers/
│   │   └── item_controller.go
│   ├── models/
│   │   ├── item.go
│   │   └── user.go
│   ├── repositories/
│   │   └── item_repository.go
│   ├── routes/
│   │   └── routes.go
│   └── services/
│       └── item_service.go
├── .env
├── go.mod
└── go.sum
```

## API Endpoints

### Items
- GET /api/items - Get all items
- POST /api/items - Create a new item with warehouse update