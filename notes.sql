CREATE TABLE notes(
                    id SERIAL PRIMARY KEY,
                    title VARCHAR(50),
                    body VARCHAR(255),
                    user_id INTEGER,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);