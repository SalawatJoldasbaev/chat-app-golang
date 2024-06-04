new-migrate:
	@echo "Creating new migration..."
	@sql-migrate new -config=configs/dbconfig.yml -env=development $(name)
up-migrate:
	@echo "Running up migration..."
	@sql-migrate up -limit=0 -config=configs/dbconfig.yml -env=development
down-migrate:
	@echo "Running down migration..."
	@sql-migrate down -limit=$(limit) -config=configs/dbconfig.yml -env=development