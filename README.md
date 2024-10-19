# Golang links shortener

## Run the app first time

You can run the app with the command following:
```bash
make init
```
That will compile & runs web server and proceed migrations.

To stop app use:
```bash
make down
```
To get into container terminal:
```bash
make exec
```

## Migrations
Run the migrations with command:

```bash
make migrate
```

Create new migration with command:
```bash
make migration
```