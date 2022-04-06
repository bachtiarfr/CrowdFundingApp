package user

type Repository interface {
	saveUser(user User) (User, error)
}