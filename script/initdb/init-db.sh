set -euo pipefail
echo "Creating role ${APP_DB_USER} and database ${APP_DB_NAME}"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
-- Check if role exists, create or alter as needed
SELECT 'CREATE ROLE "${APP_DB_USER}" LOGIN PASSWORD ''${APP_DB_PASSWORD}'';'
WHERE NOT EXISTS (SELECT FROM pg_roles WHERE rolname = '${APP_DB_USER}')
\gexec

SELECT 'ALTER ROLE "${APP_DB_USER}" WITH LOGIN PASSWORD ''${APP_DB_PASSWORD}'';'
WHERE EXISTS (SELECT FROM pg_roles WHERE rolname = '${APP_DB_USER}')
\gexec

-- Check if database exists, create if not
SELECT 'CREATE DATABASE "${APP_DB_NAME}" OWNER "${APP_DB_USER}";'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '${APP_DB_NAME}')
\gexec
EOSQL

echo "Role ${APP_DB_USER} and database ${APP_DB_NAME} created or already exist"