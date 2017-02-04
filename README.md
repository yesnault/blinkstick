[![Build Status](https://travis-ci.org/yesnault/blinkstick.svg?branch=master)](https://travis-ci.org/yesnault/blinkstick)
[![GoDoc](https://godoc.org/github.com/yesnault/blinkstick?status.svg)](https://godoc.org/github.com/yesnault/blinkstick)
[![Go Report Card](https://goreportcard.com/badge/yesnault/blinkstick)](https://goreportcard.com/report/yesnault/blinkstick)

# Blinkstick SDK & CLI

SDK and CLI for [BlinkStick](https://www.blinkstick.com/)

**/!\ Work in progress.**

## Supported device

* Blinckstick Nano

# Hacking

* `$GOPATH/src/github.com/yesnault/blinkstick` contains SDK
* `$GOPATH/src/github.com/yesnault/blinkstick/cli/blink` contains CLI

```bash
mkdir -p $GOPATH/src/github.com/yesnault
cd $GOPATH/src/github.com/yesnault
git clone git@github.com:yesnault/blinkstick.git
cd blinkstick
go install ./...
blink version
```

You've developed a new cool feature? Fixed an annoying bug? We'd be happy
to hear from you! Make sure to read [CONTRIBUTING.md](./CONTRIBUTING.md) before.

# License

This work is under the BSD license, see the [LICENSE](LICENSE) file for details.
