* Create mysql database with db username and db password as `env/config.sh`
* import sql/token.sql to database
* Run `source env/config.sh && go run cmd/api/main.go`
* Test with api `http://localhost:8080/v1/tokenPrice`,... (api in server/server.go)
