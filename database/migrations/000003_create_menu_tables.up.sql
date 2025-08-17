USE ordering_db;

CREATE TABLE IF NOT EXISTS Categories (
    category_id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    category_name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Items (
    item_id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    item_name VARCHAR(255),
    category_id BIGINT,
    price INT UNSIGNED,
    description TEXT,
    item_image_url VARCHAR(255),
    is_veg BOOLEAN,
    spice_level TINYINT,
    FOREIGN KEY (category_id) REFERENCES Categories(category_id)
);