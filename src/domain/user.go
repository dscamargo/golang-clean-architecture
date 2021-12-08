package domain

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"-" bson:"-"`
}

type UserRepository interface {
	FindById(id string) (User, bool, error)
	FindByEmail(email string) (User, bool, error)
	Create(name, email, password string) (string, error)
}
