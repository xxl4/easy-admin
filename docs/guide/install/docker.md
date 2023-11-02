# Docker Install easy-admin
> Build a full-featured administrative interface quickly easy-admin

### 1、Ready

> tip: [how to install docker](https://nicelizhi.github.io/easy-admin/guide/install/installdocker)

```
docker pull registry.ap-southeast-1.aliyuncs.com/kuops/easy-admin:latest 
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
docker run --name api-admin -p 8000:8000 -v ./config:/easy-admin/config/ -d registry.ap-southeast-1.aliyuncs.com/kuops/easy-admin:latest
```

### 4、Test it

```
docker exec -it api-admin bash 
netstat -an | grep 8000
```

### Issue Submit
https://github.com/nicelizhi/easy-admin/issues

