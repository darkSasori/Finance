package model

// NewUser return a user
func NewUser(username, displayname, password, passwordCheck string) *User {
	return &User{
		username,
		displayname,
		[]byte(password),
		[]byte(passwordCheck),
	}
}

// User is a struct to save
type User struct {
	Username      string `bson:"_id"`
	Displayname   string
	Password      []byte
	passwordCheck []byte
}

// CheckPassword verify if 2 password as same
func (u User) CheckPassword() bool {
	return string(u.Password) == string(u.passwordCheck)
}
