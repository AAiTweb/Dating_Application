CREATE TABLE users(
    user_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(20),
    password VARCHAR(20),
    email VARCHAR(50),
    confirmation_token VARCHAR(80),
    varified INT DEFAULT 0,
    created_date DATE

);

CREATE TABLE relationship(
    relationship_id BIGSERIAL PRIMARY KEY,
    user_one_id INT,
    user_two_id INT,
    relationship_status INT
);

CREATE TABLE messages(
    messages_id BIGSERIAL PRIMARY KEY,
    from_id INT,
    to_id INT,
    messages TEXT,
    send_time TIMESTAMP
);

CREATE TABLE user_profile(
    user_profile_id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(20),
    second_name VARCHAR(20),
    profile_picture TEXT,
    sex VARCHAR(7),
    online_offline_status INT

);

CREATE TABLE user_location(
    user_location_id BIGSERIAL PRIMARY KEY,
    location_country VARCHAR(30),
    location_city VARCHAR(30)    
);


