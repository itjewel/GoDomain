package user

type UserRepository interface {
	GetByID(id string) (*User, error)
	Save(u *User) error
}
