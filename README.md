# Transaction Timestamp Service

This Go application provides a simple web service to fetch the last active timestamp of a transaction given its transaction ID. It uses the Gorilla Mux router for routing and makes HTTP requests to an external API to retrieve the transaction data.

## Features

- Home page with basic instructions.
- Endpoint to get the last active timestamp of a transaction by its ID.

## Endpoints

- `GET /`: Home page with instructions.
- `GET /api/tx/{txid}`: Fetch the last active timestamp for the transaction with the specified ID.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)

### Installation

1. Clone the repository:

    ```sh
    https://github.com/Abhinav7903/TestAPI_Txid.git
    cd TestAPI_Txid
    ```

2. Install the required packages:

    ```sh
    go get -u github.com/gorilla/mux
    ```

### Running the Application

1. Build and run the application:

    ```sh
    go run main.go
    ```

2. The server will start listening on port `8080`. You should see the following output:

    ```sh
    Server listening on port 8080
    ```

### Usage

- To access the home page, open your browser and navigate to `http://localhost:8080/`.

    You should see a message saying:

    ```
    Hello! This is the home page.
    To get the last active timestamp for a transaction ID, make a GET request to /api/tx/{txid}
    Example: /api/tx/1a2b3c4d5e6f7g8h9i0j
    ```

- To get the last active timestamp for a specific transaction ID, make a GET request to `http://localhost:8080/api/tx/{txid}`. Replace `{txid}` with the actual transaction ID.

    Example request:

    ```sh
    curl http://localhost:8080/api/tx/1a2b3c4d5e6f7g8h9i0j
    ```

    Example response:

    ```json
    {
      "message": "success",
      "last_active_timestamp": "2023-05-23 15:04:05"
    }
    ```

## Code Overview

### Main Components

- **Main Function**: Sets up the router and starts the server.
- **Home Handler**: Provides basic instructions on the home page.
- **Transaction Handler**: Fetches and returns the last active timestamp for a given transaction ID.

### Error Handling

The application handles various error scenarios, such as failing to fetch data from the external API or failing to decode the response. Appropriate error messages are returned in the response.

## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux): A powerful URL router and dispatcher for golang.


## Acknowledgements

- [Gorilla Mux](https://github.com/gorilla/mux) for the routing library.

