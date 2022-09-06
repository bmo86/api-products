DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id VARCHAR(32) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    pass VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),   
    name VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    date_brithday TIMESTAMP NOT NULL DEFAULT NOW()   
);


