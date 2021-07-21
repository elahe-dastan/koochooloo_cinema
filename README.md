# koochooloo_cinema

# Database

I use a postgres database in my docker-compose. use below commands to get into the container

```shell
docker exec -it <container_name> /bin/bash
psql --host=database --username=<username> --dbname=<dbname>
```

```sh
curl -X POST http://127.0.0.1:1373/api/signup -H 'Content-Type: application/json' -d '{
  "username": "1995parham",
  "password": "Parham123123",
  "first_name": "Parham",
  "last_name": "Alvani",
  "email": "parham.alvani@gmail.com",
  "phone": "09390909540",
  "national_number": "0017784646"
}'
```

```sh
curl -X POST http://127.0.0.1:1373/api/admin -H 'Content-Type: application/json' -d '{
  "file": "file",
  "name": "The Movie",
  "producers": [ "awesome" ],
  "production_year": 2021,
  "explanation": "awesome",
  "view": 100,
  "price": 100,
  "tags": [ "t1" ]
}'
```

```sh
curl -X POST http://127.0.0.1:1373/api/favorite -H 'Content-Type: application/json' -d '{
  "username": "1995parham",
  "film": [ 1 ],
  "album": "the album"
}'
```

```sh
curl -X POST http://127.0.0.1:1373/api/special/wallet -H 'Content-Type: application/json' -d '{
  "username": "1995parham"
}'
```
