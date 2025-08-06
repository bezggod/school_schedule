migrate-up:
	goose -dir ./db/migrations postgres "host=localhost user=bezgo dbname=school_schedule_2 password= sslmode=disable" up

migrate-create:
	goose -dir db/migrations create $(F) sql
