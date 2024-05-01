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