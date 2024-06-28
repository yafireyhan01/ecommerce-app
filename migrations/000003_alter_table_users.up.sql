alter table users
    add created_at timestamp default current_timestamp;

alter table users
    add updated_at timestamp default current_timestamp;

alter table users
    add deleted_at timestamp;
