USE ordering_db;

CREATE TABLE IF NOT EXISTS Orders (
    order_id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_id BIGINT,
    instructions TEXT,
    order_type ENUM('dine_in', 'takeaway'),
    table_number INT DEFAULT 0,
    status ENUM('preparing', 'payment_pending', 'completed'),
    total_cost INT UNSIGNED,
    FOREIGN KEY (user_id) REFERENCES User(user_id)
);
CREATE TABLE IF NOT EXISTS Ordered_Items (
    order_id BIGINT,
    item_id BIGINT,
    quantity INT,
    dish_complete BOOLEAN DEFAULT FALSE,
    chef_id BIGINT DEFAULT 1,
    PRIMARY KEY (order_id, item_id),
    FOREIGN KEY (order_id) REFERENCES Orders (order_id),
    FOREIGN KEY (item_id) REFERENCES Items(item_id),
    FOREIGN KEY (chef_id) REFERENCES User(user_id)
);
CREATE TABLE IF NOT EXISTS Payment (
    transaction_id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    order_id BIGINT,
    tip_amount INT UNSIGNED DEFAULT 0,
    discount_reward_points BIGINT,
    amount_paid INT UNSIGNED,
    payment_status ENUM('pending', 'paid', 'failed'),
    FOREIGN KEY (order_id) REFERENCES Orders (order_id)
);
CREATE TABLE IF NOT EXISTS Reviews (
    review_id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    order_id BIGINT,
    comments LONGTEXT,
    ambience_stars TINYINT DEFAULT 0,
    food_quality_stars TINYINT DEFAULT 0,
    service_stars TINYINT DEFAULT 0,
    value_for_money_stars TINYINT DEFAULT 0,
    star_rating FLOAT,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id)
);