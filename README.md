# Go Inventory API

## Overview

Go Inventory API is an open-source RESTful API designed for efficient inventory management. It provides functionalities for managing suppliers, products, and categories. Leveraging the power of Go, Gorm, and Gin, it offers a robust backend solution for stock level monitoring and CRUD operations on inventory data.

## Tech Stack

- **Go**: The primary programming language used for developing the API.
- **Gin**: A web framework used for routing and middleware.
- **Gorm**: An ORM library for Go, used for database operations.
- **Database**: I recommend using [Neon.tech's](https://neon.tech/) managed PostgreSQL database for storing inventory data.

## Features

- Stock checking for specific products.
- Listing out-of-stock and low-stock products.
- CRUD operations for suppliers, categories, and products.

## API Endpoints

### Suppliers

- **List all suppliers**: `GET /api/v1/suppliers`
- **Create a new supplier**: `POST /api/v1/suppliers`
- **Get a supplier by ID**: `GET /api/v1/suppliers/:id`
- **Update a supplier by ID**: `PUT /api/v1/suppliers/:id`
- **Delete a supplier by ID**: `DELETE /api/v1/suppliers/:id`
- **Get products by supplier ID**: `GET /api/v1/suppliers/products/:supplierId`

### Categories

- **List all categories**: `GET /api/v1/categories`
- **Create a new category**: `POST /api/v1/categories`
- **Get a category by ID**: `GET /api/v1/categories/:id`
- **Update a category by ID**: `PUT /api/v1/categories/:id`
- **Delete a category by ID**: `DELETE /api/v1/categories/:id`

### Products

- **List all products**: `GET /api/v1/products`
- **Create a new product**: `POST /api/v1/products`
- **Get a product by ID**: `GET /api/v1/products/:id`
- **Update a product by ID**: `PUT /api/v1/products/:id`
- **Delete a product by ID**: `DELETE /api/v1/products/:id`
- **Check stock for a product**: `GET /api/v1/suppliers/{supplierId}/products/{productId}/stock`
- **List out-of-stock products**: `GET /api/v1/products/out-of-stock`
- **List low-stock products**: `GET /api/v1/products/low-stock`

## Getting Started

These instructions will get you a copy of the project up and running on your local machine or in a Docker container for development and testing purposes.

### Prerequisites

- Go (Version 1.22)
- Docker (if running the application in a Docker container)
- Gorm
- Gin

### Installing and Running Locally

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/russelldias98/inventory-api-go.git
   ```

### Running with Docker

To containerize the Go Inventory API, follow these steps using Docker:

1. **Build the Docker image**:
   This step compiles your application into a Docker image. Replace `<your-image-name>` with your preferred image name.

   ```bash
   docker build -t <your-image-name> .
   ```

2. **Run the Docker container**:
   This step compiles your application into a Docker image. Replace `<your-image-name>` with your preferred image name.

```bash
docker run -p 3000:3000 <your-image-name> .
```

After executing these steps, your Go Inventory API should be running inside a Docker container and accessible via <http://localhost:3000>.

### Contributing

We welcome contributions to make this project even better. There are several ways you can contribute:

- Reporting bugs.
- Suggesting enhancements.
- Submitting pull requests.

### Acknowledgments

- Gorm for the ORM support.
- Gin for the web framework.
- All contributors and supporters of the project.
