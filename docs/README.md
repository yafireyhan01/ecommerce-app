# ecommerce app

This is a core of ecommerce system 
BUILD WITH GOLANG

## Minimum viable product
* Customer can view product list by product category
* Customer can add product to shopping cart
* Customers can see a list of products that have been added to the shopping cart
* Customer can delete product list in shopping cart
* Customers can checkout and make payment transactions
* Login and register customers

## Prerequisite run on local

### Database migration
migrate https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


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


