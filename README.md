# interview-app-srv
[![Go Report Card](https://goreportcard.com/badge/github.com/RafalSalwa/interview-app-srv)](https://goreportcard.com/report/github.com/RafalSalwa/interview-app-srv)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/7285/badge)](https://bestpractices.coreinfrastructure.org/projects/7285)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/RafalSalwa/interview-app-srv/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/RafalSalwa/interview-app-srv/tree/master)
[![Go](https://github.com/RafalSalwa/interview-app-srv/actions/workflows/go.yml/badge.svg)](https://github.com/RafalSalwa/interview-app-srv/actions/workflows/go.yml)

Interview app
Sample app presenting app structure and usage of some libraries. This app works with REST and gRPC servers included, 
based on docker containers. 

create network for containers in other repos
```bash
docker network create --driver=bridge --subnet=172.28.1.0/24 --ip-range=172.28.1.0/24 --gateway=172.28.1.1 external-interview-net
```
## Build
make up

## Run tests

make test


### build
- make proto
- make build
## Technologies used
 
 - files structure according to [best practices](https://github.com/golang-standards/project-layout)
 - gorilla/mux middleware and handlers
   - [correlationid](api/resource/middlewares/correlationid.go)
   - [cors](api/resource/middlewares/cors.go)
   - [Basic Auth](internal/auth/auth.go) 
   - jwt auth
   - api key
 - mysql:
   - [gorm](internal/services/user_gorm_service.go)
   - [database/sql](internal/services/user_sql_service.go)
 - jwt [generation and validation](internal/jwt)
 - services implementations
   - [services](internal/services)
   - [repository](internal/repository)
 - gRPC
   - [getUser](rpc_api/rpc_get_user.go)
   - [AuthUser](rpc_api/rpc_auth_user.go)
 - [Swagger](api/resource/swagger)
 - Docker with Docker compose
   - [Dockerfile](Dockerfile)
   - [docker-compose](docker-compose.yml)
 - [Postman collection](docs/Interview_API_Requests.postman_collection.json)

# In Progress
- notification hub -> api for email sendout in batches/bulks
- tests


# GRPC
usage example with responses
```bash
grpcurl -plaintext 0.0.0.0:8082 list
>
intrvproto.AuthService
intrvproto.UserService
```
get user info
``` bash
grpcurl -d '{"user_id":"1"}' -plaintext 0.0.0.0:8082 intrvproto.UserService/GetUserById
{
  "user": {
    "id": "1",
    "username": "user1",
    "createdAt": "2023-05-01T07:16:12Z"
  }
}

```

Get token and Refresh token
``` bash
grpcurl -d '{"username":"user1", "password":"password"}' -plaintext 0.0.0.0:8082 intrvproto.AuthService/SignInUser
{
  "status": "success",
  "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM4MzQ5NzYsImlhdCI6MTY4MzgzNDA3NiwibmJmIjoxNjgzODM0MDc2LCJzdWIiOjF9.B0XjeFLR0nrpKU00ga1XFR4ASRoEyVhby4RecEccQPyyxkdtj9xPeqBACFp_D0P4HZJXpQ3SRPyqoos2hSH9Lw",
  "refreshToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM4Mzc2NzYsImlhdCI6MTY4MzgzNDA3NiwibmJmIjoxNjgzODM0MDc2LCJzdWIiOjF9.b8dyhMSc3yApI0z7iftyJipqQeZWa4KB0O5Y9YZXVIk0rjsd8Z23PrkvvByjKWmLiRzY413httmPTtu4vpobuA"
}

```

# Rest API
params:
- apikey: 
  - .env['X_API_KEY']
- Basic Auth:
  - .env['AUTH_USERNAME']
  - .env['AUTH_PASSWORD']

example user:
   - user1:password
routes:

GET / Index page with routes list

GET /user/{id} get User details

POST /user create user

POST /user/change_password self explanatory

GET /auth/login call to get tokens

## Screenshots:
![grpc_list.png](docs/grpc_list.png)

---

![get_user.png](docs/get_user.png)

---

![grpc_auth.png](docs/grpc_auth.png)
