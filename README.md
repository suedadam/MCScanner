MCScanner
=============

This is a scanner for Minecraft servers, theoretically possible to scan for large server's backends (that use frontends for protection). Use at your own risk and do not use it for malicious purposes. 

Compilation
-------------

Compiling it easy!

Grab the dependencies:
```bash
$ go get github.com/geNAZt/minecraft-status/data
$ go get github.com/geNAZt/minecraft-status/protocol
```

### Lastly ###

```bash
$ go build
$ ./main
```