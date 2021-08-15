# bbrf-amass

Parses amass json output and forward it to bbrf: domain:ip with its sources. This also supports ipv6.
Normally, I wasn't able to add domain:ip with amass sources from stdin. This solves that problem.

### Installation
```
GO111MODULE=on go get github.com/k-sau/bbrf-amass
```

### Usage

```
  -bc string
    	File path for bbrf config file. Default: ~/.bbrf/config.json (default "~/.bbrf/config.json")
  -h	Prints available flags
  -p string
    	Program id. Required.
  -path string
    	Full path to amass json output. Required.

```