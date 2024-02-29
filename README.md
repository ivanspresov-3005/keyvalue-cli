# KeyValue CLI

KeyValue CLI is a user-friendly command-line key-value storage system implemented in Golang. It is designed to provide a simple yet powerful solution for managing key-value pairs with ease.

## Features

- **Set Method:** Store key-value pairs with user-specified or default Time-to-Live (TTL).
- **Get Method:** Retrieve values associated with the given keys.
- **Delete Method:** Remove key-value pairs from storage.
- **TTL Implementation:** Automatic deletion of expired data based on user-provided or default TTL.
- **Usability:** Intuitive and user-friendly command-line interface.
- **Efficient Storage:** In-memory data structure ensures efficient key-value pair management.
- **Documentation:** Inline comments for code readability, along with a comprehensive user guide.

## Project Structure

- `/cli`: Entry point for the command-line application.
- `/kvlogic`: Implementation of key-value storage functionality and tests.

## Getting Started

### Clone the repository

```bash
git clone https://github.com/ivanspresov-3005/keyvalue-cli.git```


## Navigate to the project directory

- ```cd keyvalue-cli```

## Build the application

- ```go build ./cli```


