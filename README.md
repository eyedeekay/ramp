# ramp

*[allium tricoccum](https://en.wikipedia.org/wiki/Allium_tricoccum)*

*[inclined plane](https://en.wikipedia.org/wiki/Inclined_plane)*

Convenient structures for managing sam sessions in Go. Some of them might be
reimplemented in sam3, it's initially intended to make it more intuitive to port
[cretz/bine](https://github.com/cretz/bine) based applications to I2P streaming.

Besides that, it in general exists to explore a few ease-of-use issues in sam3
and goSam, like how for now they use different, non-compatible formats for
saving keys, which are in turn not compatible with i2ptunnel, the
more-tedious-than-it-needs-to-be keeping track of configuration options, full
support for Contexts, etc.

        ./        : Base Ramp classes, the things most will need to use
        ./config  : Configuration classes to deal with compatibility issues
        ./emit    : Takes a config file and uses it to send SAM commands
