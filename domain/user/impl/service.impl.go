package impl

import (
	"dealljobs/domain/request"
	"dealljobs/domain/user"
	"dealljobs/util"
	"errors"
)

type service struct {
	repo user.Repository
}

func NewService(repo user.Repository) user.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) ConvertToRole(roleStr string) (user.Role, error) {
	roleMap := user.GetAllRoles()

	role := user.Role(roleStr)
	if _, ok := roleMap[role]; !ok {
		return "", errors.New("role does not exist")
	}

	return role, nil
}

func (s *service) CreateUserIfNotAny(req request.CreateUserRequest) (*user.User, error) {
	role, err := s.ConvertToRole(req.Role)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Username: req.Username,
		Name:     req.Name,
		Password: util.EncryptPassword(req.Password),
		Address:  req.Address,
		Role:     role,
	}

	existingUser, err := s.repo.GetUserByUsername(*u)
	if existingUser != nil {
		return existingUser, nil
	}

	u, err = s.repo.Persist(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) UpdateUser(req request.UpdateUserRequest) error {
	u := s.repo.GetUser(req.ID)
	if u == nil {
		return errors.New("user Not Found")
	}

	role, err := s.ConvertToRole(req.Role)
	if err != nil {
		return err
	}

	u.Username = req.Username
	u.Name = req.Name
	u.Password = util.EncryptPassword(req.Password)
	u.Address = req.Address
	u.Role = role

	err = s.repo.Update(u)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteUser(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUser(id int) *user.User {
	return s.repo.GetUser(id)
}

func (s *service) GetAllUsers() []*user.User {
	return s.repo.GetAllUsers()
}

func (s *service) Login(username, password string) (*user.User, string, error) {
	encryptedPass := util.EncryptPassword(password)
	u := s.repo.GetUserByUserPass(username, encryptedPass)
	if u == nil {
		return nil, "", nil
	}

	token, err := util.TokenizeData(u)
	if err != nil {
		return nil, "", err
	}

	return u, token, nil
}
