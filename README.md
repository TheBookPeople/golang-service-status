# golang-service-status

Provides a JSON object exposing status information, ideal as part of a REST API for internal monitoring.

## Installation

Import with something akin to:

```go
import servicestatus "github.com/TheBookPeople/golang-service-status"
```

## Usage

```go
// Initialize
status := servicestatus.NewServiceStatus("app name", "0.0.1")

// Add service checks
status.AddCheck("Redis Ping", func() bool {
  // Perform the check here, return true if passed 
  return true
})

// Get a JSON string containing service info:
status.Status()
```
## Format

As below. Checks are always part of the 'checks' array. If they have failed,
they are included in 'errors' as well.

'Status' is either "online" or "offline".

```javascript
{
  "name": "windows",
  "version": "3.11",
  "hostname": "clippy",
  "checks": [{
    "name": "redis",
    "description": "Can connect to Redis"
    "successful": true
  }],
  "errors": [],
  "timestamp": "2015-05-07 14:35:17",
  "uptime": "14d:23:11:21",
  "diskusage": "64%",
  "status": "online"
}
```

## Contributing

1. Fork it ( https://github.com/[my-github-username]/service-status/fork )
2. git flow init -d (assuming you have git flow)
2. Create your feature branch (`git flow feature start my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

## Copyright

Copyright (c) 2015 The Book People

See LICENSE (GPLv3)
