package todos

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"
	"untitled/util"
)

type Detail struct {
	Color string `json:"color"`
}

func (d *Detail) Scan(src interface{}) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &d)
}

type Todo struct {
	Id        int    `json:"id"`
	Body      string `json:"body"`
	Status    string `json:"status"`
	Detail    Detail `json:"detail"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func Find(id int) (*Todo, error) {
	db := util.GetDB()
	row, err := db.Query(
		"SELECT id, body, status, detail, created_at, updated_at, deleted_at FROM todos WHERE id=?",
		id,
	)
	if err != nil {
		return nil, err
	}

	todo := Todo{}
	var deletedAt sql.NullString
	for row.Next() {
		err = row.Scan(&todo.Id, &todo.Body, &todo.Status, &todo.Detail, &todo.CreatedAt, &todo.UpdatedAt, deletedAt)
		if err != nil {
			log.Fatal("error: ", err)
		}
		if deletedAt.Valid {
			todo.DeletedAt = deletedAt.String
		}
	}

	row.Close()
	return &todo, nil
}

func FindAll() ([]*Todo, error) {
	db := util.GetDB()
	rows, err := db.Query(
		"SELECT id, body, status, detail, created_at, updated_at FROM todos WHERE deleted_at IS NULL",
	)
	if err != nil {
		return nil, err
	}

	todos := []*Todo{}
	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.Id, &todo.Body, &todo.Status, &todo.Detail, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	rows.Close()
	return todos, nil
}

func Update(id int, body string, status string) error {
	db := util.GetDB()
	_, err := db.Exec(
		"UPDATE todos SET body = ?, status = ?, detail = '{}' WHERE id = ?",
		body,
		status,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func Create(body string, status string) error {
	db := util.GetDB()
	_, err := db.Exec(
		"INSERT INTO todos(body, status, detail) VALUES (?,?,'{}')",
		body,
		status,
	)
	if err != nil {
		return err
	}

	return nil
}

func Upsert(id int, body string, status string) error {
	exists, err := Find(id)
	if err != nil {
		return err
	}
	if exists.Id == 0 {
		Create(body, status)
	} else {
		Update(id, body, status)
	}
	return nil
}

func Delete(id int) error {
	db := util.GetDB()
	now := time.Now()
	_, err := db.Exec(
		"UPDATE todos SET deleted_at = ? WHERE id = ?",
		now,
		id,
	)

	if err != nil {
		return err
	}

	return nil

}
