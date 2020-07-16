# Go Migration RnD
A research project to integrate migration tool `golang-migrate/migrate` with Gin framework, Gorm and mysql database.

## Initial Steps
- Download and Install dependency packages. 
```
go mod download
```
- Install golang-migrate CLI: `https://github.com/golang-migrate/migrate/tree/master/cmd/migrate` (preferred With Go toolchain versioned with tags mysql)
- Add a .env file to root directory and add db configurations as in .env.example file.

## Start Local Server
```
go run main.go
```
or for Hot reload

```
./run_server
```

## Create migrations
- Run `create_migration` script as 
```
./create_migration -i /path/to/migration/directory/ -n migration_file_name
```

Example:
```
./create_migration -i migrations/ -n create_books_table
```
- Write SQL statements to reflect the models changes in *migration_name*.up.sql file
- Write SQL statements to revert the database changes back to the state before migration in *migration_name*.down.sql file.


## Run migrations
The migrations are runned automatically when the server is started. However, if you want to run migrations manually you can do it by running `run_migrate` script as
```
./run_migrate -d db_name -t mysql -u db_user -p db_password -H db_host -P db_port -i /path/to/migrations/directory [-c N]
```

Example:
Run all migrations:
```
./run_migrate -d test-migration -t mysql -u root -p password -H 127.0.0.1 -P 3306 -i migrations
```

Run 2 migrations:
```
./run_migrate -d test-migration -t mysql -u root -p password -H 127.0.0.1 -P 3306 -i migrations -c 2
```

## Down database migrations
Run `down_migrate` script as 
```
./down_migrate -d db_name -t mysql -u db_user -p db_password -H db_host -P db_port -i /path/to/migrations/directory [-c N]
```

Example:
Down all migrations:
```
./down_migrate -d test-migration -t mysql -u root -p password -H 127.0.0.1 -P 3306 -i migrations
```

Down 2 migrations:
```
./down_migrate -d test-migration -t mysql -u root -p password -H 127.0.0.1 -P 3306 -i migrations -c 2
```

## Force migration
You may encounter a problem while down migrations or run migrations manually and might have to force the migration. In order to do so, run `force_migrate` script as
```
./force_migrate -d db_name -t mysql -u db_user -p db_password -H db_host -P db_port -i /path/to/migrations/directory -v version
```

Example:
Force migration version 4:
```
./force_migrate -d test-migration -t mysql -u root -p password -H 127.0.0.1 -P 3306 -i migrations -v 4
```


## TODOs
- Dockerize the server.
