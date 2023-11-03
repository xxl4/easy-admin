# Docker Compose Install easy-admin
> Build a full-featured administrative interface quickly easy-admin

### 1、Ready

> 1.1、Docker-compose install

```
https://docs.docker.com/compose/install/
```

> 1.2、docker-compose.yaml (MySQL Version)
```
version: '3.8'
services:
  easy-admin:
    container_name: easy-admin
    image: registry.ap-southeast-1.aliyuncs.com/kuops/easy-admin:1.1.0
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
  mysql:
    container_name: mysql
    image: mysql:5.7.26
    hostname: mysql
    restart: always
    command: mysqld
    environment:
      - TZ=Asia/Shanghai
      - MYSQL_ROOT_PASSWORD=GoAdmin
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - ./data:/var/lib/mysql/data
      - ./data/my.cnf:/etc/mysql/my.cnf
    networks:
      - kuops
networks:
  kuops:
    driver: bridge
```

> 1.3、docker-compose.yaml (PG Version) (Option)
```
version: '3.8'
services:
  easy-admin:
    container_name: easy-admin
    image: registry.ap-southeast-1.aliyuncs.com/kuops/easy-admin:1.1.0
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

