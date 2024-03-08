
# Easy-Admin
[![Go Reference](https://godoc.org/github.com/xxl4/easy-admin?status.svg)](https://godoc.org/github.com/xxl4/easy-admin)
[![Build Status](https://github.com/xxl4/easy-admin/workflows/Build/badge.svg)](https://github.com/xxl4/easy-admin)
[![Release](https://img.shields.io/github/release/xxl4/easy-admin.svg?style=flat-square)](https://github.com/xxl4/easy-admin/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/xxl4/easy-admin)](https://goreportcard.com/report/github.com/xxl4/easy-admin)
[![HitCount](https://views.whatilearened.today/views/github/xxl4/easy-admin.svg)](https://github.com/xxl4/easy-admin)
[![License](https://img.shields.io/github/license/xxl4/easy-admin
)](https://github.com/xxl4/easy-admin)
[![Commits](https://img.shields.io/github/commit-activity/m/xxl4/easy-admin?color=ffff00)](https://github.com/xxl4/easy-admin/commits/main)
[![Docker Pulls](https://img.shields.io/docker/pulls/nicesteven/easy-admin)](https://hub.docker.com/r/nicesteven/easy-admin)


## About Easy-Admin

Easyadmin is a Quick, beautiful and modern admin generator for Go VUE application.

English | [简体中文](https://xxl4.github.io/easy-admin/zh/)


>> The front-end and back-end separation authority management system based on Gin + Vue  is extremely simple to initialize the system. You only need to modify the database connection in the configuration file. The system supports multi-instruction operations. Migration instructions can make it easier to initialize database information. Service instructions It's easy to start the api service.

## Online demo
- [Vue2 demo](https://easy-admin-ui.vercel.app)  
- [Arco demo](https://hello-arco-pro.vercel.app/)( testing now)  

##  Feature

- Based on the GIN WEB API framework, it provides rich middleware support (user authentication, cross-domain, access log, tracking ID, Cahce, Zip etc.)

- RBAC access control model based on Casbin

- JWT authentication

- Support Swagger documents (based on swaggo)

- Database storage based on GORM, which can expand multiple types of databases

- Simple model mapping of configuration files to quickly get the desired configuration

- Code generation tool

- Multi-command mode

- Multi Language

- Multi-platform (Darwin Freebsd Linux Windows)

- TimeZone Support

- Gzip Support, make your application faster.


## Internal

1. User management: The user is the system operator, this function mainly completes the system user configuration.
2. Department management: configure the system organization (company, department, group), and display the tree structure to support data permissions.
3. Position management: configure the positions of system users.
4. Menu management: configure the system menu, operation authority, button authority identification, interface authority, etc.
5. Role management: Role menu permission assignment and role setting are divided into data scope permissions by organization.
6. Dictionary management: Maintain some relatively fixed data frequently used in the system.
7. Parameter management: dynamically configure common parameters for the system.
8. Operation log: system normal operation log record and query; system abnormal information log record and query.
9. Login log: The system login log record query contains login exceptions.
1. Interface documentation: Automatically generate related api interface documents according to the business code.
1. Code generation: According to the data table structure, generate the corresponding addition, deletion, modification, and check corresponding business, and the whole process of visual operation, so that the basic business can be implemented with zero code.
1. Service monitoring: View the basic information of some servers.
1. Content management: demo function, including classification management and content management. You can refer to the easy to use quick start.


# Api Document Generate

```
swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin
```

# Online wagger Document

```
https://editor.swagger.io/?url=https://xxl4.github.io/easy-admin/admin/admin_swagger.yaml
```

# Install

- [Docker Install](https://xxl4.github.io/easy-admin/guide/install/docker)

- [Docker Composer Install](https://xxl4.github.io/easy-admin/guide/install/docker-composer)

- [K8s Install](https://xxl4.github.io/easy-admin/guide/install/k8s)

- [Binary Install](https://xxl4.github.io/easy-admin/guide/install/binary)


# Configure

[Configure Docs](https://xxl4.github.io/easy-admin/guide/configure/)

# Document
[https://xxl4.github.io/easy-admin/](https://xxl4.github.io/easy-admin/) 

# Issue
[https://github.com/xxl4/easy-admin/issues](https://github.com/xxl4/easy-admin/issues)   
[https://gitee.com/xxl4/easy-admin/issues](https://gitee.com/xxl4/easy-admin/issues)  (中国)

# Discussions
[discussions](https://github.com/xxl4/easy-admin/discussions) 
