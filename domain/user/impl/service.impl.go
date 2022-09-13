package impl

import (
	"dealljobs/domain/user"
	"dealljobs/dto/request"
	"dealljobs/pkg/auth"
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
		Password: auth.EncryptPassword(req.Password),
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
	u, err := s.repo.GetUser(req.ID)
	if u == nil {
		return errors.New("user Not Found")
	}
	if err != nil {
		return err
	}

	role, err := s.ConvertToRole(req.Role)
	if err != nil {
		return err
	}

	u.Username = req.Username
	u.Name = req.Name
	u.Password = auth.EncryptPassword(req.Password)
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
	u, err := s.repo.GetUser(id)
	if err != nil {
		return nil
	}

	return u
}

func (s *service) GetAllUsers() []*user.User {
	return s.repo.GetAllUsers()
}

func (s *service) Login(username, password string) (*user.User, string, error) {
	encryptedPass := auth.EncryptPassword(password)
	u, err := s.repo.GetUserByUserPass(username, encryptedPass)
	if u == nil {
		return nil, "", nil
	}
	if err != nil {
		return nil, "", err
	}

	token, err := auth.TokenizeData(u)
	if err != nil {
		return nil, "", err
	}

	return u, token, nil
}
