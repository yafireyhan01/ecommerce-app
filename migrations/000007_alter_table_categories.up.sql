alter table categories
    add created_at timestamp default current_timestamp;

alter table categories
    add updated_at timestamp default current_timestamp;

alter table categories
    add deleted_at timestamp;