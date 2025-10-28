# github.com/hsdfat/go-cli-mgt

## go-echo framework

## database

### postgresql
```
docker run -env=POSTGRES_PASSWORD=`db_password` --env=POSTGRES_USER=cli --env=POSTGRES_DB=cli_db -p 5432:5432 -d postgres:12.20-alpine
```
### pgadmin (database viewer)
```
docker run --env=PGADMIN_DEFAULT_EMAIL=`your_email` --env=PGADMIN_DEFAULT_PASSWORD=`password` -p 32443:443 -p 32080:80 -d dpage/pgadmin4:8.11
```