# Go Microservice Example

This repository contains a simple microservice implemented in GoLang. The microservice provides CRUD (Create, Read, Update, Delete) operations for managing customer data. It follows a modular structure and employs common design patterns for building scalable and maintainable microservices.

## Project Structure

The project structure is organized as follows:


- **internal/database:** Contains database-related functionalities including database client implementation and database-specific operations.

- **internal/dberrors:** Defines custom error types for handling specific error scenarios during database operations.

- **internal/models:** Defines data models used within the microservice.

- **internal/server:** Implements the HTTP server functionality using the Echo framework, including route handlers and server initialization.

- **main.go:** Entry point of the application. It initializes the database client, creates an instance of the server, and starts the server.

## Design Patterns Used

- **Dependency Injection:** The microservice uses dependency injection to inject dependencies, promoting modularity and testability.

- **Repository Pattern:** Although not explicitly named, the `DatabaseClient` interface serves as a repository interface, abstracting away the details of database access from the server layer.

- **Error Handling:** Custom error types are defined for handling specific error scenarios, enhancing code readability and maintainability.

- **Separation of Concerns:** The project structure separates concerns into different packages, making the codebase modular and easier to manage.

- **HTTP Routing:** The microservice uses the Echo framework for HTTP routing and request handling, providing a clean and organized way to define API endpoints.

## Usage

To run the microservice, make sure you have GoLang installed on your system. Then, execute the following command:

go run main.go


The microservice will start listening on port `8080` by default.

## Contributing

Contributions are welcome! Please feel free to open issues or submit pull requests for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
