# ovh-api-creds

A tiny CLI to manage OVH API credentials


## Installation

Download the binary [directly](https://github.com/xlucas/ovh-api-creds/releases).

Alternatively, for gophers:

```bash
go get github.com/xlucas/ovh-api-creds
```


## Usage

An application communicating with OVH APIs requires the following credentials:

* an application key
* a secret token
* a consumer key


### Retrieve all credentials at once (all-in-one)

With this approach you are given all credentials in a one-step process.

```bash
ovh-api-creds application create --service kimsufi --zone ca --all-in-one
```

### Retrieve an application key and an application secret

With this approach you aren't given a consumer key. It has to be generated
with the CLI in a second step.

```bash
ovh-api-creds application create --service ovh --zone eu
```

### Retrieve a consumer key

Given you have an application key and a secret token you can request a consumer
key.

```bash
ovh-api-creds consumer-key create --service ovh --zone eu --app-key *** --app-secret ****
```

You can also give it read-only or read-write access and/or restrict API
access to a particular path.

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
