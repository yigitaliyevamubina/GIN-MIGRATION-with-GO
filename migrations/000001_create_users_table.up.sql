CREATE TABLE users (
    id_num SERIAL PRIMARY KEY, 
    uuid UUID,
    first_name VARCHAR(35), 
    last_name VARCHAR(35),
    age INT, 
    role_id INT
    )