
Microservices to check the skills of the developer.
- User-api
- Order-api

## Project Requirements

- Golang
- Mysql 
- Redis
- ElasticSearch 

## Technologies

- Golang 1.14 (https://golang.org/)
- mysql 8 (https://www.mysql.com/)
- Elasticsearch (https://www.elastic.co/)
- Redis (https://redis.io/)
- GORM (https://gorm.io/)

## Instalation
For the installation to work properly, it's necessary to have __8000, 8001,9200 and 3306 opened__. <br><br>
The installation consists in these steps:
- Install Golang (https://golang.org/doc/install)
- mysql 8 (https://dev.mysql.com/doc/refman/8.0/en/installing.html)
- Elasticsearch (https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html)
- Redis (https://redis.io/topics/quickstart)
- GORM (It's in Go modules)
- Create 2 databases(following the config files), and change the service/config/config.example.json to config.json 

## Routes Documentation
The routes documentation were made in Postman, you can import this collection to your postman, or recreate it using the 
examples.<br>
Order api:
- https://documenter.getpostman.com/view/3611232/T17Ge7ua?version=latest#aea84270-03ff-4bad-8555-43675fbd22d3
User api:
- https://documenter.getpostman.com/view/3611232/T17GeShe?version=latest#3b260316-af3a-4d24-bd0b-62132c1a47b4
## Databases
- https://imgur.com/ROxT3Vx
- https://imgur.com/61aWzVC

## Tests
To execute the service tests<br>
Execute:<br>
`make test`
If you don't have `make` installed just run from the directory user-api/order-api:
`sh ./scripts/tests.sh`

## Checklist
- [x] Define MS structure.
- [x] Define whether or not i will use ORM, and if so, which one. (GORM)
- [x] Create database.
- [x] Accounts routes working happy way.
- [x] Accounts unit tests on service.
- [x] Services routes working happy way.
- [x] Services unit tests on service.
- [x] Errors handling (correct http responses).
- [x] Log control.
- [X] Encrypt Sensitive Data.
- [X] Cache.
- [X] Elasticsearch.
- [x] Finish documentation.

## Why we have dockerfiles here and we don't use?
I started to mount everything on docker, but there wasn't time enought, if we go to the next step of the interview i'll explain everything, if you guys need anything, you're able to contact me on: 14 99858-3391

## Elasticsearch
The same thing that happened with docker, happened here, i was able only to implement elastic on the Orders insert

# Thank you !
