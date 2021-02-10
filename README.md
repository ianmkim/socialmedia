# socialmedia
Social Media backend for DALI lab written in glorious Golang.

# Installation
You need golang and all necessary dependencies installed.
Clone this repo
```
git clone https://github.com/parvusvox/socialmedia.git
cd socialmedia
```

add database connection string in "localenv" then set environmental variables
```
source localenv
```
Run this project
```
go run server.go
```
... or with air
```
air .
```

# Upload data to database
navigate to the pytools directory and run script
```
source localenv
cd pytools
python3 upload.py
```

# Upload to heroku
This repo comes with a Dockerfile you can use with heroku, just replace the <YOUR APP NAME> with your app name on heroku
```
heroku container:push web --app <YOUR APP NAME> && heroku container:release --app <YOUR APP NAME>
```

# Authentication
This backend uses Javascript Web Tokens (JWT). h
# API Usage
```
GET /users?page=<int>&limit=<int>
BODY empty
RETURNS 
{
  "ok": boolean,
  "users": [
    {
        _id : string,
        created_at : string,
        "updated_at": string,
        "name": string,
        "year": string,
        "picture": string,
        "gender": string,
        "race": string,
        "major": string,
        "minor": string,
        "modification":string,
        "birthday": string,
        "role": string,
        "home": string,
        "quote": string,
        "favoriteShoe": string,
        "favoriteArtist": string,
        "favoriteColor": string,
        "phoneType": string,
        "likes": int
    },...
  ]
}
```
```
POST /user
BODY 
{
  "name": string,
  "year": string,
  "picture": string,
  "gender": string,
  "race": string,
  "major": string,
  "minor": string,
  "modification":string,
  "birthday": string,
  "role": string,
  "home": string,
  "quote": string,
  "favoriteShoe": string,
  "favoriteArtist": string,
  "favoriteColor": string,
  "phoneType": string,
  "likes": int
}
RETURNS
{
  "id": string,
  "ok": boolean
}
```

```
GET /user/<id>
BODY empty
RETURNS
{
  "ok" : boolean,
  "user": {
    "name": string,
    "year": string,
    "picture": string,
    "gender": string,
    "race": string,
    "major": string,
    "minor": string,
    "modification":string,
    "birthday": string,
    "role": string,
    "home": string,
    "quote": string,
    "favoriteShoe": string,
    "favoriteArtist": string,
    "favoriteColor": string,
    "phoneType": string,
    "likes": int
  }
}
```

```
GET /index
BODY empty 
RETURNS 
{
  "ok" : boolean,
}
```

```
GET /searchQuotes?q=<query>
BODY empty
RETURNS
{
  "ok": boolean,
  "length": int,
  "users": [
    {
      "name": string,
      "year": string,
      "picture": string,
      "gender": string,
      "race": string,
      "major": string,
      "minor": string,
      "modification":string,
      "birthday": string,
      "role": string,
      "home": string,
      "quote": string,
      "favoriteShoe": string,
      "favoriteArtist": string,
      "favoriteColor": string,
      "phoneType": string,
      "likes": int
    }
  ]
}
```
