# Myservice

## Overview
This is simple app to add and get "users"<br>
**Backend:** Golang with few frameworks: [go-macaron](https://go-macaron.com), [sirupsen/logrus](https://github.com/sirupsen/logrus), [jmoiron/sqlx](https://github.com/jmoiron/sqlx) <br>
**Frontend:** Bootstrap 4.3.1 + jquery 3.6.0 <br>
**Build:** Go binary / Docker <br>
**CI/CD:** GitHub actions <br>
**Environment:**

### How to start
1. Specify environment variables - `APP_VERSION`, `DB_URL`, `DB_USERNAME`, `DB_PASSWORD`, `DOCKER_REGISTRY_URL`
```shell
#example values:
APP_VERSION=0.2.0
DB_URL=localhost:3306
DB_USERNAME=root
DB_PASSWORD=supersecret
```
2. Start mysql instance
```shell
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} -v ${PWD}:/var/lib/mysql -p 3306:3306 mysql:5.7.33 --innodb_rollback_on_timeout=1
```
3. Build and run
```shell
docker build -t myservice:${APP_VERSION} .
docker run --name myservice -p 80:4000 myservice:${APP_VERSION}
```
4. Send GET / POST requests to check
```shell
# create user example
curl -XPOST http://localhost/api/v1/users -H "Content-Type: application/json" -d '{"name":"test2","email":"test@mail.com"}'
# get user by name example
curl -XGET http://localhost/api/v1/users?name=test2
```

## Contributing

Thanks for considering contributing! There’s information about how to [get started with Myservice here](CONTRIBUTE.md).

When PR will be created - new AWS est environment will be ready in 20 minutes. You will get link to test web service from github-bot
in PR comments.

After closing PR by ny reason - test environment will be destroyed

## License

[The MIT License (MIT)](LICENSE.md)

Copyright © 2021 [Denis Lavrushko](https://dlavrushko.de)