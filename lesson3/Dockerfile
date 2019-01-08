FROM postgres:11.1

ADD --chown=postgres:postgres \
    https://github.com/devrimgunduz/pagila/raw/master/pagila-schema.sql \
    /docker-entrypoint-initdb.d/01-pagila-schema.sql

ADD --chown=postgres:postgres \
    https://github.com/devrimgunduz/pagila/raw/master/pagila-data.sql \
    /docker-entrypoint-initdb.d/02-pagila-data.sql
