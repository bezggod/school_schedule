-- +goose Up
-- +goose StatementBegin
create table teachers
(
    id         bigserial primary key,
    surname    varchar(10) not null,
    name       varchar(36) not null,
    patronymic varchar(36) ,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table teachers;
-- +goose StatementEnd
