--liquibase formatted sql

--changeset dan.hart:1 labels:tasks
--comment: create table for tasks
create table tasks
(
    id   integer generated always as identity,
    task varchar not null
)
--rollback DROP TABLE tasks;
