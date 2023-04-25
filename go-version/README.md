# Create Database

CREATE DATABASE IF NOT EXISTS homework-manager;

CREATE USER homework-manager WITH SUPERUSER CREATEDB CREATEROLE LOGIN

PASSWORD 'homework-manager';

ALTER ROLE homework-manager SET client_encoding TO 'UTF-8';

ALTER ROLE homework-manager SET default_transaction_isolation TO 'read committed';

ALTER ROLE homework-manager SET timezone TO 'UTC';

GRANT ALL PRIVILEGES DATABASE homework-manager TO 'homework-manager';
