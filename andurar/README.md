# Booking System API

Welcome to the Booking System API! This project is designed to provide a robust and scalable backend for a booking system, leveraging the GoFiber web framework and adhering to Clean Architecture principles. The entire system is containerized using Docker for ease of deployment and management.

# Table of Contents

Features
Technologies Used
Architecture
Installation
Usage
API Endpoints

# Features

User registration and authentication
Hotel management (CRUD operations)
Room booking and management
Booking management
Review and rating system

# Technologies Used

GoFiber: An Express-inspired web framework for Go, known for its performance and minimal footprint.
Docker: Container platform to build, ship, and run distributed applications.
Mysql: A powerful, open-source object-relational database system.

# Architecture

This project follows the principles of Clean Architecture, ensuring that the system is:

Independent of frameworks: The architecture does not depend on the existence of some library of feature-laden software.
Testable: The business rules can be tested without the UI, database, web server, or any other external element.
Independent of UI: The UI can change easily, without changing the rest of the system.
Independent of Database: The business rules are not bound to the database.
Independent of any external agency: Business rules simply do not know anything at all about the outside world.

# Installation

# Prerequisites

Docker and Docker Compose installed on your machine
Makefile (optional, for simplified command execution)

# change directory into the tracker directory to run the docker

cd booking-system-api
Build and Run with Docker

docker-compose up --build

# Access the API

API: http://localhost:3007
ENV: contains more information about the env variables you might need to know

# Usage

Running the Development Server
To start the GoFiber development server using Docker, you can use the Makefile if it is installed:

make watch
Alternatively, you can directly use the Docker Compose command:

docker-compose up --build

# Building and Running the Project with Docker

Ensure that Docker is running on your machine. Use the following command to build and run the containers:

docker-compose up --build

# API Endpoints

# User

POST /api/users/register: Register a new user
POST /api/users/login: User login

# Hotels

GET /api/hotels: Get all hotels
POST /api/hotels: Create a new hotel
GET /api/hotels/
: Get a hotel by ID
PUT /api/hotels/
: Update a hotel by ID
DELETE /api/hotels/
: Delete a hotel by ID

# Rooms

GET /api/rooms: Get all rooms
POST /api/rooms: Create a new room
GET /api/rooms/
: Get a room by ID
PUT /api/rooms/
: Update a room by ID
DELETE /api/rooms/
: Delete a room by ID

# Bookings

GET /api/bookings: Get all bookings
POST /api/bookings: Create a new booking
GET /api/bookings/
: Get a booking by ID
PUT /api/bookings/
: Update a booking by ID
DELETE /api/bookings/
: Cancel a booking by ID

# Reviews

POST /api/reviews: Submit a review
GET /api/reviews/
: Get reviews for a hotel
