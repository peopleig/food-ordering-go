USE ordering_db;

CREATE TABLE IF NOT EXISTS User (
        user_id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
        first_name VARCHAR(255),
        last_name VARCHAR(255),
        email_id VARCHAR(255),
        mobile_number VARCHAR(10),
        password VARCHAR(255),
        role ENUM('admin', 'chef', 'customer'),
        reward_points INT DEFAULT 0,
        last_visited DATE DEFAULT '2000-01-01',
        approved BOOLEAN DEFAULT TRUE
    );