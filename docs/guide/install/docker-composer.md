# Docker Compose Install easy-admin
> Build a full-featured administrative interface quickly easy-admin

### 1、Ready

> 1.1、Docker-compose install

```
https://docs.docker.com/compose/install/
```

> 1.2、docker-compose.yaml (MySQL Version)
```
version: "3.7"

networks:
  easy-admin-network:
    ipam:
      driver: default
      config:
        - subnet: '176.7.0.0/16'
        
volumes:
  mysql:
  redis:
  
services:
  server:
    image: nicesteven/easy-admin
    restart: always
    ports:
      - '8000:8000'
    depends_on:
      - mysql
      - redis
    links:
      - mysql
      - redis
    networks:
      easy-admin-network:
        ipv4_address: 176.7.0.1
    healthcheck:
      test: ["CMD", "curl", "-f", "-X GET", "http://176.7.0.1:8080/api/v1/getinfo"]
      interval: 1m30s
      timeout: 10s
      retries: 3
      start_period: 40s
    command:
      /easy-admin server -c=/config/settings.yml

  mysql:
    image: mysql:8.0.21
    container_name: mysql
    command: mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
    restart: always
    volumes:
      - $PWD/mysql_data/:/var/lib/mysql:rw
      - $PWD/mysql_init/:/docker-entrypoint-initdb.d/:rw
      - $PWD/config/:/config/
      - /etc/localtime:/etc/localtime:ro
    environment:
      MYSQL_ALLOW_EMPTY_PASSWORD: "yes"
    networks:
      easy-admin-network:
        ipv4_address: 176.7.0.2

  redis:
    image: redis:6.0.6
    container_name: redis
    restart: always
    environment:
      ALLOW_ANONYMOUS_LOGIN: "yes"
    volumes:
      - redis:/data
    networks:
      easy-admin-network:
        ipv4_address: 176.7.0.3
```
> docker-composer.yml from https://github.com/nicelizhi/easy-admin/tree/main/deploy/docker-compose

> 1.3、docker-compose.yaml (PG Version) (Option)
```
version: '3.8'
services:
  easy-admin:
    container_name: easy-admin
    image: nicesteven/easy-admin
    privileged: true
    restart: always
    ports:
      - 8000:8000
    volumes:
      - ./config/:/easy-admin/config/
      - ./static/:/easy-admin/static/
      - ./temp/:/easy-admin/temp/
    networks:
      - kuops
  postgres:
    image: postgres:14.1-alpine
    healthcheck:
      test: [ "CMD", "pg_isready", "-q", "-d", "postgres", "-U", "root" ]
      timeout: 45s
      interval: 10s
      retries: 10
    restart: always
    environment:
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=easyadmin
      - APP_DB_USER=easyadmin
      - APP_DB_PASS=easyadmin
      - APP_DB_NAME=easyadmin
    volumes:
      - ./db:/docker-entrypoint-initdb.d/
    expod:
      - 5432:5432
networks:
  kuops:
    driver: bridge
```
### 2、Configure

[Configure Docs](https://nicelizhi.github.io/easy-admin/guide/configure/)

### 3、Start It

> You need have setting.yml file in ./config dir
```
docker-compose up -d // start
docker-compose down // down
```

### 4、Test it

```
docker exec -it easy-admin bash 
netstat -an | grep 8000
```

### Issue Submit
https://github.com/nicelizhi/easy-admin/issues

