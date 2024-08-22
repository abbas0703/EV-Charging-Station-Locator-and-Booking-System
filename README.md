# EV-Charging-Station-Locator-and-Booking-System

## Overview

This project is a web application that allows users to find nearby electric vehicle (EV) charging stations and book time slots for charging. It uses the Gin framework for the backend, Google Maps API for locating charging stations, and serves static files and HTML templates.

## Features

- **Find Charging Stations**: Users can search for nearby EV charging stations based on their location.
- **Book Charging Slots**: Users can book time slots for charging at selected stations.

## Prerequisites

Before you begin, ensure you have the following:

1. **Go**: Version 1.18 or higher installed on your system.
2. **Gin**: Gin-Gonic web framework.
3. **Google Maps API Key**: You need a Google Maps API key for accessing the Maps API.

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/ev-charging-booking-system.git
   cd ev-charging-booking-system
   ```
2. **Install Dependencies**

Ensure you have Go installed, then install the required Go packages:
  ```bash
  go mod tidy
  ```

3. **Create a .env File**

Create a .env file in the root directory of the project with the following content:

  ```bash
  GOOGLE_MAPS_API_KEY=your_google_maps_api_key
  ```

Replace your_google_maps_api_key with your actual Google Maps API key.

4. **Directory Structure**
 
  ```go
  ├── static/
├── templates/
├── main.go
├── go.mod
├── go.sum
└── .env
  ```

-static/ – Directory for static files (e.g., CSS, JavaScript).
-templates/ – Directory for HTML templates.


##Running the Application

1. **Start the Server**

 ```bash
 go run main.go
 ```

2. **Access the Application**

Open your web browser and navigate to http://localhost:8080 to use the application.


## Endpoints

### GET /charging-stations

- **Query Parameters**: 
  - `location` (optional, format: `"latitude,longitude"`)

- **Response**: 
  - JSON object containing nearby charging stations.

### POST /book-slot

- **Request Body**: 
  - JSON object with booking details:
    ```json
    {
      "user_id": "string",
      "charging_station_id": "string",
      "start_time": "string",
      "end_time": "string"
    }
    ```

- **Response**: 
  - JSON object indicating booking status.

## Error Handling

- **API Key Not Set**: 
  - Check if the `.env` file contains the correct Google Maps API key.

- **Invalid Location Format**: 
  - Ensure the location parameter is in `"latitude,longitude"` format.

- **Invalid Booking Data**: 
  - Ensure the JSON body contains all required fields.


## Contributing

Feel free to fork the repository, submit pull requests, and contribute to the project. For any issues or feature requests, please open an issue on the GitHub repository.


MIT License

Copyright (c) [2024] [Jidaar Abbas]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.








