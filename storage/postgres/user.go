package postgres

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"main.go/models"
	"main.go/storage"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) storage.IUserStorage {
	return userRepo{
		db: db,
	}
}

func (u userRepo) Create(createuser models.CreateUser) (models.User, error) {
	id := uuid.New()
	if _, err := u.db.Exec(`insert into users values($1,$2,$3,$4,$5,$6)`, id, createuser.FullName, createuser.Phone, createuser.Password, createuser.Cash, createuser.UserType); err != nil {
		fmt.Println("error while insertinf user ", err.Error())
		return models.User{}, err
	}
	return models.User{}, nil

}

func (u userRepo) GetByID(id models.PrimaryKey) (models.User, error) {
	users := models.User{}
	if err := u.db.QueryRow(`select id, fullname, phone,password_, cash, user_role where id=$1`, id.ID).Scan(
		&users.ID,
		&users.FullName,
		&users.Phone,
		&users.Password,
		&users.Cash,
		&users.UserType,
	); err != nil {
		return models.User{}, err
	}
	return models.User{}, nil

}

func (u userRepo) GetListRequest(models.GetAllRequest) (models.UsersResponse, error) {
	rows, err := u.db.Query(`select * from users`)
	if err != nil {
		return models.UsersResponse{}, err
	}
	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.ID, &user.FullName, &user.Phone, &user.Password, &user.Cash, &user.UserType); err != nil {
			return models.UsersResponse{}, err
		}
		users = append(users, user)

	}
	response := models.UsersResponse{
		Users: users,
		Count: len(users),
	}

	return response, nil
}

func (u userRepo) Update(users models.UpdateUser) (models.User, error) {
	_, err := u.db.Exec(`update users set fullname=$1, phone=$2,password_=$3,cash=$4 where id=$5`, users.FullName, users.Phone, users.Password, users.Cash, users.ID)
	if err != nil {
		return models.User{}, err
	}

	updateduser, err := u.GetByID(models.PrimaryKey{ID: users.ID})
	if err != nil {
		return models.User{}, err
	}
	return updateduser, nil
}

func (u userRepo) Delete(users models.PrimaryKey) error {
	if _, err := u.db.Exec(`delete from basket where id=$1`, users.ID); err != nil {
		return err
	}
	if _, err := u.db.Exec(`delete from users where id=$1`, users.ID); err != nil {
		return err
	}
	return nil

}
