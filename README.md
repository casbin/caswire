<h1 align="center" style="border-bottom: none;">📦⚡️ Caswire</h1>
<h3 align="center">An open-source logging and auditing system developed by Go and React.</h3>
<p align="center">
  <a href="#badge">
    <img alt="semantic-release" src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
  </a>
  <a href="https://hub.docker.com/r/casbin/caswire">
    <img alt="docker pull casbin/caswire" src="https://img.shields.io/docker/pulls/casbin/caswire.svg">
  </a>
  <a href="https://github.com/casbin/caswire/releases/latest">
    <img alt="GitHub Release" src="https://img.shields.io/github/v/release/casbin/caswire.svg">
  </a>
  <a href="https://hub.docker.com/repository/docker/casbin/caswire">
    <img alt="Docker Image Version (latest semver)" src="https://img.shields.io/badge/Docker%20Hub-latest-brightgreen">
  </a>
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/casbin/caswire">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/casbin/caswire?style=flat-square">
  </a>
  <a href="https://github.com/casbin/caswire/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/casbin/caswire?style=flat-square" alt="license">
  </a>
  <a href="https://github.com/casbin/caswire/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/casbin/caswire?style=flat-square">
  </a>
  <a href="#">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/casbin/caswire?style=flat-square">
  </a>
  <a href="https://github.com/casbin/caswire/network">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/casbin/caswire?style=flat-square">
  </a>
</p>

## Architecture
Caswire contains 2 parts:
Name | Description | Language | Source code
----|------|----|----
Frontend | Web frontend UI for Caswire | Javascript + React | https://github.com/casbin/caswire/tree/master/web
Backend | RESTful API backend for Caswire | Golang + Beego + MySQL | https://github.com/casbin/caswire

## Installation
Caswire uses Casdoor to manage members. So you need to create an organization and an application for Caswire in a Casdoor instance.

### Necessary configuration

#### Get the code
```bash
go get github.com/casbin/casdoor
go get github.com/casbin/caswire
```

or

```bash
git clone https://github.com/casbin/casdoor
git clone https://github.com/casbin/caswire
```

#### Setup database

Caswire will store its users, nodes and topics informations in a MySQL database named: `caswire`, will create it if not existed. The DB connection string can be specified at: https://github.com/casbin/caswire/blob/master/conf/app.conf

```ini
dataSourceName = root:123@tcp(localhost:3306)/
```
Caswire uses XORM to connect to DB, so all DBs supported by XORM can also be used.

#### Configure Casdoor

After creating an organization and an application for Caswire in a Casdoor, you need to update `clientID`, `clientSecret`, `casdoorOrganization` and `casdoorApplication` in app.conf.

#### Run Caswire

- Configure and run Caswire by yourself. If you want to learn more about caswire.
- Open browser: http://localhost:16001/

### Optional configuration

#### Setup your Caswire to enable some third-party login platform

  Caswire uses Casdoor to manage members. If you want to log in with oauth, you should see [casdoor oauth configuration](https://casdoor.org/docs/provider/oauth/overview).

#### OSS, Mail, and SMS services

  Caswire uses Casdoor to upload files to cloud storage, send Emails and send SMSs. See Casdoor for more details.

## Contribute

For Caswire, if you have any questions, you can give Issues, or you can also directly start Pull Requests(but we recommend giving issues first to communicate with the community).

## License

[Apache-2.0](LICENSE)
