# bbrf-amass

Parses amass json output and forward it to bbrf: domain:ip with source
```
/home/op/.local/bin/bbrf domain add sub.example.com:127.0.0.1 -p test -s CertSpotter,DNS

```

### Installation
```
GO111MODULE=on go get github.com/k-sau/bbrf-amass
```

### Usage

```
  -bbrf string
    	Path to bbrf. Default: /home/op/.local/bin/bbrf (default "/home/op/.local/bin/bbrf")
  -h	Prints available flags
  -p string
    	Program id. Required.
  -path string
    	Full path to amass json output. Required.

```