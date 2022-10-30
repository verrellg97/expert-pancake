# Microservice projects using GO

## Stacks used in this project

- Database migration versioning using sqitch
- Postgresql for the database

## Steps to setup the app

- Run the docker containers (see How to run the app)
- Create database for each services if not exists
- Run migration for each database for each available services

```Sqitch
// template
sqitch deploy db:pg://{db_username}:{db_password}@{db_host}/{db_name} --verify

// example
sqitch deploy db:pg://postgres:changeme@localhost:6000/account --verify
```

- Check if each services is active, if not check docker ps and logs

## How to run the app

```Makefile
make up_build
```
