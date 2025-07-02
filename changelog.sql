--liquibase formatted sql

--changeset dan.hart:1 labels:tasks
--comment: create table for tasks
CREATE TABLE tasks
(
    id   INTEGER GENERATED ALWAYS AS IDENTITY,
    task VARCHAR NOT NULL
)
--rollback DROP TABLE tasks;
