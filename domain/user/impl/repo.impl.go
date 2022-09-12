package impl

import (
	"database/sql"
	"dealljobs/domain/user"
	"strings"
	"time"
)

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) user.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Persist(u *user.User) (*user.User, error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	statement, _ := r.db.Prepare("INSERT INTO users (username, name, password, address, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	res, err := statement.Exec(u.Username, u.Name, u.Password, u.Address, u.Role, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	_ = statement.Close()

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.Id = int(lastId)

	return u, nil

}

func (r *repo) Update(u *user.User) error {
	u.UpdatedAt = time.Now()

	statement, _ := r.db.Prepare("UPDATE users set username = ?, name = ?, password = ?, address = ?, role = ?, updated_at = ? where id = ?")
	_, err := statement.Exec(u.Username, u.Name, u.Password, u.Address, u.Role, u.UpdatedAt, u.Id)
	if err != nil {
		return err
	}

	err = statement.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(id int) error {
	statement, _ := r.db.Prepare("UPDATE users set updated_at = now() and deleted_at = now() where id = ?")
	_, err := statement.Exec(id)
	if err != nil {
		return err
	}

	err = statement.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetUser(id int) *user.User {
	row := r.db.QueryRow("SELECT id, username, name, password, address, role, created_at, updated_at, deleted_at  FROM users where id= ? and deleted_at is null", id)
	u := user.User{}

	err := row.Scan(&u.Id, &u.Username, &u.Name, &u.Password, &u.Address, &u.Role, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		println("Exec err:", err.Error())
	}

	if u.Id == 0 {
		return nil
	}

	return &u
}

func (r *repo) GetUserByUserPass(username, password string) *user.User {
	row := r.db.QueryRow("SELECT id, username, name, password, address, role, created_at, updated_at, deleted_at  FROM users where username= ? and password = ? and deleted_at is null", username, password)
	u := user.User{}

	err := row.Scan(&u.Id, &u.Username, &u.Name, &u.Password, &u.Address, &u.Role, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		println("Exec err:", err.Error())
	}

	if u.Id == 0 {
		return nil
	}

	return &u
}

func (r *repo) GetUserByUsername(i user.User) (*user.User, error) {
	row := r.db.QueryRow("SELECT id, username, name, password, address, role, created_at, updated_at, deleted_at  FROM users where username = ? and deleted_at is null", i.Username)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	u := user.User{}
	err = row.Scan(&u.Id, &u.Username, &u.Name, &u.Password, &u.Address, &u.Role, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return nil, err
	}

	if u.Id == 0 {
		return nil, nil
	}

	return &u, nil
}

func convertStringListToInterface(list []string) []interface{} {
	args := make([]interface{}, 0)
	for _, el := range list {
		args = append(args, el)
	}

	return args
}

func (r *repo) GetUsers(idsRequest []string) []*user.User {
	query := "SELECT id, username, name, password, address, role, created_at, updated_at, deleted_at  FROM users where deleted_at is null and id in (?" + strings.Repeat(",?", len(idsRequest)-1) + ")"
	args := convertStringListToInterface(idsRequest)
	rows, err := r.db.Query(query, args...)
	defer rows.Close()

	if err != nil {
		println("Exec err:", err.Error())
	}

	var list []*user.User
	for rows.Next() {
		u := user.User{}

		err = rows.Scan(&u.Id, &u.Username, &u.Name, &u.Password, &u.Address, &u.Role, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			println("Exec err:", err.Error())
		}

		list = append(list, &u)
	}
	return list
}

func (r *repo) GetAllUsers() []*user.User {
	query := "SELECT id, username, name, password, address, role, created_at, updated_at, deleted_at FROM users where deleted_at is null"
	rows, err := r.db.Query(query)
	defer rows.Close()

	if err != nil {
		println("Exec err:", err.Error())
	}

	var list []*user.User
	for rows.Next() {
		u := user.User{}

		err = rows.Scan(&u.Id, &u.Username, &u.Name, &u.Password, &u.Address, &u.Role, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			println("Exec err:", err.Error())
		}

		list = append(list, &u)
	}
	return list
}
