FROM mysql:8.0

LABEL version="v0.0.1" description="mysql-server" by="yangchnet"

COPY build/docker/mysql/ddl /docker-entrypoint-initdb.d/

EXPOSE 3306

