# ipfs-webserver

A simple webserver that connects to a (local) IPFS daemon to find files. This
is very much a work-in-progress, and might not grow beyond being a toy.

This README is also a WIP.

## Intentions

I intend to run my personal blog (and possibly a few other sites) from this
webserver. For now it's a way to learn about IPFS and get to know the APIs of
it.

I intend to have my webserver(s) run this tool + an IPFS daemon, where both
can be configured separately, with this tool having as little configuration as
necessary. Possibly only a list of domains to serve.

When a request for a (configured) domain comes in, its DNSLink is used to
fetch the path from the IPFS daemon. Whether this object is pinned or even
present on the daemon should not matter, apart from performance obviously.
Best would be to ensure the resource is pinned before exposing it through the
DNSLink.

To further speed things up, I intend to add a
[ristretto](https://github.com/dgraph-io/ristretto) cache.
