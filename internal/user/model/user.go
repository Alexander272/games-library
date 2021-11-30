package models

type Token struct {
	AccessToken string `json:"accessToken"`
}

type User struct {
	Id       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type SignInUserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func NewUser(dto CreateUserDTO) User {
	return User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     dto.Role,
	}
}

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=128"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
	Role     string `json:"role" binding:"required"`
}

func UpdateUser(dto UpdateUserDTO) User {
	return User{
		Id:       dto.Id,
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.NewPassword,
		Role:     dto.Role,
	}
}

type UpdateUserDTO struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	Role        string `json:"role"`
}
