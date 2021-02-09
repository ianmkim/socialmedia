package models

import (
    "github.com/Kamva/mgm/v2"
)

type AnonUser struct {
    mgm.DefaultModel `bson:",inline"`
    Year string `json:"year" bson:"year"`
    Gender string `json:"gender" bson:"gender"`
    HeightInches int `json:"heightInches" bson:"heightInches"`
    Happiness int `json:"happiness" bson:"happiness"`
    Stressed int `json:"stressed" bson:"stressed"`
    SleepPerNight int`json:"sleepPerNight" bson:"sleepPerNight"`
    SocialDinnerPerWeek int`json:"socialDinnerPerWeek" bson:"socialDinnerPerWeek"`
    AlcoholDrinksPerWeek int`json:"alcoholDrinksPerWeek" bson:"alcoholDrinksPerWeek"`
    CaffeineRating int`json:"caffeineRating" bson:"caffeineRating"`
    Affiliated int`json:"affiliated" bson:"affiliated"`
    NumOfLanguages int`json:"numOfLanguages" bson:"numOfLanguages"`
    GymPerWeek int`json:"gymPerWeek" bson:"gymPerWeek"`
    HoursOnScreen int`json:"hoursOnScreen" bson:"hoursOnScreen"`
    PhoneType string `json:"phoneType" bson:"phoneType"`
}

func CreateAnonUser (
    year,
    gender string,
    heightInches,
    happiness,
    stressed, 
    sleepPerNight,
    socialDinnerPerWeek,
    alcoholDrinksPerWeek,
    caffeineRating,
    affiliated,
    numOfLanguages,
    gymPerWeek,
    hoursOnScreen int,
    phoneType string,
) *AnonUser {
    return &AnonUser {
        Year : year,
        Gender : gender,
        HeightInches : heightInches,
        Happiness : happiness,
        Stressed : stressed,
        SleepPerNight : sleepPerNight,
        SocialDinnerPerWeek : socialDinnerPerWeek,
        AlcoholDrinksPerWeek : alcoholDrinksPerWeek,
        CaffeineRating : caffeineRating,
        Affiliated : affiliated,
        NumOfLanguages : numOfLanguages,
        GymPerWeek : gymPerWeek,
        HoursOnScreen  : hoursOnScreen,
        PhoneType : phoneType,
    }
}
