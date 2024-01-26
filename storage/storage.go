package storage

import (
	"GIN_MIGRATION/models"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	connection := "user=postgres password=mubina2007 dbname=migration sslmode=disable"
	mydb, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	return mydb, err
}

func CreateUser(reqUser models.User) (*models.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	uuid := uuid.NewString()
	reqUser.UUID = uuid

	query := `INSERT INTO users(uuid, first_name, last_name, age, role_id) VALUES($1, $2, $3, $4, $5) RETURNING id_num, uuid, first_name, last_name, age, role_id`
	rowUser := db.QueryRow(query, reqUser.UUID, reqUser.FirstName, reqUser.LastName, reqUser.Age, reqUser.RoleId)

	query = `INSERT INTO roles_users(user_id, role_id) VALUES($1, $2)`

	_ = db.QueryRow(query, reqUser.ID, reqUser.RoleId)

	var respUser models.User
	if err := rowUser.Scan(&respUser.ID, &respUser.UUID, &respUser.FirstName, &respUser.FirstName, &respUser.Age, &respUser.RoleId); err != nil {
		return nil, err
	}

	return &respUser, nil
}

func UpdateUser(id int, user models.User) (*models.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	query := `UPDATE users SET first_name = $1, last_name = $2, age = $3 WHERE id_num = $4 RETURNING id_num, uuid, first_name, last_name, age, role_id`
	rowUser := db.QueryRow(query, user.ID, user.FirstName, user.LastName, user.Age, id)

	var updatedUser models.User
	if err := rowUser.Scan(&updatedUser.ID, &updatedUser.UUID, &updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Age, &updatedUser.RoleId); err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func DeleteUser(id int) (*models.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	query := `DELETE FROM roles_users WHERE user_id = $1`
	_, err = db.Exec(query, id)

	query = `DELETE FROM users WHERE id_num = $1 RETURNING id_num, uuid, first_name, last_name, age, role_id`
	rowUser := db.QueryRow(query, id)

	var deletedUser models.User
	if err := rowUser.Scan(&deletedUser.ID, &deletedUser.UUID, &deletedUser.FirstName, &deletedUser.LastName, &deletedUser.Age, &deletedUser.RoleId); err != nil {
		return nil, err
	}

	return &deletedUser, err
}

func GetUserById(id int) (*models.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	query := `SELECT id_num, uuid, first_name, last_name, age, role_id FROM users WHERE id_num = $1`
	rowUser := db.QueryRow(query, id)

	var respUser models.User
	if err := rowUser.Scan(&respUser.ID, &respUser.UUID, &respUser.FirstName, &respUser.LastName, &respUser.Age, &respUser.RoleId); err != nil {
		return nil, err
	}

	return &respUser, nil
}

func GetAllUsers(limit int, page int) ([]*models.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	offset := limit * (page - 1)
	query := `SELECT id_num, uuid, first_name, last_name, age, role_id FROM users LIMIT $1 OFFSET $2`
	rows, err := db.Query(query, limit, offset)

	var users []*models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UUID, &user.FirstName, &user.LastName, &user.Age, &user.RoleId); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func FilterByName(name string, limit int, page int) ([]*models.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	offset := limit * (page - 1)
	query := `SELECT id_num, 
				uuid, 
				first_name, 
				last_name, 
				age, 
				role_id 
				FROM users 
				WHERE 
				(first_name || ' ' || last_name ILIKE $1)
				OR 
				(last_name || ' ' || first_name ILIKE $1)
				LIMIT $2 OFFSET $3`

	var respUsers []*models.User
	rows, err := db.Query(query, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var respUser models.User
		if err := rows.Scan(&respUser.ID, &respUser.UUID, &respUser.FirstName, &respUser.LastName, &respUser.Age, &respUser.RoleId); err != nil {
			return nil, err
		}
		respUsers = append(respUsers, &respUser)
	}

	return respUsers, nil
}

func GetUsersByRole(roleId int, limit int, page int) ([]*models.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	offset := limit * (page - 1)

	query := `
		SELECT u.id_num, 
		u.uuid,
		u.first_name, 
		u.last_name, 
		u.age,
		r.id,
		r.name
		FROM users u
		JOIN roles_users ru 
		ON u.id_num = ru.user_id
		JOIN roles r 
		ON r.id = ru.role_id
		WHERE ru.role_id = $1
		LIMIT $2 OFFSET $3
	`

	rows, err := db.Query(query, roleId, limit, offset)
	if err != nil {
		return nil, err
	}

	var respUsers []*models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.UUID, &user.FirstName, &user.LastName, &user.Age, &user.RoleId, &user.RoleName)
		if err != nil {
			return nil, err
		}

		respUsers = append(respUsers, &user)
	}

	return respUsers, nil
}
