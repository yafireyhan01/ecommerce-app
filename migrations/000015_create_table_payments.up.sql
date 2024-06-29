ALTER TABLE checkouts
    ADD CONSTRAINT unique_checkout_guid UNIQUE (guid);

CREATE TABLE payments
(
    "id"   INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    "guid" varchar(36) NOT NULL,
    "checkout_guid" varchar(36) NOT NULL,
    "payment_date" TIMESTAMP DEFAULT current_timestamp,
    "created_at" timestamp default current_timestamp,
    "updated_at" timestamp default current_timestamp,
    "deleted_at" timestamp,
    CONSTRAINT fk_payment_checkout FOREIGN KEY (checkout_guid) REFERENCES checkouts(guid)
);