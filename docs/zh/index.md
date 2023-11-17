# easy-admin


[![Release](https://img.shields.io/github/release/nicelizhi/easy-admin.svg?style=flat-square)](https://github.com/nicelizhi/easy-admin/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/nicelizhi/easy-admin)

[English](https://github.com/nicelizhi/easy-admin/blob/master/README.md) | 简体中文

基于Gin + Vue + Element UI OR Arco Design OR Ant Design的前后端分离权限管理系统,系统初始化极度简单，只需要配置文件中，修改数据库连接，系统支持多指令操作，迁移指令可以让初始化数据库信息变得更简单，服务指令可以很简单的启动api服务

## 在线DEMO
- [Vue2 demo](https://easy-admin-ui.vercel.app)  
- [Arco demo](https://hello-arco-pro.vercel.app/)( testing now)  


## ✨ 特性

- 遵循 RESTful API 设计规范

- 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 支持 Swagger 文档(基于swaggo)

- 基于 GORM 的数据库存储，可扩展多种类型数据库

- 配置文件简单的模型映射，快速能够得到想要的配置

- 代码生成工具

- 表单构建工具

- 多指令模式

- 多租户的支持

- 多语言支持 （已支持中文，英文）

- 时区功能支持

- 静态打包，从而简化部署实施流程

- Gzip 的压缩支持，从而可以让您的应用跑的更快


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
1. 接口文档：根据业务代码自动生成相关的api接口文档。
1. 代码生成：根据数据表结构生成对应的增删改查相对应业务，全程可视化操作，让基本业务可以零代码实现。
1. 表单构建：自定义页面样式，拖拉拽实现页面布局。
1. 服务监控：查看一些服务器的基本信息。
1. 内容管理：demo功能，下设分类管理、内容管理。可以参考使用方便快速入门。
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


# How to Configure

[Configure Docs](https://nicelizhi.github.io/easy-admin/guide/configure/)

# Document
[https://nicelizhi.github.io/easy-admin/](https://nicelizhi.github.io/easy-admin/) 

# Issue
[https://github.com/nicelizhi/easy-admin/issues](https://github.com/nicelizhi/easy-admin/issues)   
[https://gitee.com/nicelizhi/easy-admin/issues](https://gitee.com/nicelizhi/easy-admin/issues)  (中国)

# Discussions
[https://github.com/nicelizhi/easy-admin/discussions](https://github.com/nicelizhi/easy-admin/discussions) 