# Start Up HFCS Network

To start up a 4-orgs HFCS network:

```bash
# Download release tools first and decompress it.
$ ./image-pull.sh
$ export PATH=<path to current location>/bin:$PATH
$ cd start-network
$ ./byfn.sh -m generate
$ ./byfn.sh -m up
```



