# Hospital Management System REST API

This project implements a RESTful API for managing hospital-related operations such as authentication, clinic management, patient management, scheduling, and more.

## Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/wussh/hospital.git
   ```

2. Navigate to the project directory:
   ```bash
   cd hospital
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## Running with Docker Compose

1. Make sure you have Docker and Docker Compose installed on your machine.

2. Create a `.env` file based on the provided `.env.example` file and customize the environment variables as needed.

3. Build and run the Docker containers:
   ```bash
   docker-compose up -d
   ```

4. Access the API at `http://localhost:9000`.

## Shutting Down

To shut down the application and Docker containers, run:
```bash
docker-compose down
```

## Changing Environment Variables

You can customize environment variables by editing the `.env` file. Refer to the provided `.env.example` file for a list of required variables.

## Database Design

The project uses PostgreSQL as the database. You can find the database schema and design in the `src/repository/postgres` directory.

## Project Structure

The project follows a standard Go project structure:
- `src`: Contains all source code files.
  - `delivery/http/echo`: HTTP delivery layer with Echo framework.
  - `domain`: Domain layer containing business logic interfaces and models.
  - `entity`: Entity layer containing database models.
  - `repository/postgres`: Repository layer for PostgreSQL database operations.
  - `service`: Service layer for application-specific logic.
  - `usecase`: Usecase layer for orchestrating business logic.
  - `util`: Utility functions and helpers.
- `docker-compose.yml`: Docker Compose configuration file.
- `Dockerfile`: Dockerfile for building the application image.
- `go.mod` and `go.sum`: Go module files.
- `Makefile`: Makefile for common tasks.
- `.env.example`: Example environment variable file.
- `main.go`: Entry point of the application.

## Routes

### Authentication

- `POST /login`: User login.
- `POST /logout`: User logout (requires JWT token).
- `PUT /authentications`: Update authentication information (requires JWT token).

### Clinic

- `POST /clinics`: Add a new clinic (requires JWT token).
- `GET /clinics`: Get all clinics (requires JWT token).
- `GET /clinics/:clinicID`: Get clinic by ID (requires JWT token).
- `PUT /clinics/:clinicID`: Update clinic by ID (requires JWT token).
- `DELETE /clinics/:clinicID`: Delete clinic by ID (requires JWT token).

### Doctor

- `POST /doctors`: Add a new doctor (requires JWT token).
- `GET /doctors`: Get all doctors (requires JWT token).
- `GET /doctors/:doctorID`: Get doctor by ID (requires JWT token).
- `PUT /doctors/:doctorID`: Update doctor by ID (requires JWT token).
- `DELETE /doctors/:doctorID`: Delete doctor by ID (requires JWT token).

### Hello

- `GET /hello`: Hello World endpoint.

### Medical Record

- `POST /sessions/:sessionID/records`: Add a new medical record to a session (requires JWT token).

### Patient

- `POST /patients`: Add a new patient (requires JWT token).
- `GET /patients`: Get all patients (requires JWT token).
- `GET /patients/:patientID`: Get patient by ID (requires JWT token).
- `PUT /patients/:patientID`: Update patient by ID (requires JWT token).
- `DELETE /patients/:patientID`: Delete patient by ID (requires JWT token).
- `GET /nik/:nik`: Get patient by National Identification Number (requires JWT token).

### Schedule

- `POST /users/:userID/schedules`: Add a new schedule for a user (requires JWT token).
- `GET /users/:userID/schedules`: Get all schedules for a user (requires JWT token).
- `GET /users/:userID/schedules/:scheduleID`: Get schedule by ID for a user (requires JWT token).
- `PUT /users/:userID/schedules/:scheduleID`: Update schedule by ID for a user (requires JWT token).
- `DELETE /users/:userID/schedules/:scheduleID`: Delete schedule by ID for a user (requires JWT token).

### Session

- `POST /sessions`: Add a new session (requires JWT token).
- `GET /sessions`: Get all sessions (requires JWT token).
- `POST /sessions/:sessionID/complete`: Complete a session (requires JWT token).
- `POST /sessions/:sessionID/cancel`: Cancel a session (requires JWT token).
- `POST /sessions/:sessionID/activate`: Activate a session (requires JWT token).

### Staff

- `POST /staffs`: Add a new staff member.

### Static Files

- `GET /avatar/:avatar`: Serve avatar images.

### User

- `POST /users`: Add a new user (requires JWT token).
- `PUT /users/:userID/avatar`: Update user avatar (requires JWT token).
- `DELETE /users/:userID/avatar`: Delete user avatar (requires JWT token).
- `GET /users/:userID`: Get user by ID (requires JWT token).
- `GET /users/current`: Get currently authenticated user (requires JWT token).