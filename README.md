#### running environment development
```
docker-compose up -d
make runserver
```
#### running load test
```
cd scripts && wrk -c 80 -d 10s -s ./login.lua http://localhost:8080/api/v1/login
```
