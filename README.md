# LSH Recursive Hasher
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/lshsum/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/lshsum/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/lshsum/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/lshsum?status.png)](http://godoc.org/github.com/pedroalbanese/lshsum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/lshsum)](https://goreportcard.com/report/github.com/pedroalbanese/lshsum)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/lshsum)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/lshsum)](https://github.com/pedroalbanese/lshsum/releases)
### TTAK.KO-12.0276 LSH Recursive Hasher written in Go:
<PRE>
Usage of lshsum:
lshsum [-c &lt;hash.ext&gt;] [-b N] [-r] &lt;file.ext&gt;
  -b int
        Bits: 224, 256, 384 and 512. (default 256)
  -c string
        Check hashsum file.
  -r    Process directories recursively.
</PRE>

### LSH-256 Hash-based Message Autentication Code Utility:
<PRE>
Usage of lshmac:
lshmac [-k &lt;secret&gt;] -f &lt;file.ext&gt;
  -f string
        Target file. ('-' for STDIN)
  -k string
        Secret key.
</PRE>
## Examples
### Generate hashsum list:
```sh
$ ./lshsum [-r] "*.*" > hash.txt
```

### Check hashsum file:
```sh
$ ./lshsum -c hash.txt
$ echo $?
```

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2022 Pedro F. Albanese - ALBANESE Research Lab.
