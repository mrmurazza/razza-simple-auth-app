package response

import (
	"dealljobs/domain/user"
	"time"
)

type (
	UserResponse struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		Address   string `json:"address"`
		Role      string `json:"role"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
)

func FromUserToResponse(u *user.User) *UserResponse {
	return &UserResponse{
		ID:        u.Id,
		Username:  u.Username,
		Name:      u.Name,
		Address:   u.Address,
		Role:      string(u.Role),
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}
