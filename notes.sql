CREATE TABLE notes(
                    id SERIAL PRIMARY KEY,
                    title VARCHAR(50),
                    body VARCHAR(255),
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO notes(title, body, created_at, updated_at)VALUES
                                            ('Сходить в магазин','Нужно не забыть сходить в магазин и купить молока', '2025-01-30 15:02:30', '2025-01-30 15:02:30');
