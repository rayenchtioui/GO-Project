# Go project
## Clone the project:

`
git clone https://github.com/rayenchtioui/GO-Project.git
`
## Mysql database environment setup, execute the following commands:

`
sudo mysql -u root -p
`

`
create database quanticfy_test;
`

`
create database quanticfy_test;
`

`
create user 'candidat2020'@'127.0.0.1' indentified by 'dfskj_878$*=';
`

`
grant all on quanticfy_test.* to 'candidat2020'@'localhost';
`

`
flush privileges
`

## Add the following variables to the .env file:
`
DB_HOST=localhost 
`

`
DB_PORT=3306
`

`
DB_USER=candidat2020
`

`
DB_PASSWORD=candidat
`

`
DB_NAME=quanticfy_test
`

## Run the project:

`
cd project/location
`

`
go build && ./go-project
`

## Project structure:
csv package: random data generated.

data package: functions to generate random data.

pkg/:

- database package: DB connection, data insertion and schemas.
- exporter package: exports results of the calculations.
- model package: database tables representation in struct.
- processing: revenue calculations and quantile analysis.

main.go: entry file of the project.
