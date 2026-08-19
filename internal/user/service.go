package user

type UserService struct{
	Repository *UserRepository 
}

func NewUserService(repository *UserRepository) *UserService{
	return &UserService{
		Repository: repository,
	}
}

func (s *UserService) CreateUser(email string, password string, userName string) (*User, error){
	newUser := User{
		Email: email,
		Password: password,
		UserName: userName,
	}

	_ = s.Repository.Create(&newUser)

	return nil, nil

}
