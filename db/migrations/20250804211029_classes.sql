-- +goose Up
-- +goose StatementBegin
create table classes
(
    id           bigserial primary key,
    grade        varchar(10) not null unique ,
    created_at   timestamptz  not null,
    updated_at   timestamptz  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table classes;
-- +goose StatementEnd
