CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(20)
);

INSERT INTO roles(name) VALUES ('teacher'), ('student'), ('admin');