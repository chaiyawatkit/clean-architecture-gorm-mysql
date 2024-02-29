# Go Clean-Architecture-Gorm-Mysql




## Manual Installation

Clone the repo:

```bash
1. git clone https://github.com/chaiyawatkit/clean-architecture-gorm-mysql.git
2. cp config.json.example config.json
3. docker-compose up  --build  -d
```



## Features

- **Golang v1.21**: Stable version of go
- **Framework**: A stable version of [gin-go](https://github.com/gin-gonic/gin)
- **Token Security**: with [JWT](https://jwt.io)
- **API documentation**: with [swaggo](https://github.com/swaggo/swag) 
    of [swagger](https://swagger.io/)
- **SQL databaseSQL**: [MariaDB](https://mariadb.org/) using internal sql package of
  go [sql](https://golang.org/pkg/databaseSQL/sql/)
- **Testing**: unit and integration tests using package of go [testing](https://golang.org/pkg/testing/)

- **Dependency management**: with [go modules](https://golang.org/ref/mod)
- **Environment variables**: using [viper](https://github.com/spf13/viper)
- **Docker support**

## Commands





### Swagger Implementation
#### install swaggo/swag

```bash
go install github.com/swaggo/swag/cmd/swag@latest

touch ~/.bash_profile;  open ~/.bash_profile

copy paste  export PATH=$PATH:$(go env GOPATH)/bin to bash_profile
```

#### Generate swagger documentation
```bash

swag init -g src/infrastructure/rest/routes/routes.go
```


#### use swagger documentation

http://localhost:8080/v1/swagger/index.html


![swagger ](https://i.postimg.cc/mkGjrYdM/Screenshot-2567-02-29-at-16-08-54.png)




