package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type AdminProfile struct {
	Login             string `json:"login"`
	EncryptedPassword string `json:"password"`
}

func (u *AdminProfile) ValidateForCreate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Login,
			validation.Required,
			validation.Length(6, 64).Error("Логин должен содержать от 6 до 64 символов")),
		validation.Field(&u.EncryptedPassword,
			validation.By(requiredIF(u.EncryptedPassword == "")),
			validation.Length(6, 32).Error("Пароль должен содержать от 6 до 32 символов")),
	)
}
