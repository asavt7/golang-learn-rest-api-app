## db

```
docker run --name=todo-db -e POSTGRES_PASSWORD=password -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 --rm postgres
```

## migrate

https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

``` 
migrate create -ext sql -dir ./schema -seq init

# fill files with sql scripts

migrate -path ./schema/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' up


# check migration result
# connect to docker container
docker exec -it $(docker ps -q) /usr/bin/psql -U postgres
# list tables
\d

```

fix dirty migration
```
#connect to docker container and run psql
docker exec -it $(docker ps -q) /usr/bin/psql -U postgres

#fix table row (000001 - migration version )
update schema_migrations set version='000001', dirty=false;

```


revert migration 
```
migrate -path ./schema/ -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' down
```

