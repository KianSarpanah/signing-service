# Signature Service - Coding Challenge

## Overview

The **Signature Service** is a simple Go-based service designed to manage devices by creating, signing, and verifying them. This service includes API endpoints for interacting with the devices and checking the application’s health.

## Quick Start

### Prerequisites:
- Language **Go**

### API Endpoints

You can interact with the Signature Service using tools like Postman or cURL. Below are the available endpoints:


1. Create a New Device
    Creates a new device entry in the system.
    Endpoint: POST /api/v0/devices/create
    Example:
    ```bash
    curl -X POST http://localhost:8080/api/v0/devices/create \
    -H "Content-Type: application/json" \
    -d '{
        "algorithm": "RSA"
    }'

2. Sign a Device
    Sign into a specific device using its ID.
    Endpoint: POST /api/v0/devices/{device_id}/sign
    Example:
    ```bash
    curl -X POST http://localhost:8080/api/v0/devices/{id}/sign \
    -H "Content-Type: application/json" \
    -d '{
        "data": "transactiontest"
    }'

3. Verify a signed device 
    Verify if a device has been signed by checking its status.
    Endpoint: GET /api/v0/devices/{id}
    `curl -X GET http://localhost:8080/api/v0/devices/{id}`

4. Application Health Check:
    check if the service is running and operational.
    Endpoint: GET /api/v0/health
    Example:
    `curl -X GET http://localhost:8080/api/v0/health`

5. List all signed devices:
    Retrieve a list of all devices that have been signed.
    Endpoint: GET /devices
    Example:
    `curl -X GET http://localhost:8080/devices`

#### QA / Testing

- `go test ./...`



```bash
go run main.go