# Sales Data Application

## Overview

The Sales Data Application is a Golang-based service for managing and analyzing historical sales data. It performs the following key functions:

- **Data Loading:** Imports sales data from a CSV file into a PostgreSQL database.
- **Data Refresh:** Supports periodic and on-demand data refresh.
- **Revenue Analysis:** Provides RESTful API endpoints to retrieve various revenue-related metrics.

## Prerequisites

- **Go** (version 1.18 or higher)
- **PostgreSQL** (installed and running)
- **CSV File**: A CSV file (`sales_data.csv`) with sales data

## Setup

1. **Clone the Repository**

    ```bash
    git clone https://github.com/yourusername/sales-data-app.git
    cd sales-data-app
    ```

2. **Install Dependencies**

    Install the required Go packages:

    ```bash
    go mod tidy
    ```

3. **Set Up PostgreSQL**

    Ensure PostgreSQL is installed and running. Create a database named `sales_data`:

    ```sql
    CREATE DATABASE sales_data;
    ```

    Update the `dsn` (Data Source Name) in `main.go` to match your PostgreSQL setup:

    ```go
    dsn := "host=localhost user=postgres password=yourpassword dbname=sales_data port=5432 sslmode=disable"
    ```

4. **Create Tables**

    The application automatically creates tables on the first run using GORM's `AutoMigrate`. Ensure the `initDB` function in `main.go` is correctly set up with your database connection details.

## Running the Application

1. **Start the Application**

    To run the application, use:

    ```bash
    go run main.go
    ```

    This will start the server on `http://localhost:8080`.

## API Endpoints

### Data Refresh

- **Trigger Data Refresh**
  - **Endpoint:** `POST /refresh-data`
  - **Description:** Initiates the loading of CSV data into the database.
  - **Request Body:** None
  - **Response:**
    ```json
    {
      "message": "Data refresh initiated"
    }
    ```

### Revenue Analysis

- **Total Revenue**
  - **Endpoint:** `GET /revenue/total`
  - **Query Parameters:**
    - `start_date` (e.g., `2023-01-01`)
    - `end_date` (e.g., `2024-01-01`)
  - **Description:** Returns the total revenue for the specified date range.
  - **Response:**
    ```json
    {
      "total_revenue": 12345.67
    }
    ```

## Project Structure

- **main.go:** Main application file, including server setup and route definitions.
- **models.go:** Defines database schema and GORM models.
- **data_loader.go:** Contains functions for loading CSV data into the database.
- **cron.go:** Configures and starts the cron job for periodic data refresh.

