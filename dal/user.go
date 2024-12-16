package dal

import (
	"chat/dal/dao"
	"chat/dal/model"
	"context"
	"errors"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var UserRepo *userRepo

type userRepo struct{}

func (u *userRepo) IsNameExist(ctx context.Context, name string, q ...*dao.Query) (bool, error) {
	tx := GetTx(q)
	cond := []gen.Condition{
		tx.User.Name.Eq(name),
		tx.User.IsDeleted.Eq(0),
	}
	data, err := tx.User.WithContext(ctx).Where(cond...).First()
	if err != nil || data == nil {
		return false, err
	}
	return true, nil
}

func (u *userRepo) CreateUser(ctx context.Context, name string, password string, q ...*dao.Query) error {
	tx := GetTx(q)
	val := &model.User{
		Name:     name,
		Password: password,
	}
	err := tx.User.WithContext(ctx).Create(val)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) Login(ctx context.Context, name string, password string, q ...*dao.Query) (*model.User, error) {
	tx := GetTx(q)
	cond := []gen.Condition{
		tx.User.Name.Eq(name),
		tx.User.Password.Eq(password),
		tx.User.IsDeleted.Eq(0),
	}
	data, err := tx.User.WithContext(ctx).Where(cond...).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return data, nil
}

func (u *userRepo) GetUserByID(ctx context.Context, id int64, q ...*dao.Query) (*model.User, error) {
	tx := GetTx(q)
	cond := []gen.Condition{
		tx.User.ID.Eq(id),
		tx.User.IsDeleted.Eq(0),
	}
	data, err := tx.User.WithContext(ctx).Where(cond...).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return data, nil
}
