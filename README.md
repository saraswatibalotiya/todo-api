# TODO API

This is a TODO API built with Golang and ScyllaDB that supports basic CRUD operations and includes pagination functionality.

## Requirements

- Go 1.16 or higher
- ScyllaDB

## Setup

1. Clone the repository.
2. Install Go and ScyllaDB.
3. Set up ScyllaDB with the provided keyspace and table schema.
4. Run `go mod tidy` to install dependencies.
5. Run the application with `go run main.go`.

## Endpoints

- `POST /todo`: Create a new TODO item.
- `GET /todo`: Get a paginated list of TODO items.
- `GET /todo/{id}`: Get a TODO item by ID.
- `PUT /todo/{id}`: Update a TODO item by ID.
- `DELETE /todo/{id}`: Delete a TODO item by ID.

## Query Parameters

- `user_id`: Filter TODO items by user ID.
- `status`: Filter TODO items by status (e.g., pending, completed).
- `page`: Page number for pagination (default: 1).
- `limit`: Number of items per page (default: 10).

- Pagination is implemented to handle large datasets and improve performance.
