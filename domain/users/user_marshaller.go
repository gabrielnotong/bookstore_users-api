package users

import "encoding/json"

type PublicUser struct {
	Id        int64  `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type PrivateUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	uu := make([]interface{}, len(users))
	for i, user := range users {
		uu[i] = user.Marshall(isPublic)
	}

	return uu
}

func (u *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:        u.Id,
			Status:    u.Status,
			CreatedAt: u.CreatedAt,
		}
	}

	uj, _ := json.Marshal(u)
	var pu PrivateUser
	_ = json.Unmarshal(uj, &pu)

	return pu
}
