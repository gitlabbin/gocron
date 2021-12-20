# gocron - Schedule Job Manage System
[![Downloads](https://img.shields.io/github/downloads/ouqiang/gocron/total.svg)](https://github.com/ouqiang/gocron/releases)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/ouqiang/gocron/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/ouqiang/gocron.svg?label=Release)](https://github.com/ouqiang/gocron/releases)

*Fork of [ouqiang/gocron](https://github.com/ouqiang/gocron)*

# Description
Using Golang to implement a light schedule job running and managing system, it can replace Linux-crontab [Check doc](https://github.com/ouqiang/gocron/wiki)

Original `DelayJob` was put it into another project [DelayQueue](https://github.com/ouqiang/delay-queue)  

## Features
* Web console to manage schedule job
* Crontab expression, acurate to second level
* Job failed can use retry
* Job timeout, will force exit
* Job dependency, Job A --> Job B
* User access permission
* Job type
    * shell job
    > In node to run shell command, support to run in multiple nodes same time
    * HTTP job
    > use http URL, will run by gocron server, no need node
* Job running log
* Notification: email、Slack、Webhook

### Diagram
```shell

     Scheduler ----|--> Node 1
                   |--> Node 2
                   |--> Node 3

```

### OS system
> Linux, Mac OS, Windows

### Database
-  Sqlite3
-  MySQL
-  Postgres

## Download
[releases](https://github.com/gitlabbin/gocron/releases)  

## Install

###  Binary Install
1. tar -xvzf [xxx.gz] 
2. `cd [unziped folder]`   
3. Run apps        
* Run Scheduler        
  * Windows: `gocron.exe web`   
  * Linux、Mac OS:  `./gocron web`
* Run Node, default on: 0.0.0.0:5921
  * Windows:  `gocron-node.exe`
  * Linux、Mac OS:  `./gocron-node`
4. Browse the url: http://localhost:5920

### Source code Install 

- Install Golang 1.11+
- `go get -d github.com/ouqiang/gocron`
- `export GO111MODULE=on` 
- compile `make`
- Run
    * gocron `./bin/gocron web`
    * gocron-node `./bin/gocron-node`
  
### docker

```shell
docker run --name gocron --link mysql:db -p 5920:5920 -d ouqg/gocron
```

Setting: /app/conf/app.ini

Log: /app/log/cron.log


### Develop

1. Install Golang 1.9+, Node.js, Yarn
2. Front end dependency `make install-vue`
3. Run gocron, gocron-node `make run`
4. Run node server `make run-vue`, browse http://localhost:8080

request to http://localhost:8080, API request will forward to `gocron`

`make` compile

`make run` compile and run

`make package` package
> produce the gzip package for app: gocron-v1.5-darwin-amd64.tar.gz gocron-node-v1.5-darwin-amd64.tar.gz

`make package-all` produce Windows、Linux、Mac gzip package

### CMD

* gocron
    * -v version

* gocron web
    * --host default 0.0.0.0
    * -p port, default 5920
    * -e environment, dev|test|prod, dev mode having more logs, default `prod`
    * -h help
* gocron-node
    * -allow-root *nix allow to run as `root`
    * -s ip:port listen on address  
    * -enable-tls enable TLS    
    * -ca-file   CA files   
    * -cert-file cert file  
    * -key-file  key file
    * -h help
    * -v version

## To Do List
- [x] version upgrade
- [x] batch start, clos, delete jobs
- [x] Scheduler communicate with Node support https
- [x] Job groups
- [x] Multiple user
- [x] User Access permissions

## Libs and Framework
* Web [Macaron](http://go-macaron.com/)
* Scheduler [Cron](https://github.com/robfig/cron)
* ORM [Xorm](https://github.com/go-xorm/xorm)
* UI [Element UI](https://github.com/ElemeFE/element)
* RPC [gRPC](https://github.com/grpc/grpc)

## Feedback
raise issue

## ChangeLog

v1.6.0
- Support `Sqlite3`
- Better handle long run shell job
    * shell cmd output result to log
    * keep max 2000 lines in buffer to minimize memory usage
- Fix `shell` mode timeout goroutine leaking
- UI support international (en, zh)
- Backend support international (en, zh)
- GoCron default lang setting in app.ini (app.lang = en)

v1.6.1
- Fix Panic possibly due to x/sys version
  Go 1.17 segmentation violation on MacOS Big Sur
  logrus old version https://github.com/sirupsen/logrus/issues/1285