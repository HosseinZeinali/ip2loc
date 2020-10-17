# ip2Loc
# Installation
1-Make sure you have installed docker/docker-compose and Make.

2-clone this repository
```
git clone git@github.com:HosseinZeinali/ip2loc.git
```
3-make .env file based on env.dist

4-Then run below command:
```sh
make up
```
That's it.
# Endpoints
```
GET http://localhost:8082/api/v1/ip/${ip}
```
