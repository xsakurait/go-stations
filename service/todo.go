package service

import (
	"context"
	"database/sql"
	"github.com/TechBowl-japan/go-stations/model"
	"log"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = "INSERT INTO todos(subject, description) VALUES(?, ?)"
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	// ctx=関数の第１引数はrequestオブジェクトみたいなものでcontext.Context型と呼ばれる　エラー等を制御する役割を担う
	stmt, err := s.db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		log.Fatal("server/todo.go s.db.ExecContext(ctx,insert,subject,description)", err)
		return  nil,err
	}
	insert_id, err := stmt.LastInsertId()
	if err != nil {
		log.Fatal("server/todo.go stmt.LastInsertId()", err)
		return nil,err

	}
	var todo model.TODO
	err = s.db.QueryRowContext(ctx, confirm, insert_id).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Fatal("server/todo.go s.db.QueryRowContext(ctx,confirm,insert_id).Scan(&todo.Subject,&todo.Description,&todo.CreatedAt,&todo.UpdatedAt)", err)
		return nil,err

	}

	return &todo, err
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
