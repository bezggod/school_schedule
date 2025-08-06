-- +goose Up
-- +goose StatementBegin
create table lessons
(
    id         bigserial primary key,
    teacher_id bigint      not null,
    class_id   bigint      not null,
    office     text        not null,
    time_slot  smallint    not null,
    subject    text        not null,
    date       date        not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,


    CONSTRAINT fk_teacher
        FOREIGN KEY (teacher_id)
            REFERENCES teachers (id) ON DELETE CASCADE,
    CONSTRAINT fk_class
        FOREIGN KEY (class_id)
            REFERENCES classes (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table lessons;
-- +goose StatementEnd
