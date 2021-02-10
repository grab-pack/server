# Grab Server

This is the server for the Grab Package Manager.

## Usage

For now, use the following:

```bash
docker-compose up -d
go run main.go
```

By the end of this, you can check that all migrations have been successfull.

```
user at Computer in ~/P/G/server
↪ sudo docker-compose exec db psql -U grab
# This will connect you to the PostgreSQL
# database running within the Docker Container.

psql (13.1)
Type "help" for help.

grab=# select * from maintainer;
  handle  |            name            |        email
----------+----------------------------+---------------------
 sharpvik | Viktor A. Rozenko Voitenko | sharp.vik@gmail.com
(1 row)
```
