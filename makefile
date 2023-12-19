migrations:
	sql-migrate new ${name}

migrate-up:
	sql-migrate up

migrate-down:
	sql-migrate down

migrate-redo:
	sql-migrate redo