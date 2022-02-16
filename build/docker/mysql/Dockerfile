FROM mysql:8.0

LABEL version="v0.0.1" description="mysql-server" by="yangchnet"

COPY db/ddl /docker-entrypoint-initdb.d/
# COPY .conf/my.cnf /etc/mysql

EXPOSE 3306

