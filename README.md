# deployer
[![Build Status](https://travis-ci.org/valentin2105/deployer.svg?branch=master)](https://travis-ci.org/valentin2105/deployer)

![](https://i.imgur.com/Je8FbDT.png)

## Description
a Go docker-compose multi-stages deployer.

Use case example[ here](https://opsnotice.xyz/deployer-multi-stage-dockercompose/
).

## Build

To build it, use `go get` and `gopm` for dependencies :

```bash
$ go get -d github.com/valentin2105/deployer
$ go get -u github.com/gpmgo/gopm
$ cd $GOPATH/src/github.com/valentin2105/deployer && $GOPATH/bin/gopm get 
$ go build && ./deployer -h
```

## Example

#### First, you need a config.json file :
```
{
   “config”:
 {
     “WpImage”: “wordpress:latest”,
     “DBImage”: “mysql:latest”,
     “NginxImage”: “nginx:latest”
 },
   “dev”:
 {
     “Tag”: “dev”,
     “Vhost”: “dev.example.com”,
     "DBPassword": "AnyGoodPassword",
     "DBName": "mydevsite",
     "ExpositionPort": "8001:80"
 },
   “prod”:
 {
     “Tag”: “integration”,
     “Vhost”: “integration.example.com”,
     "DBPassword": "AnyBetterPassword",
     "DBName": "myprodsite",
     "IPv6Network": "ff00:c210::/64",
     "IPv6": "ff00:c210::121"
 }
}
```
#### Then, you can create your compose/dev.tmpl.yml file :
```
version: '2'
services:
  wordpress:
    image: {{.config_WpImage}}
    ports:
      - {{.dev_ExpositionPort}}
    volumes:
      - /var/www/html 
    environment:
      WORDPRESS_DB_HOST: db
      WORDPRESS_DB_NAME: {{.dev_DBName}}
      WORDPRESS_DB_USER: root
      WORDPRESS_DB_PASSWORD: {{.dev_DBPassword}}
    depends_on:
      - db
    links:
      - db

  db:
    image: {{.config_DBImage}}
    volumes:
      - /var/lib/mysql
    environment:
      MYSQL_DATABASE: {{.dev_DBName}}
      MYSQL_ROOT_PASSWORD: {{.dev_DBPassword}}
```
#### Finally, you can deploy your dev environement :
```
deployer add dev
```

## Usage

![](http://i.imgur.com/ngkdqr0.gif)

```bash
NAME:
   deployer

USAGE:
   deployer [global options] command [command options] [arguments...]

VERSION:
   0.1.5

AUTHOR:
   Valentin Ouvrard

COMMANDS:
     deploy
     delete
     list
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Contribution

1. Fork ([https://github.com/valentin2105/deployer/fork](https://github.com/valentin2105/deployer/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Valentin Ouvrard](https://github.com/valentin2105)
