# Easy Admin
[![Build Status](https://github.com/nicelizhi/easy-admin/workflows/Build/badge.svg)](https://github.com/nicelizhi/easy-admin)
[![Release](https://img.shields.io/github/release/nicelizhi/easy-admin.svg?style=flat-square)](https://github.com/nicelizhi/easy-admin/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nicelizhi/easy-admin)](https://goreportcard.com/report/github.com/nicelizhi/easy-admin)
[![HitCount](https://views.whatilearened.today/views/github/nicelizhi/easy-admin.svg)](https://github.com/nicelizhi/easy-admin)
[![License](https://img.shields.io/github/license/nicelizhi/easy-admin
)](https://github.com/nicelizhi/easy-admin)
[![Commits](https://img.shields.io/github/commit-activity/m/nicelizhi/easy-admin?color=ffff00)](https://github.com/nicelizhi/easy-admin/commits/main)

[English](https://nicelizhi.github.io/easy-admin/) | 简体中文

## 关于 Easy Admin
> 由于自己在具体的业务场景中，经常是会需要管理后台的来维护对应的业务系统，从而就有了这么一个想法，通过提供基础的管理后台系统的方式，方便大家可以免费使用。

> 基于Gin + Arco Design的前后端分离权限管理系统,系统初始化极度简单.实现了Docker,docker-compose,二进制下载安装，K8S 部署方式。完美的实现了静态文件打包到GO中，很好的简化用户的使用体验。

## 在线DEMO
- [Vue2 demo](https://easy-admin-ui.vercel.app)  
- [Arco demo](https://hello-arco-pro.vercel.app/)( testing now)  


## ✨ 特性

- 遵循 RESTful 接口 设计规范

- 基于 GIN 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID，缓存，压缩等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 支持 Swagger 文档(基于swaggo)

- 基于 GORM 的数据库存储，可扩展多种类型数据库, 现支持 MySQL, PostgreSQL, SQlite, SQL Server, Tidb

- 配置文件简单的模型映射，快速能够得到想要的配置

- 代码生成工具

- 表单构建工具

- 多指令模式

- 多租户的支持

- 时区功能支持

- 静态打包，从而简化部署实施流程

- Gzip 的压缩支持，从而可以让您的应用跑的更快

- 多语言支持，现支持 中文与英语，采用灵活的接口配置，方便添加更多语言的支持

- 多平台支持，多CPU架构支持。


## 🎁 内置

1. 多租户：系统默认支持多租户，按库分离，一个库一个租户。
1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3. 岗位管理：配置系统用户所属担任职务。
4. 菜单管理：配置系统菜单，操作权限，按钮权限标识，接口权限等。
5. 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6. 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7. 参数管理：对系统动态配置常用参数。
8. 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9. 登录日志：系统登录日志记录查询包含登录异常。
1. 接口文档：根据业务代码自动生成相关的接口文档。
1. 代码生成：根据数据表结构生成对应的增删改查相对应业务，全程可视化操作，让基本业务可以零代码实现。
1. 表单构建：自定义页面样式，拖拉拽实现页面布局。
1. 服务监控：查看一些服务器的基本信息。
1. 定时任务：自动化任务，目前支持接口调用和函数调用。

# Api 文档生成

```
swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin
```

# 在线文档查看

```
http://localhost:8000/swagger/admin/index.html
```

# 如何安装部署

- [Docker 安装部署](https://nicelizhi.github.io/easy-admin/guide/install/docker)

- [Docker Composer 部署安装](https://nicelizhi.github.io/easy-admin/guide/install/docker-composer)

- [K8s 部署安装](https://nicelizhi.github.io/easy-admin/guide/install/k8s)

- [二进制部署安装](https://nicelizhi.github.io/easy-admin/guide/install/binary)


# 配置

[Configure Docs](https://nicelizhi.github.io/easy-admin/guide/configure/)

# 在线文档
[https://nicelizhi.github.io/easy-admin/](https://nicelizhi.github.io/easy-admin/) 

# Issue
[https://github.com/nicelizhi/easy-admin/issues](https://github.com/nicelizhi/easy-admin/issues)   
[https://gitee.com/nicelizhi/easy-admin/issues](https://gitee.com/nicelizhi/easy-admin/issues)  (中国)

# Discussions
[https://github.com/nicelizhi/easy-admin/discussions](https://github.com/nicelizhi/easy-admin/discussions) 
