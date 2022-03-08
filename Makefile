.PHONY: sql createdb dropdb migrationsup migrationsdown server test

sql:
	sqlc generate

generatedbdoc: 
	echo 'to be done'
	# npm install -g @softwaretechnik/dbml-renderer
	# dbml-renderer -i db/schema.dbml -o output.svg


createdb:
	echo 'to be done'

dropdb:
	echo 'to be done'

migrationsup:
	echo 'running migrations up'

migrationsdown:
	echo 'running migrations down'

server:
	go run main.go

test:
	echo 'to be done'