DROP DATABASE IF EXISTS homework_manager;

DROP USER IF EXISTS homework_manager;

SELECT 'CREATE DATABASE homework_manager' WHERE NOT EXISTS(SELECT FROM pg_database WHERE datname = 'homework_manager' )\gexec

CREATE USER homework_manager
WITH
    SUPERUSER CREATEDB CREATEROLE LOGIN PASSWORD 'homework_manager';

ALTER ROLE homework_manager SET client_encoding TO 'UTF-8';

ALTER ROLE homework_manager
SET
    default_transaction_isolation TO 'read committed';

ALTER ROLE homework_manager SET timezone TO 'UTC';

GRANT
    ALL PRIVILEGES ON DATABASE homework_manager TO homework_manager;

