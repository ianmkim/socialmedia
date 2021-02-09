package models

import (
    "github.com/Kamva/mgm/v2"
)

type User struct {
    mgm.DefaultModel `bson:",inline"`
    Name string `json:"name" bson:"name"`
    Year string `json:"year" bson:"year"`
    Picture string `json:"picture" bson:"picture"`
    Gender string `json:"gender" bson:"gender"`
    Race string `json:"race" bson:"race"`
    Major string `json:"major" bson:"major"`
    Minor string `json:"minor" bson:"minor"`
    Modification string `json:"modification" bson:"modification"`
    Birthday string `json:"birthday" bson:"birthday"`
    Role string `json:"role" bson:"role"`
    Home string `json:"home" bson:"home"`
    Quote string `json:"quote" bson:"quote"`
    FavoriteShoe string `json:"favoriteShoe" bson:"favoriteShoe"`
    FavoriteArtist string `json:"favoriteArtist" bson:"favoriteArtist"`
    FavoriteColor string `json:"favoriteColor" bson:"favoriteColor"`
    PhoneType string `json:"phoneType" bson:"phoneType"`
    Likes int `json:"likes" bson:"likes"`
}

func CreateUser(
    name,
    year,
    picture,
    gender,
    race, 
    major,
    minor,
    modification,
    birthday,
    role,
    home,
    quote,
    favoriteShoe,
    favoriteArtist,
    favoriteColor,
    phoneType string,
) *User {
    return &User {
        Name : name,
        Year : year,
        Picture: picture,
        Gender: gender,
        Race: race,
        Major: major,
        Minor: minor,
        Modification : modification,
        Birthday : birthday,
        Role : role,
        Home : home,
        Quote : quote,
        FavoriteShoe : favoriteShoe,
        FavoriteArtist : favoriteArtist,
        FavoriteColor : favoriteColor, 
        PhoneType : phoneType,
        Likes : 0,
    }
}
