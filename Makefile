sqlc-synchronize:
	curl -s -L https://raw.githubusercontent.com/Goboolean/shared/main/api/sql/schema.sql -o ./api/sql/schema.sql; \
	curl -s -L https://raw.githubusercontent.com/Goboolean/shared/main/api/sql/schema.test.sql -o ./api/sql/schema.test.sql; \

sqlc-generate: \
	sqlc-synchronize; \
	sqlc generate -f ./api/sql/sqlc.yml

sqlc-check: \
	sqlc-synchronize; \
	sqlc compile -f ./api/sql/sqlc.yml

all-generate: \
	sqlc-generate