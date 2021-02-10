package models

import (
    "github.com/Kamva/mgm/v2"
)

type LoginUser struct {
    mgm.DefaultModel `bson:",inline"`
    Username string `json:"username" bson:"username"`
    Email string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
}

func CreateLoginUser(username, email, password string) *LoginUser{
    return &LoginUser {
        Username : username,
        Email : email,
        Password : password,
    }
}


