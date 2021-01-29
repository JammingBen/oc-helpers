# ownCloud Helpers

A small helper project for ownCloud including handy commands and services to make developing easier.

## Prerequisites

* docker
* docker-compose
* Go
* build-essentials

Copy the .env.dist file (`cp .env.dist .env`) and adjust its values.

## Usage

Either use the command via go run:
```
$ go run cmd/*.go
```

...or build an executable in build/oc-helper:
```
$ make build
```

## Commands

```
$ oc-helpers reset-oc # reset your OC installation
$ oc-helpers import-app-config # import app configs which are located in the app_config directory
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
