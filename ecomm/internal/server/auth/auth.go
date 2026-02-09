package auth

type User struct {
	Name  string
	Email string
	// ... other details
}

type Authentication interface {
	Login() error
	Callback() error
	UserInfo() (*User, error)
	Logout() error
}
