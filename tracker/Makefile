run:
	go run cmd/main.go
restart:
	@echo Stopping down Docker images if running...
	docker compose -f docker-compose.yml down -v
	@echo Restarting Docker images in watch mode...
	docker compose -f docker-compose.yml up --build
	@echo Docker images started!
watch:
	@echo Starting Docker images in watch mode...
	docker compose -f docker-compose.yml up
	@echo Docker images started!
watch-d:
	@echo Starting Docker images in detatch mode...
	docker compose -f docker-compose.yml up -d
	@echo Docker images started!
watch-down:
	@echo Stopping Docker images ...
	docker compose -f docker-compose.yml down -v
	@echo Docker images stopped!
prune-image:
	@echo Removing all images without at least one container associated to them...
	docker image prune -a
	@echo Docker images pruned!
prune-system:
	@echo Removing all images without at least one container associated to them...
	docker system prune -a
	@echo Docker images pruned!

ent:
	@echo Creating ent Model ...
	cd .././andurar && go run entgo.io/ent/cmd/ent new $(name) \

	@echo Model created successfully!
gen-ent:
	@echo Gnerating ent Model ...
	cd .././andurar && go generate ./ent && go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema && go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./ent/schema \
	
	@echo Model generated successfully!
upset-ent:
	@echo Gnerating ent Model ...
	cd .././andurar && go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema
	@echo Model generated successfully!