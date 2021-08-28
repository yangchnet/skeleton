sqlc:
	sqlc generate

migrate-up: 
	migrate -source=file://schema/iam -database "mysql://root:admin123@tcp(localhost:13306)/iam" -verbose up
	migrate -source=file://schema/hello -database "mysql://root:admin123@tcp(localhost:13306)/hello" -verbose up

migrate-down:
	migrate -source=file://schema/iam -database "mysql://root:admin123@tcp(localhost:13306)/iam" -verbose down
	migrate -source=file://schema/hello -database "mysql://root:admin123@tcp(localhost:13306)/hello" -verbose down

sql-build:
	docker build -t skeleton-mysql .

sql-up:
	docker run --name skeleton-sql --restart=always -p 13306:3306 -v ~/.skeleton-data/mysql:/var/lib/mysql -d skeleton-mysql 

sql-down:
	docker stop skeleton-sql && \
	docker rm skeleton-sql
