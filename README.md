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
# Reset your ownCloud instance while removing all your data.
$ make reset-oc-data
```

## Services

List of available services as docker containers. Just hit `docker-compose up`.

* Collabora
* OnlyOffice
* MailHog
* Selenium
* LDAP
* Redis
* SFTP-Server
