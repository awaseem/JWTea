<p align="center">
  <img src="http://i.imgur.com/3uwThuK.png)">
</p>

# JWTea

[![Go Report Card](https://goreportcard.com/badge/github.com/awaseem/jwtea)](https://goreportcard.com/report/github.com/awaseem/jwtea)

Simple HTTP service for user auth with json web tokens.

## Create

To create a user simply post to `/create` with a username and password, JWTea will parse the information and generate
a new user for you.

```
POST /create

Body {
  "username": "awaseem",
  "password": "shhh!"
}
```

## Login

To login post to `/login` with the username and password of an existing user, JWTea will parse the username and password and send you back 
the information.

```
POST /decode

Body {
  "username": "awaseem",
  "password": "shhh!"
}
```

## Check 

To check a token to see if it's valid, post to `/check` and JWTea will send a 200 if success otherwise it failed!

```
POST /check

Body {
  "token": "sdfjsdfjdsfkljsdfkldsjlksfjsdf"
}
```
