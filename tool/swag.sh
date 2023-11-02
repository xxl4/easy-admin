#!/bin/bash

# how to install swag https://github.com/swaggo/swag

swag i -g init_router.go -dir app/admin/router --instanceName admin --parseDependency -o docs/admin
