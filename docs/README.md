# synapsis-test Backend Challenge

This is a backend app RESTful API Challenge from Synapsis

## Minimum viable product
* Customer can view product list by product category
* Customer can add product to shopping cart
* Customers can see a list of products that have been added to the shopping cart
* Customer can delete product list in shopping cart
* Customers can checkout and make payment transactions
* Login and register customers

## System Design : ERD

Here is the the erd link: https://dbdiagram.io/d/Synapsis-64d2361d02bd1c4a5e6e5d50
![Alt text](docs/erd.png?raw=true "Entity Relationship Diagram")

## API docummentation
Here's a link to the postman documentation:
https://documenter.getpostman.com/view/23469031/2sA3duGZCal

## Prerequisite run on local

### Set Environment Variable
```bash
cp app.env.example app.env
```
### Set Database Account
DB_HOST=localhost
DB_PORT=5432
DB_NAME=synapsis_db
DB_USER=postgres
DB_PASSWORD=

### Database migration
migrate https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

#### Export url
```bash
export POSTGRESQL_URL='postgres://postgres:root@localhost:5432/synapsis_db?sslmode=disable'
```

#### Migrate UP
```bash
migrate -database ${POSTGRESQL_URL} -path migrations up
```

### Migrate New Scheme
```bash
make migrate_new name=create_table_namatable
```
### ORM Builder
https://gorm.io/
## RUN SERVER
```bash
make serve 
```


