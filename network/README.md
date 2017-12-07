# Start Up HFCS Network

To start up a 3-orgs HFCS network:

```bash
$ curl -sSL https://goo.gl/5ftp2f | bash
$ export PATH=<path to current location>/bin:$PATH
$ cd first-network
$ ./byfn.sh -m generate
$ ./byfn.sh -m up
```

For IBM LinuxONE:

```bash
$ rm -r bin/
$ curl -sSL https://goo.gl/5ftp2f | bash
$ export PATH=<path to current location>/bin:$PATH
$ cd first-network
$ ./byfn.sh -m generate
$ ./byfn.sh -m up
```


