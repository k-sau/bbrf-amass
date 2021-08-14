# bbrf-amass

Parses amass json output and forward it to bbrf: domain:ip & sources

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