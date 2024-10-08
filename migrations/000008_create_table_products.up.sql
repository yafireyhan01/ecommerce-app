ALTER TABLE categories
    ADD CONSTRAINT unique_guid UNIQUE (guid);

CREATE TABLE "products"
(
    "id"   INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "guid" varchar(36) NOT NULL,
    "category_guid" varchar(36) NOT NULL,
    "name" varchar(255) DEFAULT null,
    "description" text,
    "price" decimal(10,2) NOT NULL,
    "stock_qty" int NOT NULL,
    "created_at" timestamp default current_timestamp,
    "updated_at" timestamp default current_timestamp,
    "deleted_at" timestamp,
    CONSTRAINT fk_category_product FOREIGN KEY (category_guid) REFERENCES categories(guid)
)