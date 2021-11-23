# Scheduler

An application that allows the checking of availability, booking, and cancelling of appointments.


# Steps To Run
1. Setup your Golang environment. These steps can be found [Here](https://go.dev/learn/)
2. Setup your Postgres environment. You can download Postgres client [Here](https://wiki.postgresql.org/wiki/PostgreSQL_Clients)
3. Import the provided Postman collection from the root application directory for ease of use endpoints
4. Create a ```.env``` file in the root project directory with the following ```USERNAME, PASSWORD, HOST, DBNAME, SSL```
5. Install the CLI for [Golang Migrate](https://github.com/golang-migrate/migrate)
6. Navigate to ```internal/db/migrations``` for the migration scripts
7. Run ```migrate -path . -database "postgres://{USERNAME}:{PASSWORD}@{HOST}:5432/{DBNAME}?sslmode=disable" up```
8. Run ```go run ./cmd/api``` in the root directory of the application
9. Use Postman with the imported collection to hit the RESTful endpoints


# Project Structure
- cmd: contains controller logic to accept and handle RESTful requests
- internal: contains all application specific business logic and data layer code
- pkg: contains all application agnostic code that is reusable

# Stack/Tools
- [Golang](https://go.dev/)
- [Postgres](https://www.postgresql.org/)
- [Golang Migrate](https://github.com/golang-migrate/migrate)
- [Viper](https://github.com/spf13/viper)




