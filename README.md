# Movie CRUD API

This is a simple RESTful API written in Go for managing movie data. It provides basic CRUD (Create, Read, Update, Delete) operations for movie resources.

## Getting Started

To get started with this project, follow the steps below:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/e-aub/Movie-CRUD-API-With-golang
   ```
2. **Navigate to the project directory:**
    ```bash
    cd Movie-CRUD-API-With-golang
    ```
3. **Run the application:**
    ```bash
    go run main.go
    ```
4. **Access the API endpoints:**
- GET /movies: Retrieve all movies.
- GET /movies/{id}: Retrieve a specific movie by ID.
- POST /movies: Create a new movie.
- PUT /movies/{id}: Update an existing movie.
- DELETE /movies/{id}: Delete a movie.

##Dependencies
This project uses the following dependencies:

- [gorilla/mux](https://github.com/gorilla/mux): A powerful HTTP router and URL matcher for building Go web servers.
Make sure to have the dependencies installed before running the application.
