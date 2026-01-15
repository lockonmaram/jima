#!/bin/sh
set -eu

echo "============================================================================================"
echo "Creating role ${JIMA_POSTGRES_DB_USER} and database ${JIMA_POSTGRES_DB_NAME}"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
-- Check if role exists, create or alter as needed
SELECT 'CREATE ROLE "${JIMA_POSTGRES_DB_USER}" LOGIN PASSWORD ''${JIMA_POSTGRES_DB_PASSWORD}'';'
WHERE NOT EXISTS (SELECT FROM pg_roles WHERE rolname = '${JIMA_POSTGRES_DB_USER}')
\gexec

SELECT 'ALTER ROLE "${JIMA_POSTGRES_DB_USER}" WITH LOGIN PASSWORD ''${JIMA_POSTGRES_DB_PASSWORD}'';'
WHERE EXISTS (SELECT FROM pg_roles WHERE rolname = '${JIMA_POSTGRES_DB_USER}')
\gexec

-- Check if database exists, create if not
SELECT 'CREATE DATABASE "${JIMA_POSTGRES_DB_NAME}" OWNER "${JIMA_POSTGRES_DB_USER}";'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '${JIMA_POSTGRES_DB_NAME}')
\gexec
EOSQL

echo "Role ${JIMA_POSTGRES_DB_USER} and database ${JIMA_POSTGRES_DB_NAME} created or already exist"
echo "============================================================================================"