CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,  
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT,
    phone_number TEXT,
    token TEXT,
    role TEXT,
    expires TIMESTAMP WITH TIME ZONE
);

INSERT INTO users (username, email, password, role)
VALUES ('administrator', 'administrator@example.com', 'administrator', 'administrator');