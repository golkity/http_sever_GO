# Go HTTP Service with Middleware

This is a simple Go-based HTTP service that demonstrates the use of middleware for sanitization, setting default values, and basic routing. It includes handlers for greeting users and ensures certain parameters (like the `name` query parameter) are sanitized before processing.

## Features
- **Sanitization**: Ensures the `name` query parameter contains only alphabetic characters.
- **Default Value**: Sets a default name as "stranger" if no `name` parameter is provided.
- **Context Usage**: Uses Go's `context` package to pass the `name` value between middlewares and the final handler.
- **RPC Handler**: A placeholder middleware demonstrating how request handling can be extended for RPC-like functionality.

## Endpoints

- **GET /**: Returns a greeting message with the `name` parameter from the query string. If no `name` is provided, it defaults to "stranger".

  Example requests:
    - `GET /?name=John` → `{ "status": "ok", "result": { "greetings": "hello", "name": "John" } }`
    - `GET /` → `{ "status": "ok", "result": { "greetings": "hello", "name": "stranger" } }`
    - `GET /?name=John123` → `{ "status": "error", "result": {} }` (Invalid name, contains non-alphabet characters)

## How it Works

### Middleware Functions
1. **Sanitize**: Validates the `name` query parameter to ensure it only contains alphabetic characters. If the `name` contains non-alphabetic characters, it returns an error response.
2. **RPC**: A placeholder middleware for extending future RPC functionalities.
3. **SetDefaultName**: If the `name` query parameter is empty, this middleware sets the default name to "stranger" and passes it along the request context.

### Handlers
- **StrangerHandler**: The final handler that retrieves the `name` from the request context and sends a JSON response with a greeting.
- **HelloHandler**: The entry point that chains together the middlewares in the correct order.

### Example Flow

1. **Request**: `GET /?name=Alice`
    - The `name` parameter passes through the `Sanitize` middleware (valid name).
    - The `SetDefaultName` middleware is skipped because `name` is provided.
    - The `StrangerHandler` is invoked with the name "Alice", and the response is `{ "status": "ok", "result": { "greetings": "hello", "name": "Alice" } }`.

2. **Request**: `GET /`
    - The `name` parameter is empty.
    - The `SetDefaultName` middleware assigns "stranger" to the name.
    - The `StrangerHandler` is invoked with the name "stranger", and the response is `{ "status": "ok", "result": { "greetings": "hello", "name": "stranger" } }`.

3. **Request**: `GET /?name=John123`
    - The `name` parameter contains invalid characters.
    - The `Sanitize` middleware rejects the request and returns an error response `{ "status": "error", "result": {} }`.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go-http-service.git
   cd go-http-service

2. Install dependencies (if any):
    ```bash
    go mod tidy
    ```
3. Run the application:
    ```bash
    go run main.go

    ```
4. The server will start on http://localhost:8080.You can now test the service using a browser or a tool like curl or Postman.

# Example Usage

- Valid request:
    ```bash
    curl "http://localhost:8080/?name=Alice"
    ```
    Response:
    ```json
    {
      "status": "ok",
      "result": {
        "greetings": "hello",
        "name": "Alice"
      }
    }
    ```
- Invalid request (non-alphabetic characters in name):
    ```bash
    curl "http://localhost:8080/?name=John123"
    ```
    Response:
    ```json
    {
       "status": "error",
       "result": {}
    }
    ```
