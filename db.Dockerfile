FROM mysql:8.0.23

COPY ./db/migrations/*.sql /docker-entrypoint-initdb.d/