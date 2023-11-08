# Docker Install easy-admin
> Build a full-featured administrative interface quickly easy-admin

### 1、Ready

> tip: [how to install docker](https://nicelizhi.github.io/easy-admin/guide/install/installdocker)

```
docker pull nicesteven/easy-admin 
```
> view docker images
```
docker images
```

### 2、Configure

[Configure Docs](https://nicelizhi.github.io/easy-admin/guide/configure/)

### 3、Start It

> You need have setting.yml file in ./config dir
```
docker run --name easy-admin -p 8000:8000 -v ./config:/config/ -d nicesteven/easy-admin
```

### 4、Test it

```
docker exec -it easy-admin bash 
netstat -an | grep 8000
```

### 5、FAQ


### Issue Submit
https://github.com/nicelizhi/easy-admin/issues

