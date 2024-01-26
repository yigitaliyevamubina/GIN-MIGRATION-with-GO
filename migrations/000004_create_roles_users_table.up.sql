CREATE TABLE IF NOT EXISTS roles_users(user_id INT, role_id INT,
    FOREIGN KEY(user_id) REFERENCES users(id_num),
    FOREIGN KEY(role_id) REFERENCES roles(id)
)
