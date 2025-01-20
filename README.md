# Delivery Management System

A web application for managing deliveries. This project includes a frontend built with React (TypeScript and ShadCN UI) and a RESTful backend API developer in Golang.

---

## Summary

- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Environment Configuration](#environment-configuration)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
- [License](#license)

---

## Getting Started

Follow these steps to set up and run the project on your local machine.

---

## Prerequisites

Make sure you have the following installed:

- [Docker](https://www.docker.com/get-started)
- [Golang](https://go.dev/doc/tutorial/getting-started)

---

## Installation

1. Clone this repository:

   ```bash
   git clone git@github.com:benevidesjohns/client-geolocation-api.git
   cd client-geolocation-api
   ```

2. Navigate to the `frontend` or `backend` folder as needed.

---

## Environment Configuration

1. Copy the `.env.example` file and rename it to `.env` in the `root of the project` and inside the `backend` directory:

   ```bash
   cp .env.example .env
   ```

2. Update the variables in the `.env` file according to your environment. Example:

   For the backend:

   ```env
   DB_HOST=mysql
   DB_PORT=3306
   DB_USER=user
   DB_PASSWORD=password
   DB_NAME=my-database
   ```

---

## Usage

1. Build and start the Docker containers:

   ```bash
   docker compose up -d --build
   ```

2. Access the containers:

   - Frontend:
     ```bash
     docker exec -it go-front bash
     ```
     Now run the following commands:

     ```bash
     npm install
     ```
     To install the required packages (node_modules).
      
     ```bash
     npm run dev
     ```
     To run the front-end application.

   - Backend:
     ```bash
     docker exec -it go-api sh
     ```

     Navigate to the go-api container and start the server using the command:
     ```bash
     air
     ```

   - MySQL:
     ```bash
     docker exec -it go-mysql sh
     ```

3. Open the frontend in your browser:

   ```
   http://localhost:3000
   ```

4. Backend API will be available at:

   ```
   http://localhost:8080/api/deliveries
   ```

---

## Project Structure

```plaintext
.
├── frontend         # React (TypeScript + ShadCN UI) application
├── backend          # Golang (HTTP Mux)
├── .env.example     # Environment variables example
├── README.md        # Project documentation
├── docker-compose.yml
```

---

## API Endpoints

### Delivery Endpoints

- `GET /deliveries` - Get all deliveries
- `GET /deliveries/:id` - Get a delivery by ID
- `POST /deliveries` - Create a new delivery
- `PUT /deliveries/:id` - Update a delivery
- `DELETE /deliveries/:id` - Delete a delivery
- `DELETE /deliveries` - Delete all delivery

---

### Developed by João Benevides