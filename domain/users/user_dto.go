package users

type User struct {
	Id        int64  `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	CreatedAt string `json:"created_at" validate:"required"`
	Status    string `json:"status"`
	Password  string `json:"-"`
}
