package pkg



func GetUserByEmail(email string) User{

	for _, user := range Users {
		if user.Email == email {
			return user
		}
	}
	return User{}
}

func (u *User) VerifyPassword(password string) bool{
	if u.Password == password {
		return true
	}
	return false
}