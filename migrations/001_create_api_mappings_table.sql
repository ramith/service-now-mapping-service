-- Migration: 001_create_api_mappings_table
-- Description: Creates the api_mappings table for storing ServiceNow to API mappings

-- Create api_mappings table
CREATE TABLE IF NOT EXISTS api_mappings (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sys_id VARCHAR(32) NOT NULL COMMENT 'ServiceNow System ID',
    api_id VARCHAR(100) NOT NULL COMMENT 'API Identifier',
    api_name VARCHAR(255) NOT NULL COMMENT 'API Name',
    api_version VARCHAR(50) NOT NULL COMMENT 'API Version',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    -- Unique constraint on sys_id
    UNIQUE KEY uk_sys_id (sys_id),
    
    -- Unique constraint on api_name and api_version combination
    UNIQUE KEY uk_api_name_version (api_name, api_version),
    
    -- Index for searching by api_id
    INDEX idx_api_id (api_id),
    
    -- Index for soft delete queries
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB 
  DEFAULT CHARSET=utf8mb4 
  COLLATE=utf8mb4_unicode_ci
  COMMENT='Stores ServiceNow system ID to API mappings';
