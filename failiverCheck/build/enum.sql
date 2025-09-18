-- Active: 1757856792238@@127.0.0.1@5434@db
CREATE TYPE enum_status AS ENUM(
        'DRAFT',
        'DELETED',
        'COMPLITED',
        'FORMED',
        'REJECTED'
);