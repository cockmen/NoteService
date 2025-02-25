CREATE TABLE users (
                    id SERIAL PRIMARY KEY,
                    email VARCHAR(255),
                    password VARCHAR(255),
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (email, password, created_at)VALUES(
                    ('kevin.carter1997@gmail.com','asdfg','2025-01-30 15:02:30');
)