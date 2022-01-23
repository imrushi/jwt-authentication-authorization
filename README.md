# JWT AUTHENTICATION AND AUTHORIZATION

### Run Postgres on Docker

`$ docker run --name postgres -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres:9.6`

### Create Database in Postgres

`$ docker exec -it postgres`

Run below commands inside container

`# su postgres`

`$ psql`

`#postgres= CREATE DATABASE userdb;`

### RUN Go Server

`$ go install`

`$ API_PORT=8090 go run main.go`

### API Endpoint

POST `http://localhost:8090/signup`

Request Body:

```
{
	"name": "test",
	"email": "test@gmail.com",
	"password":"p@ssw0rd",
	"role": "admin" // role can be "admin" or "user"
}
```

Response Body:

```
{
  "ID": 2,
  "CreatedAt": "2022-01-23T18:54:36.929113+05:30",
  "UpdatedAt": "2022-01-23T18:54:36.929113+05:30",
  "DeletedAt": null,
  "name": "test",
  "email": "test@gmail.com",
  "password": "$2a$14$iBfRa81RQclHCndi274wv.aYEtEALwkOAMCITJ0aY2QMMiJ5mXN06",
  "role": "admin"
}
```

POST `http://localhost:8090/signin`

Request Body:

```
{
	"email": "test@gmail.com",
	"password":"p@ssw0rd"
}
```

Response Body:

```
{
  "role": "admin",
  "email": "test@gmail.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNjQyOTQ0NDc0LCJyb2xlIjoiYWRtaW4ifQ.m92qOGMsluZu446JBKfZfiJw4x9XAVqIQxm5_lwc28w"
}
```

GET `http://localhost:8090/isauth`

Set Headers `Token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNjQyOTQ0NDc0LCJyb2xlIjoiYWRtaW4ifQ.m92qOGMsluZu446JBKfZfiJw4x9XAVqIQxm5_lwc28w`

Response Body:

```
{
  "msg": "Welcome, Admin."
}
```
