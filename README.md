# Go Work Benefits API

This API, developed in Go, calculates common labor benefits in Brazil, such as net salary, vacation pay, and 13th salary. It uses the Gin framework for routing and request handling.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Contributing](#contributing)

## Prerequisites

- [Go](https://golang.org/) 1.23.2+
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) for easy environment setup

## Installation

Clone the repository:

```bash
git clone https://github.com/amorimluiz/go-work-benefits-api.git
cd go-work-benefits-api
```

## Configuration

Set up environment variables for MySQL database connection:

```bash
APP_PORT=
INSS_BRACKETS=
INSS_RATES=
IRRF_BRACKETS=
IRRF_RATES=
IRRF_DEDUCTIONS=
DEDUCTION_PER_DEPENDENT=
```

Create a `.env` file in the project root and define these variables.

## Running the Application

To start the environment using Docker Compose:

```bash
docker-compose up
```

To run the API locally:

```bash
go build -o ./tmp/main.exe ./cmd/workbenefitsapi/main.go
./tmp/main.exe
```

## Contributing

To contribute:

1. Fork the project
2. Create a branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
