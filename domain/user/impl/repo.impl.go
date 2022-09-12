package impl

import (
	"dealljobs/domain/user"
	"github.com/jinzhu/gorm"
	"time"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) user.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Persist(u *user.User) (*user.User, error) {
	res := r.db.Save(u)

	err := res.Error
	if err != nil {
		return nil, err
	}

	res.Last(u)

	return u, nil

}

func (r *repo) Update(u *user.User) error {
	err := r.db.Model(user.User{}).Where("id = ?", u.ID).
		Update(u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(id int) error {
	err := r.db.Model(user.User{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"deleted_at": time.Now(),
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetUser(id int) (*user.User, error) {
	u := user.User{}
	err := r.db.Model(user.User{}).Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	return &u, nil
}

func (r *repo) GetUserByUserPass(username, password string) (*user.User, error) {
	u := user.User{}
	err := r.db.Model(user.User{}).Where("username = ? AND password = ?", username, password).First(&u).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	return &u, nil
}

func (r *repo) GetUserByUsername(i user.User) (*user.User, error) {
	u := user.User{}
	err := r.db.Model(user.User{}).Where("username = ?", i.Username).First(&u).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	return &u, nil
}

func (r *repo) GetAllUsers() (list []*user.User) {
	err := r.db.Model(user.User{}).Find(&list).Error
	if err != nil {
		return list
	}

	return list
}
