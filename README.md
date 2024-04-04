### Kahoot clone

**Tech Stack :**
1. Svelte
2. TypeScript
3. TailwindCSS
4. Golang
5. Fiber
6. PostgreSQL

```bash
sudo systemctl start docker
docker-compose up -d

docker exec -it bwizz-db psql -U habi bwizz -W

make migrate state=migration_state
make migrate-up
make migrate-down
```