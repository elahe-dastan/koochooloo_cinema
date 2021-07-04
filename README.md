# koochooloo_cinema

# Database
I use a postgres database in my docker-compose. use below commands to get into the container
```shell
docker exec -it <container_name> /bin/bash
psql --host=database --username=<username> --dbname=<dbname>
```