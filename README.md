# goura

![](https://github.com/paveg/goura/workflows/.github/workflows/test.yaml/badge.svg)

goura is an API client of oura cloud and functionally command line tool. 

## Overview

[oura cloud api document](https://cloud.ouraring.com/docs/)

goura is an API client and command line tool.

You can easily call the API from the command line to get the json string.

## Installation

- preparing Go runtime

```bash
$ goenv install -s
```

- main installation process

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

```bash
$ goura sleeps --target 2019-03-01 | jq .
2019/09/07 10:14:01 HTTP Request: 200 OK
{
  "sleep": [
    {
      "summary_date": "2019-03-01",
      "period_id": 0,
      "is_longest": 1,
      "time_zone": 0,
      "bedtime_start": "2019-03-02T00:57:59+09:00",
      "bedtime_end": "2019-03-02T07:46:59+09:00",
      "score": 53,
      "score_total": 43,
      "score_disturbances": 55,
      "score_efficiency": 53,
      "score_latency": 81,
      "score_rem": 43,
      "score_deep": 49,
      "score_alignment": 68,
      "total": 17670,
      "duration": 24540,
      "awake": 6870,
      "light": 11550,
      "rem": 3180,
      "deep": 2940,
      "onset_latency": 180,
      "restless": 42,
      "efficiency": 72,
      "midpoint_time": 11160,
      "hr_lowest": 51,
      "hr_average": 60.625,
      "rmssd": 50,
      "breath_average": 14.75,
      "temperature_delta": 0.05,
      "hypnogram_5min": "4222444444222222122111221123344333322222222212221234222222144332334222244244444444"
    }
  ]
}
```

This example using the jq which is a lightweight and flexible command-line JSON processor.

jq is [here](https://stedolan.github.io/jq/) .
