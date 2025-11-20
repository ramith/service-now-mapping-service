# ServiceNow Mapping Service API Documentation

## Overview

RESTful API for managing mappings between ServiceNow system IDs and API configurations.

**Base URL**: `http://localhost:8080`

## API Endpoints

### Health Check

#### GET /health
Check if the service is running.

**Response** (200 OK):
```json
{
  "status": "healthy",
  "service": "servicenow-mapping-service"
}
```

---

### Mappings

#### POST /api/v1/mappings
Create a new API mapping.

**Request Body**:
```json
{
  "sys_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "api_id": "api-001",
  "api_name": "User Management API",
  "api_version": "v1.0.0"
}
```

**Response** (201 Created):
```json
{
  "id": 1,
  "sys_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "api_id": "api-001",
  "api_name": "User Management API",
  "api_version": "v1.0.0",
  "created_at": "2025-11-21T10:00:00Z",
  "updated_at": "2025-11-21T10:00:00Z"
}
```

---

#### GET /api/v1/mappings
Get all API mappings.

**Response** (200 OK):
```json
[
  {
    "id": 1,
    "sys_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
    "api_id": "api-001",
    "api_name": "User Management API",
    "api_version": "v1.0.0",
    "created_at": "2025-11-21T10:00:00Z",
    "updated_at": "2025-11-21T10:00:00Z"
  }
]
```

---

#### GET /api/v1/mappings/sys-id/:sys_id
Get mapping by ServiceNow system ID.

**Parameters**:
- `sys_id` (path) - ServiceNow System ID

**Example**: `GET /api/v1/mappings/sys-id/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6`

**Response** (200 OK):
```json
{
  "id": 1,
  "sys_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "api_id": "api-001",
  "api_name": "User Management API",
  "api_version": "v1.0.0",
  "created_at": "2025-11-21T10:00:00Z",
  "updated_at": "2025-11-21T10:00:00Z"
}
```

---

#### GET /api/v1/mappings/api-id/:api_id
Get mapping by API ID.

**Parameters**:
- `api_id` (path) - API Identifier

**Example**: `GET /api/v1/mappings/api-id/api-001`

**Response** (200 OK):
```json
{
  "id": 1,
  "sys_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "api_id": "api-001",
  "api_name": "User Management API",
  "api_version": "v1.0.0",
  "created_at": "2025-11-21T10:00:00Z",
  "updated_at": "2025-11-21T10:00:00Z"
}
```

---

#### GET /api/v1/mappings/search
Get mapping by API name and version.

**Query Parameters**:
- `api_name` (required) - API Name
- `api_version` (required) - API Version

**Example**: `GET /api/v1/mappings/search?api_name=User%20Management%20API&api_version=v1.0.0`

**Response** (200 OK):
```json
{
  "id": 1,
  "sys_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "api_id": "api-001",
  "api_name": "User Management API",
  "api_version": "v1.0.0",
  "created_at": "2025-11-21T10:00:00Z",
  "updated_at": "2025-11-21T10:00:00Z"
}
```

---

#### PUT /api/v1/mappings/:id
Update an existing API mapping.

**Parameters**:
- `id` (path) - Mapping ID

**Request Body**:
```json
{
  "api_id": "api-001-updated",
  "api_name": "Updated User Management API",
  "api_version": "v2.0.0"
}
```

**Response** (200 OK):
```json
{
  "id": 1,
  "sys_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "api_id": "api-001-updated",
  "api_name": "Updated User Management API",
  "api_version": "v2.0.0",
  "created_at": "2025-11-21T10:00:00Z",
  "updated_at": "2025-11-21T10:15:00Z"
}
```

---

#### DELETE /api/v1/mappings/:id
Delete an API mapping (soft delete).

**Parameters**:
- `id` (path) - Mapping ID

**Example**: `DELETE /api/v1/mappings/1`

**Response** (204 No Content)

---

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message description"
}
```

### Common Status Codes

- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `204 No Content` - Resource deleted successfully
- `400 Bad Request` - Invalid request parameters
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

---

## Example Usage with curl

### Create a mapping
```bash
curl -X POST http://localhost:8080/api/v1/mappings \
  -H "Content-Type: application/json" \
  -d '{
    "sys_id": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
    "api_id": "api-001",
    "api_name": "User Management API",
    "api_version": "v1.0.0"
  }'
```

### Search by sys_id
```bash
curl http://localhost:8080/api/v1/mappings/sys-id/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6
```

### Search by api_id
```bash
curl http://localhost:8080/api/v1/mappings/api-id/api-001
```

### Search by api_name and api_version
```bash
curl "http://localhost:8080/api/v1/mappings/search?api_name=User%20Management%20API&api_version=v1.0.0"
```

### Get all mappings
```bash
curl http://localhost:8080/api/v1/mappings
```

### Update a mapping
```bash
curl -X PUT http://localhost:8080/api/v1/mappings/1 \
  -H "Content-Type: application/json" \
  -d '{
    "api_name": "Updated User Management API",
    "api_version": "v2.0.0"
  }'
```

### Delete a mapping
```bash
curl -X DELETE http://localhost:8080/api/v1/mappings/1
```
