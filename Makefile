build:
	docker compose build
up:
	docker compose up
down:
	docker compose down --remove-orphans

#TODO: add one database server and create multiple databases
#TODO: add wait for it script
#TODO: add env file for each service