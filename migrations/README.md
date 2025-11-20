# Database Migrations

This directory contains MySQL migration scripts for the ServiceNow Mapping Service.

## Running Migrations

### Using MySQL CLI

```bash
# Connect to MySQL
mysql -u root -p

# Run the migration
source migrations/001_create_api_mappings_table.sql
```

### Using Docker

```bash
# Start MySQL with Docker Compose
docker-compose up -d mysql

# Run migration
docker exec -i servicenow-mysql mysql -uroot -ppassword < migrations/001_create_api_mappings_table.sql
```

### Using MySQL Workbench or DBeaver

1. Open the SQL file in your tool
2. Execute the script

## Migration Files

- `001_create_api_mappings_table.sql` - Creates initial schema with api_mappings table and audit logs

## Schema Overview

### api_mappings Table

Stores the mapping between ServiceNow system IDs and API configurations.

| Column | Type | Description |
|--------|------|-------------|
| id | BIGINT | Primary key (auto-increment) |
| sys_id | VARCHAR(32) | ServiceNow System ID (unique) |
| api_id | VARCHAR(100) | API Identifier |
| api_name | VARCHAR(255) | API Name |
| api_version | VARCHAR(50) | API Version |
| created_at | TIMESTAMP | Record creation time |
| updated_at | TIMESTAMP | Last update time |
| deleted_at | TIMESTAMP | Soft delete timestamp |

### Indexes

- `uk_sys_id`: Unique index on sys_id for fast lookups
- `uk_api_name_version`: Unique constraint on (api_name, api_version)
- `idx_api_id`: Index for searching by api_id
- `idx_api_name_version`: Composite index for (api_name, api_version) searches
- `idx_deleted_at`: Index for soft delete filtering

## Search Capabilities

The schema supports efficient searching by:
1. **sys_id**: Direct unique lookup (O(1))
2. **api_id**: Indexed search
3. **(api_name, api_version)**: Composite index search
