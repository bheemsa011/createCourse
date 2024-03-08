# Course API

This API provides endpoints to manage courses.

## Endpoints

### Home

- **URL:**  http://localhost:4000/
- **Method:** GET
- **Description:** Returns a welcome message.

### Create Course

- **URL:** http://localhost:4000//create-course
- **Method:** POST
- **Description:** Creates a new course.
- **Request Body:** JSON object representing the course to be created.
  - Example:
    ```json
    {
        "CourseId": 43,
        "CourseName": "OS",
        "CoursePrice": 9999,
        "Author": {
            "Fullname": "Ajay",
            "Website": "osbyajay.com"
        }
    }
    ```
- **Response:** JSON object representing the created course.

### Get All Courses

- **URL:** http://localhost:4000/get-all
- **Method:** GET
- **Description:** Retrieves all courses.
- **Response:** JSON array containing all courses.

## Usage

1. Start the server:
    ```bash
    go run main.go
    ```

2. Access the endpoints using appropriate HTTP methods and paths.

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router used for defining API routes.
