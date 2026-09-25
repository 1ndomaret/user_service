# User Service

The User Service is a RESTful API built with Go and the Echo framework, serving as the user management and authentication backend for a blood donation platform. It utilizes PostgreSQL for data persistence through GORM.

## Features

* **User Authentication:** Supports user registration and login. Passwords are encrypted using bcrypt, and authentication sessions are managed via JWT tokens.
* **Donor Profile Management:** Automatically generates a linked donor profile upon user registration. Users can securely view and update their profiles with information such as blood type, city, geographic coordinates, and availability status.
* **Service Integration:** Exposes dedicated, service-to-service endpoints protected by a static `SERVICE_TOKEN`. This allows other microservices to search for eligible donors by blood type and city, or retrieve specific donor profiles by their UUID.

## Tech Stack

* **Language:** Go (1.26.4).
* **Web Framework:** Echo v5.
* **Database & ORM:** PostgreSQL managed via GORM.
* **API Documentation:** Swaggo (Swagger UI).
* **Containerization:** Docker.

## Environment Variables

To configure the application locally, create a `secret.env` file in the root directory based on the provided `secret.env.example`. The application requires the following variables:

* `DB_HOST`: The database host address.
* `DB_PORT`: The database port.
* `DB_USER`: The database username.
* `DB_PASSWORD`: The database password.
* `DB_NAME`: The database name.
* `JWT_SECRET`: The secret key used for signing and validating JWT bearer tokens.
* `SERVICE_TOKEN`: The authorization token used for internal service-to-service communication.

## Running the Application

### Using Docker

The project includes a multi-stage `Dockerfile` that builds a lightweight Alpine Linux container. 

```bash
docker build -t user_service .
docker run -p 1323:1323 --env-file secret.env user_service