# 二进制下载安装
> 通过此方案，简单的通过下载对应平台的可执行文件，做下简单的文件配置就可以完成应用的开启使用

[English](https://nicelizhi.github.io/easy-admin/guide/install/binary) | 简体中文

### 1、准备
> [Easy Admin 官方下载页面](https://github.com/nicelizhi/easy-admin/releases) 下载可执行文件


### 2、配置

[配置文件](https://nicelizhi.github.io/easy-admin/zh/guide/configure/)

### 3、开始运行 （Linux）

```
chmod +x ./easy-admin
./easy-admin server -c=config/settings.yml

```

### 4、测试（Linux）

```
ps aux | grep "easy-admin" // 
netstat -an | grep 8000
```

### 提交BUG与建议
https://github.com/nicelizhi/easy-admin/issues

