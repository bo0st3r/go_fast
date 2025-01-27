# Go Fast - Telemetry Service

A simple telemetry service built with Go that allows you to store and retrieve metric data.

## Prerequisites

- Go 1.23.5 or higher
- Docker and Docker Compose
- Make (optional)
- Postman (optional)

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/bo0st3r/go-fast.git
cd go-fast
```

2. Create a `.env` file in the root directory with the following content:
```bash
POSTGRES_USER=your_username
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_database_name
APP_PORT=8080
```

3. Start the database:
```bash
make db-up
```

4. Run the application:
```bash
make dev
```

The service will be available at `http://localhost:8080`

## API Endpoints

### Health Check
```bash
curl http://localhost:8080/health
```

### Create Telemetry Entry
```bash
curl -X POST http://localhost:8080/v1/telemetry \
  -H "Content-Type: application/json" \
  -d '{"metric": "nb_players", "value": 5}'
```

### Get All Telemetry Entries
```bash
curl http://localhost:8080/v1/telemetry
```

### Get Peak Values Per Metric
```bash
curl http://localhost:8080/v1/telemetry/peak
```

## Cleanup

To stop the database:
```bash
make db-down
```

To clean built artifacts:
```bash
make clean
```

## Postman Collection

A Postman collection is included in the repository for easy API testing. Import `go-fast.postman_collection.json` into Postman to get started. 