# goura

![](https://github.com/paveg/goura/workflows/static%20code%20test/badge.svg)

goura is an unofficial Go client for the [Oura Ring API v2](https://cloud.ouraring.com/v2/docs) with a command-line interface.

## Overview

goura provides both a Go library and CLI tool to access your Oura Ring health data through the official Oura API v2.

### Supported Data Types

| Category | Endpoints |
|----------|-----------|
| **User** | Personal Info, Ring Configuration |
| **Sleep** | Daily Sleep, Sleep Periods, Sleep Time Recommendations |
| **Activity** | Daily Activity |
| **Readiness** | Daily Readiness |
| **Heart** | Heart Rate (time-series), SpO2 (blood oxygen) |
| **Wellness** | Daily Stress, Daily Resilience, Cardiovascular Age, VO2 Max |
| **Activities** | Workouts, Sessions (guided/unguided) |
| **Tags** | Tags, Enhanced Tags, Rest Mode Periods |

## Requirements

- Go 1.21 or later
- Oura Ring with active membership (Gen3/Ring 4 require membership for API access)

## Installation

### Using Go Install

```bash
go install github.com/paveg/goura@latest
```

### From Source

```bash
git clone https://github.com/paveg/goura
cd goura
make install
```

## Configuration

### 1. Create an Oura Application

Go to [cloud.ouraring.com/oauth/applications](https://cloud.ouraring.com/oauth/applications) and create a new application.

- Set **Redirect URL** to `http://localhost:8989`

### 2. Set Environment Variables

```bash
export OURA_CLIENT_ID=your_client_id
export OURA_CLIENT_SECRET=your_client_secret
```

Alternatively, you can use a Personal Access Token directly:

```bash
export OURA_ACCESS_TOKEN=your_personal_access_token
```

> **Note**: Personal Access Tokens will be deprecated by end of 2025. OAuth2 is recommended.

### 3. Authenticate

```bash
goura configure
```

This opens your browser for OAuth2 authorization and stores the token in `~/.goura.yaml`.

## Usage

### Available Commands

```
goura [command]

Commands:
  userinfo      Fetch user personal information
  ring          Fetch ring configuration
  sleep         Fetch daily sleep scores
  sleep-periods Fetch detailed sleep period data
  sleep-time    Fetch sleep time recommendations
  activity      Fetch daily activity data
  readiness     Fetch daily readiness scores
  heartrate     Fetch heart rate data
  spo2          Fetch daily SpO2 (blood oxygen) data
  stress        Fetch daily stress data
  resilience    Fetch daily resilience data
  workout       Fetch workout data
  session       Fetch session data
  configure     Authenticate with Oura API
  version       Print version information
```

### Common Flags

Most data commands support these flags:

```
-s, --start string   Start date (YYYY-MM-DD), default: 1 month ago
-e, --end string     End date (YYYY-MM-DD), default: today
-t, --target string  Fetch data for a specific day (YYYY-MM-DD)
```

The `heartrate` command uses datetime flags:

```
--start-datetime string  Start datetime (ISO 8601)
--end-datetime string    End datetime (ISO 8601)
```

## Examples

### Get User Information

```bash
$ goura userinfo
{
  "id": "abc123",
  "age": 30,
  "weight": 70.5,
  "height": 175,
  "biological_sex": "male",
  "email": "user@example.com"
}
```

### Get Daily Sleep Data

```bash
$ goura sleep -t 2024-01-15
{
  "data": [
    {
      "id": "sleep-123",
      "day": "2024-01-15",
      "score": 85,
      "timestamp": "2024-01-15T07:30:00+00:00",
      "contributors": {
        "deep_sleep": 80,
        "efficiency": 90,
        "latency": 85,
        "rem_sleep": 75,
        "restfulness": 88,
        "timing": 92,
        "total_sleep": 82
      }
    }
  ],
  "next_token": null
}
```

### Get Heart Rate Data

```bash
$ goura heartrate --start-datetime 2024-01-15T00:00:00Z --end-datetime 2024-01-15T23:59:59Z
{
  "data": [
    {
      "bpm": 65,
      "source": "rest",
      "timestamp": "2024-01-15T03:00:00+00:00"
    },
    {
      "bpm": 72,
      "source": "awake",
      "timestamp": "2024-01-15T10:30:00+00:00"
    }
  ],
  "next_token": null
}
```

### Get Workout Data

```bash
$ goura workout -s 2024-01-01 -e 2024-01-31
{
  "data": [
    {
      "id": "workout-123",
      "activity": "running",
      "calories": 350.5,
      "day": "2024-01-15",
      "distance": 5000,
      "intensity": "moderate",
      "source": "manual",
      "start_datetime": "2024-01-15T07:00:00+00:00",
      "end_datetime": "2024-01-15T07:45:00+00:00"
    }
  ],
  "next_token": null
}
```

### Using with jq

For formatted output, pipe to [jq](https://stedolan.github.io/jq/):

```bash
$ goura sleep -t 2024-01-15 | jq '.data[0].score'
85
```

## Library Usage

```go
package main

import (
    "context"
    "fmt"
    "net/http"

    "github.com/paveg/goura/api"
    "github.com/paveg/goura/oura"
)

func main() {
    client, _ := api.NewClient(
        "https://api.ouraring.com",
        &http.Client{},
        "myapp/1.0",
        "your_access_token",
    )

    ctx := context.Background()

    // Get personal info
    info, _ := client.GetPersonalInfo(ctx)
    fmt.Printf("Email: %s\n", info.Email)

    // Get daily sleep data
    period := oura.DatePeriod{
        StartDate: "2024-01-01",
        EndDate:   "2024-01-31",
    }
    sleep, _ := client.GetDailySleep(ctx, period)
    for _, s := range sleep.Data {
        fmt.Printf("Day: %s, Score: %d\n", s.Day, *s.Score)
    }

    // Get heart rate data
    hrPeriod := oura.DateTimePeriod{
        StartDateTime: "2024-01-15T00:00:00Z",
        EndDateTime:   "2024-01-15T23:59:59Z",
    }
    hr, _ := client.GetHeartRate(ctx, hrPeriod)
    for _, h := range hr.Data {
        fmt.Printf("BPM: %d at %s\n", h.BPM, h.Timestamp)
    }
}
```

## API Reference

See the [Oura API v2 Documentation](https://cloud.ouraring.com/v2/docs) for detailed information about available data and response formats.

## License

MIT
