CREATE TABLE schema_migrations (
                                   version BIGINT NOT NULL,
                                   dirty BOOLEAN NOT NULL,
                                   PRIMARY KEY (version)
);
