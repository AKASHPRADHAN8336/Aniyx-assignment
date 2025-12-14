-- Initialize database schema
CREATE DATABASE IF NOT EXISTS userdb;
USE userdb;

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data (optional)
INSERT IGNORE INTO users (name, dob) VALUES
('Akash Kumar Pradhan', '1990-05-15'),
('Rohit Yadva', '1985-11-22'),
('ramesh kumar', '1995-03-08'),
('Eklavya singh rathore', '1992-07-30');

-- Create indexes for better performance
CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_users_dob ON users(dob);