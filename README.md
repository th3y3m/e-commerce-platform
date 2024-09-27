[![CodeQL](https://github.com/th3y3m/e-commerce-platform/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/th3y3m/e-commerce-platform/actions/workflows/github-code-scanning/codeql)
[![E-Commerce Platform (CI applid)](https://github.com/th3y3m/e-commerce-platform/actions/workflows/ci-script.yml/badge.svg)](https://github.com/th3y3m/e-commerce-platform/actions/workflows/ci-script.yml)
![MIT License](https://img.shields.io/badge/License-MIT-yellow.svg)


# Welcome to E-Commerce Platform in Go

This project is an e-commerce platform built using Go, leveraging a range of powerful libraries and services to handle various aspects of the application, including user authentication, payment processing, product management, and more.

---

## Technologies Used

<p align="center">
  <img src="https://www.vectorlogo.zone/logos/golang/golang-official.svg" alt="Go" width="100" height="100"/>
  <img src="https://www.vectorlogo.zone/logos/postgresql/postgresql-icon.svg" alt="PostgreSQL" width="100" height="100"/>
  <img src="https://avatars.githubusercontent.com/u/2370800?s=200&v=4" alt="Gin Gonic" width="100" height="100"/>
  <img src="https://www.vectorlogo.zone/logos/google_cloud/google_cloud-icon.svg" alt="Google Cloud" width="100" height="100"/>
  <img src="https://avatars.githubusercontent.com/u/10409771?s=200&v=4" alt="Casbin" width="100" height="100"/>
  <img src="https://www.vectorlogo.zone/logos/jsonwebtokenio/jsonwebtokenio-icon.svg" alt="JWT" width="100" height="100"/>
  <img src="https://www.vectorlogo.zone/logos/facebook/facebook-icon.svg" alt="Facebook OAuth" width="100" height="100"/>
  <img src="https://www.vectorlogo.zone/logos/google/google-icon.svg" alt="Google OAuth" width="100" height="100"/>
  <img src="https://www.vectorlogo.zone/logos/redis/redis-icon.svg" alt="Redis" width="100" height="100"/>
</p>

---

## Features

- **User Authentication**: Supports Google and Facebook authentication via `goth`, alongside email and JWT authentication.
- **Storage**: Uses Google Cloud Storage for storing product images and other media files.
- **RBAC Authorization**: Implements Role-Based Access Control (RBAC) using `casbin` to manage user permissions efficiently.
- **Task Scheduling**: Utilizes `gocron` for scheduling and managing periodic tasks such as promotions, product updates, etc.
- **Database Management**: Powered by GORM with PostgreSQL as the primary database for storing user data, products, and orders.
- **Session Management**: Uses `gin-contrib/sessions` for secure session handling.
- **Custom Product UUIDs**: Generates unique identifiers for products using `google/uuid`.
- **Configuration Management**: Environment variables are managed using `godotenv` for easier configuration.

## Tech Stack

- **Backend Framework**: `gin` (A lightweight web framework in Go).
- **Database**: PostgreSQL with `gorm` for ORM.
- **Authentication**: `goth` for Google and Facebook OAuth, `jwt-go` for JWT tokens.
- **Storage**: Google Cloud Storage for media assets.
- **Scheduling**: `gocron` for background tasks.
- **RBAC Authorization**: `casbin` for role-based access control.
- **Unique Identifiers**: `google/uuid` for generating UUIDs.
- **Task Queue**: `asynq` for background job processing (indirect dependency).

## Requirements

- Go 1.23.0 or above.
- PostgreSQL database.
- Google Cloud credentials for Cloud Storage.
- OAuth credentials for Google and Facebook authentication.
  
## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/th3y3m/e-commerce-platform.git
    cd e-commerce-platform
    ```

2. **Install Go dependencies**:
    ```bash
    go mod download
    ```

3. **Set up environment variables** by creating a `.env` file in the root directory:
    ```bash
    touch .env
    ```

4. **Run the application**:
    ```bash
    go run main.go
    ```

## Modules & Libraries

- **Gin**: Web framework for building the HTTP server and RESTful APIs.
- **GORM**: ORM library for database management.
- **Casbin**: Authorization library for role-based access control (RBAC).
- **JWT-Go**: For creating and verifying JWT tokens.
- **Google Cloud Storage**: For storing media files.
- **GoCron**: For scheduling tasks.
- **Goth**: For Google and Facebook OAuth.
- **Godotenv**: For environment variable management.
- **UUID**: For generating unique product and order IDs.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


#### Connect me via: truongtanhuy3006@gmail.com

##### &#169; 2024 th3y3m

