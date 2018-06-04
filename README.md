# ovh-api-creds

A tiny CLI to manage OVH API credentials


## Installation

Download the binary [directly](https://github.com/xlucas/ovh-api-creds/releases).

Alternatively, for gophers:

```bash
go get github.com/xlucas/ovh-api-creds
```


## Usage

### Create an application

Create an application via a two-step process. At first the request will return an application key and application secret token then you will have to request a consumer key using `consumer-key create`.

```bash
ovh-api-creds application create --service ovh --zone eu
```

Create a new application via a one-step process. All credentials (application key, application secret and consumer key) will be generated at once and you'll have to pick consumer key access rules during the process.

```bash
ovh-api-creds application create --service kimsufi --zone ca --all-in-one
```

### Create a consumer key

Create a new consumer key given the application key and secret token.

```bash
ovh-api-creds consumer-key create --service ovh --zone eu --app-key *** --app-secret ****
```

Create a new consumer key with read-only access to `/sms`:
```bash
ovh-api-creds consumer-key create --service ovh --zone eu --app-key *** --app-secret *** --path "/sms/*" --ro
```

## Services and Zones

Valid services are:
* kimsufi
* ovh
* soyoustart

Valid zones are:
* ca
* eu
* us

*Warning*: not all services are available in all zones!
