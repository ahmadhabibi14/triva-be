# Triva - Quiz App

### Tech Stack
1. [Svelte](https://svelte.dev/)
2. [TypeScript](https://www.typescriptlang.org/)
3. [TailwindCSS](https://tailwindcss.com/)
4. [Golang](https://go.dev)
5. [Fiber](https://gofiber.io/)
6. [PostgreSQL](https://www.postgresql.org/)
7. [Redis](https://redis.io/)

### Start docker service
```bash
sudo systemctl start docker
docker-compose up -d
```

### Make database migration
```bash
make migrate state=migration_state
make migrate-up
make migrate-down
make migrate-fix version=migration_version
```

### Enter to PostgreSQL
```bash
docker exec -it triva-db psql -U habi -d triva -W
```

### Enter to Redis
```bash
docker exec -it triva-redis redis-cli -h localhost -p 6379
```

### Start development
```bash
# install intial tool
make setup

# install libraries or dependencies
go mod download

# run go server with air hot reload
air

# run svelte
cd client
pnpm dev
```

```bash
├─ _docker-data       # Docker container data
├─ bin                # Binary compiled
├─ client             # Frontend (Svelte)
├─ cmd                # Apps
├─ configs            # Configs for service/dependency
├─ databases          # Database migration stuff
├─ docs               # Config generated swagger API Docs
├─ helper             # Other codes, can be imported anywhere
├─ internal           # Most logical, including app wrapper
│   ├─ bootstrap      # App components
│   ├─ controller     # Business logic, http handler
│   ├─ repository     # Repository layer, database integration
│   ├─ service        # Service layer
│   └─ web            # Web-Server stuff
├─ logs               # Log files
├─ script             # Automation scripts, including CI/CD
├─ test               # Unit test, integration test
├─ tmp                # Temporary files, for development

```