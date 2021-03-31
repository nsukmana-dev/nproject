package user

type UserFormmater struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormmater {
	formatter := UserFormmater{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Role:       user.Role,
		Token:      token,
	}

	return formatter
}
