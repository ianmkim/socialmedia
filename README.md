# socialmedia
Social Media backend for DALI lab written in glorious Golang.

# Installation and usage
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
