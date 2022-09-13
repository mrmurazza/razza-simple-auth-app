package request

import "errors"

type (
	CreateUserRequest struct {
		Username             string `json:"username"`
		Name                 string `json:"name"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
		Address              string `json:"address"`
		Role                 string `json:"role"`
	}

	UpdateUserRequest struct {
		ID                   int    `json:"id"`
		Username             string `json:"username"`
		Name                 string `json:"name"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
		Address              string `json:"address"`
		Role                 string `json:"role"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (cr CreateUserRequest) Validate() error {
	if cr.Username == "" {
		return errors.New("username is required")
	}

	if cr.Password != cr.PasswordConfirmation {
		return errors.New("password & password confirmed must be the same")
	}

	if cr.Role == "" {
		return errors.New("role is required")
	}

	return nil
}

func (ur UpdateUserRequest) Validate() error {
	if ur.ID == 0 {
		return errors.New("ID is required")
	}

	if ur.Username == "" {
		return errors.New("username is required")
	}

	if ur.Password != ur.PasswordConfirmation {
		return errors.New("password & password confirmed must be the same")
	}

	if ur.Role == "" {
		return errors.New("role is required")
	}

	return nil
}
