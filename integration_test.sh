#!/bin/bash

echo "create 1995parham with invalid password"
curl -X POST http://127.0.0.1:1373/api/signup -H 'Content-Type: application/json' -d '{
  "username": "1995parham",
  "password": "123123",
  "first_name": "Parham",
  "last_name": "Alvani",
  "email": "parham.alvani@gmail.com",
  "phone": "09390909540",
  "national_number": "0017784646"
}'

echo "create 1995parham"
curl -X POST http://127.0.0.1:1373/api/signup -H 'Content-Type: application/json' -d '{
  "username": "1995parham",
  "password": "Parham123123",
  "first_name": "Parham",
  "last_name": "Alvani",
  "email": "parham.alvani@gmail.com",
  "phone": "09390909540",
  "national_number": "0017784646"
}'

echo "create elda"
curl -X POST http://127.0.0.1:1373/api/signup -H 'Content-Type: application/json' -d '{
  "username": "elda",
  "password": "Elahe123123",
  "first_name": "Elahe",
  "last_name": "Dastan",
  "email": "elahe.dstn@gmail.com",
  "phone": "09336005978",
  "national_number": "0012254646"
}'

echo 'create The Movie 1'
curl -X POST http://127.0.0.1:1373/api/admin -H 'Content-Type: application/json' -d '{
  "file": "file",
  "name": "The Movie 1",
  "producers": [ "awesome" ],
  "production_year": 2021,
  "explanation": "awesome",
  "view": 1,
  "price": 100,
  "tags": [ "t1" ]
}'

echo 'create The Movie 2'
curl -X POST http://127.0.0.1:1373/api/admin -H 'Content-Type: application/json' -d '{
  "file": "file",
  "name": "The Movie 2",
  "producers": [ "awesome" ],
  "production_year": 2021,
  "explanation": "awesome",
  "view": 0,
  "price": 100,
  "tags": [ "t1" ]
}'

echo 'create an album for 1995parham without being special'
curl -X POST http://127.0.0.1:1373/api/favorite -H 'Content-Type: application/json' -d '{
  "username": "1995parham",
  "film": [ 1, 2 ],
  "album": "the album"
}'

echo 'promote 1995parham to be special without credit'
curl -X POST http://127.0.0.1:1373/api/special/wallet -H 'Content-Type: application/json' -d '{
  "username": "1995parham"
}'

echo 'increase 1995parham credit'
curl -X POST http://127.0.0.1:1373/api/wallet -H 'Content-Type: application/json' -d '{
  "username": "1995parham",
  "credit": 10000
}'

echo 'promote 1995parham to be special with credit'
curl -X POST http://127.0.0.1:1373/api/special/wallet -H 'Content-Type: application/json' -d '{
  "username": "1995parham"
}'

echo 'create an album for 1995parham'
curl -X POST http://127.0.0.1:1373/api/favorite -H 'Content-Type: application/json' -d '{
  "username": "1995parham",
  "film": [ 1, 2 ],
  "album": "the album 1"
}'

echo 'create another album for 1995parham'
curl -X POST http://127.0.0.1:1373/api/favorite -H 'Content-Type: application/json' -d '{
  "username": "1995parham",
  "film": [ 1 ],
  "album": "the album 2"
}'

echo 'watch a movie without credit'
curl http://127.0.0.1:1373/api/film/1/watch/elda

echo 'watch a movie with credit'
curl http://127.0.0.1:1373/api/film/1/watch/1995parham

echo 'get a movie with id'
curl http://127.0.0.1:1373/api/film/1

echo 'create vote with 1995parham for the movie'
curl -X POST http://127.0.0.1:1373/api/comment/1995parham/1 -H 'Content-Type: application/json' -d '{
  "score": 1,
  "comment": "Hello World"
}'

echo 'get a movie with id'
curl http://127.0.0.1:1373/api/film/1
