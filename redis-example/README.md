# Redis example

### Command to run redis in docker

[Redis Dockerhub Tags](https://hub.docker.com/_/redis?tab=tags&page=1&ordering=last_updated)

```bash
# Create Redis Server with 6379
docker run -d --rm -p 6379:6379 --name redis-server redis:6.2.5-alpine3.14

# Cheat Shet of Redis commands https://lzone.de/cheat-sheet/Redis
docker exec -it $(docker ps | grep redis-server | awk '{ print $1}') redis-cli

# Run golang with parameters
go run main.go -keys first,second,third
```
