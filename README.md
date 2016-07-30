<p align="center">
  <img src="http://i.imgur.com/3uwThuK.png)">
</p>

# JWTea

Simple HTTP service to create and parse JWT tokens.

## Create

To create a token simply post to `/create` with anything you'd like, JWTea will parse the information and generate
a token for you to use.

```
POST /create

Body {
  "username": "awaseem",
  "userPath": "/"
}
```

## Decode

To decode a token simply post to `/decode` with the token you'd like, JWTea will parse the token and send you back 
the information.

```
POST /decode

Body {
  "token": "j389rfnf487bgaiw4g.398f9unr9w8gn.904fuin"
}
```