# bbrf-amass

Parses amass json output and forward it to bbrf: domain:ip with its sources. This also supports ipv6.
Normally, I wasn't able to add domain:ip with amass sources from stdin. This solves that problem.

### Installation
```
GO111MODULE=on go install github.com/k-sau/bbrf-amass@latest
```

### Usage

```
  -bc string
        File path for bbrf config file. Default: ~/.bbrf/config.json (default "~/.bbrf/config.json")
  -h    Prints available flags
  -p string
        Program id. Required.
  -path string
        Full path to amass json output.
  -service
        Takes input from stdin in format of ip;port;service-name. Supports ipv6
  -wildcard
        Adds everything excepts domains which explicitly mentioned in out of scope.
```
