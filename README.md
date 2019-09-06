# goura

goura is an API client of oura cloud and functionally command line tool. 

## Overview

[oura cloud api document](https://cloud.ouraring.com/docs/)

goura is an API client and command line tool.

You can easily call the API from the command line to get the json string.

## Installation

```bash
$ git clone https://github.com/paveg/goura
$ cd path/to/goura
$ make install
```

### Configuration

First, go to [cloud.ouraring.com](https://cloud.ouraring.com/oauth/applications), create an application, and get a ClientID and ClientSecret.

And set it in the environment variable.

```bash
export OURA_CLIENT_ID=your_client_id
export OURA_CLIENT_SECRET=your_client_secret
```

RedirectURL should be http://localhost:8989 .

```bash
$ goura configure
# apply oura cloud
```

## Examples

```bash
$ goura userinfo | jq .
2019/09/07 08:17:57 HTTP Request: 200 OK
{
  "age": 27,
  "weight": 58.1,
  "height": 176,
  "gender": "male",
  "email": "example@gmail.com",
  "user_id": "ABCDEF12345"
}
```

This example using the jq which is a lightweight and flexible command-line JSON processor.

jq is [here](https://stedolan.github.io/jq/) .
