package model

type UserResponse struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type RegisterUser struct {
	Name     string `json:"name" validate:"required,max=50"`
	Password string `json:"password" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,max=100"`
	NoHp     string `json:"no_hp" validate:"required,max=100"`
}
