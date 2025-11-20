-- Migration: 001_create_api_mappings_table
-- Description: Creates the api_mappings table for storing ServiceNow to API mappings
-- Author: sql-pro agent
-- Date: 2025-11-21

-- Create database if not exists
CREATE DATABASE IF NOT EXISTS servicenow_mapping
    DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_unicode_ci;

USE servicenow_mapping;

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
    
    -- Unique constraint: one sys_id can only map to one API configuration
    UNIQUE KEY uk_sys_id (sys_id),
    
    -- Unique constraint: prevent duplicate API configurations
    UNIQUE KEY uk_api_name_version (api_name, api_version, deleted_at),
    
    -- Index for searching by api_id
    INDEX idx_api_id (api_id, deleted_at),
    
    -- Index for searching by api_name and version combination
    INDEX idx_api_name_version (api_name, api_version, deleted_at),
    
    -- Index for soft delete queries
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB 
  DEFAULT CHARSET=utf8mb4 
  COLLATE=utf8mb4_unicode_ci
  COMMENT='Stores ServiceNow system ID to API mappings';

-- Create audit log table for tracking changes
CREATE TABLE IF NOT EXISTS api_mapping_audit_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    mapping_id BIGINT UNSIGNED NOT NULL,
    action VARCHAR(20) NOT NULL COMMENT 'CREATE, UPDATE, DELETE',
    old_value JSON NULL COMMENT 'Previous state',
    new_value JSON NULL COMMENT 'New state',
    changed_by VARCHAR(100) DEFAULT 'system',
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    INDEX idx_mapping_id (mapping_id),
    INDEX idx_action (action),
    INDEX idx_changed_at (changed_at),
    
    FOREIGN KEY (mapping_id) 
        REFERENCES api_mappings(id) 
        ON DELETE CASCADE
) ENGINE=InnoDB 
  DEFAULT CHARSET=utf8mb4 
  COLLATE=utf8mb4_unicode_ci
  COMMENT='Audit trail for API mapping changes';

-- Insert sample data for testing
INSERT INTO api_mappings (sys_id, api_id, api_name, api_version) VALUES
    ('a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6', 'api-001', 'User Management API', 'v1.0.0'),
    ('b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7', 'api-002', 'Order Processing API', 'v2.1.0'),
    ('c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8', 'api-003', 'Inventory Service API', 'v1.5.2'),
    ('d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9', 'api-004', 'Payment Gateway API', 'v3.0.0'),
    ('e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0', 'api-005', 'Notification Service API', 'v1.2.0')
ON DUPLICATE KEY UPDATE updated_at = CURRENT_TIMESTAMP;
