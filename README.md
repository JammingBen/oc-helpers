# ownCloud Helpers

A small helper project for ownCloud including handy commands and services to make developing easier.

## Prerequisites

* docker
* docker-compose
* Go
* build-essentials

Copy the .env.dist file (`cp .env.dist .env`) and adjust its values.

## Commands

```
# Use the command via go run
$ go run cmd/*.go

# ...or build an executable in build/oc-helper
$ make build
```

## Services

List of available services as docker containers. Just hit `docker-compose up`.

* Collabora
* OnlyOffice (use docker IP or switch to host mode in docker-compose.yml)
* MailHog
* Selenium
* LDAP
* Redis
* SFTP-Server
* FTP-Server
* Elasticsearch
* Samba
