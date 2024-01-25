# Rest API - Golang Project

This is a sample RESTful API project built using Golang, featuring book categories and user entities. The application supports CRUD operations for managing categories, users, and books. The project is dockerized and utilizes PostgreSQL as its database.

## Table of Contents

- [Rest API - Golang Project with Docker Compose](#rest-api---golang-project-with-docker-compose)
  * [Table of Contents](#table-of-contents)
  * [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
  * [Category Endpoints](#category-endpoints)
  * [User Endpoints](#user-endpoints)
  * [Book Endpoints](#book-endpoints)
  * [Swagger Implementation](#swagger-implementation)
- [License](#license)


## Getting Started

To get started with this project, follow the steps below:

1. Clone the repository:

     ```bash
     git clone https://github.com/your-username/rest_api_sample.git
     cd rest_api_sample
   
2. Set up your database:
    PostgreSQL database and update the database configuration in the config/config.go file.
3. Dockerize the API:

    ```bash
    docker build -t my-api
  
4. Compose up the docker-compose.yml
   
    ```bash
    docker-compose up -d
  

# API Documentation

## Category Endpoints

- **GET /category**: Retrieve all categories.
- **GET /category/{id}**: Retrieve a specific category by ID.
- **POST /category**: Create a new category.
- **PUT /category/{id}**: Update a category by ID.
- **DELETE /category/{id}**: Delete a category by ID.

## User Endpoints

- **GET /user**: Retrieve all users.
- **GET /user/{id}**: Retrieve a specific user by ID.
- **POST /user**: Create a new user.
- **PUT /user/{id}**: Update a user by ID.
- **DELETE /user/{id}**: Delete a user by ID.

## Book Endpoints

- **GET /book**: Retrieve all books.
- **GET /book/{id}**: Retrieve a specific book by ID.
- **POST /book**: Create a new book.
- **PUT /book/{id}**: Update a book by ID.
- **DELETE /book/{id}**: Delete a book by ID.

## Swagger Implementation

The API documentation is available using Swagger. Visit `/swagger/index.html` after running the application to explore and interact with the API.
The application will be accessible at localhost.

# License
This project is licensed under the MIT License. Feel free to use and modify the code as needed.

