# golang-graphql


### Start mysql in local
```
$ mysql -u root -p
```

### generate resolver based on latest schema file
```
$ go run github.com/99designs/gqlgen generate
```


### set local env variables
```
$ export DBUSER=username
$ export DBPASS=password
```


### go to project root directory and run server
```
$ go run .
```
