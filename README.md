# gochallenge

1. Prerequisites: 

- Install Go version 1.10.3
- Install PostgreSQL 10.4 version for windows (we use "postgres" for default user and P@ssw0rd as default password)
- Add postgresql instance to your PATH. In a command line execute:

  set PATH=%PATH%;C:\Program Files\PostgreSQL\10\bin

2. Get dependencies. 

On a windows machine, open a command line and run these commands:

cd %GOPATH%\src

go get github.com/adolfovega/gochallenge

go get github.com/gorilla/mux

go get github.com/lib/pq

go get gopkg.in/gin-gonic/gin.v1


In case you are unable to download above dependencies, check your proxy settings.

3. Create the database, 

psql -U postgres -c "CREATE DATABASE tododb WITH OWNER postgres;"

4. Run the app

go run main.go

