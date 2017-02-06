[![Build Status](https://travis-ci.org/yesnault/blinkstick.svg?branch=master)](https://travis-ci.org/yesnault/blinkstick)
[![GoDoc](https://godoc.org/github.com/yesnault/blinkstick?status.svg)](https://godoc.org/github.com/yesnault/blinkstick)
[![Go Report Card](https://goreportcard.com/badge/yesnault/blinkstick)](https://goreportcard.com/report/yesnault/blinkstick)

# Blinkstick SDK & CLI

SDK and CLI for [BlinkStick](https://www.blinkstick.com/)

**/!\ Work in progress.**

## Download
See https://github.com/yesnault/blinkstick/releases

## Supported device

* Blinckstick Nano
* Blinckstick Flex
* Blinckstick Strip & Square

## Usage

### Main Menu

``` bash
$blink
Blink - Command Line for Blinkstick

Usage:
  blink [command]

Available Commands:
  color       Color list
  device      device <command>
  flex        flex <command>
  nano        nano <command>
  strip       strip <command>
  update      Update blink to the latest release version: blink update
  version     Display Version of blink: blink version

Flags:
      --log-level string   Log Level : debug, info or warn

Use "blink [command] --help" for more information about a command.
```

### Nano

Example:

```bash
$ blink nano --help
  nano <command>

  Usage:
    blink nano [flags]
    blink nano [command]

  Available Commands:
    color       Color a blinkstick nano: blink nano color [<color>] [--brightness=n] [--top=<color>] [--bottom=<color>] [--serial=s] [--duration=n] [--repeats=n] [--blink]
    list        List all blinkstick nano
```

```bash
$blink nano color --help
  Color a blinkstick nano:

  Set the same color for both led with 50% brightness :
    blink nano color orange --brightness 50

  Set a color for bottom Led and another for top Led:
    blink nano color --bottom red --top green

  Examples:
   blink nano color --top purple --brightness 1
   blink nano color --bottom red --brightness 100
   blink nano color green --brightness 12
   blink nano color --serial BS008173-3.0 --duration=500 --repeats=10 --brightness 30 --blink --bottom red

  Turn off light:
    blink nano color black

  Usage:
    blink nano color [flags]

  Flags:
        --blink            Blink LED
        --bottom string    Color for botton led
        --brightness int   Limit the brightness of the color 0..100 (default 1)
        --duration int     Set duration of transition in milliseconds (use with --blink) (default 100)
        --repeats int      Number of repetitions (use with --blink) (default 10)
        --serial string    Select device by serial number. If unspecified, action will be performed on all BlinkSticks Strip
        --top string       Color for top led: blink nano color red
```

### Flex

Example:

```bash
$ blink flex --help
  flex <command>

  Usage:
    blink flex [flags]
    blink flex [command]

  Available Commands:
    color       Color a blinkstick flex: blink flex color [<color>] [--brightness=n] [--led=n] [--serial=s] [--duration=n] [--repeats=n] [--blink]
    list        List all blinkstick flex

```

```bash
$ blink flex color --help
  Color a blinkstick flex:

  Set the same color for all leds with 50% brightness :
    blink flex color orange --brightness 50

  Color led 0 and 7
   blink flex color red --led 0 --led 7

  Examples:
   blink flex color powderblue
   blink flex color powderblue --brightness 60 --blink --repeats 1
   blink flex color ghostwhite --led 0 --led 2 --led 3 --led 5 --led 7 --led 11 --led 13 --led 17 --led 19 --led 23 --led 29 --led 31

  Turn off light:
    blink flex color black

  Usage:
    blink flex color [flags]

  Flags:
        --blink            Blink LED
        --brightness int   Limit the brightness of the color 0..100 (default 10)
        --duration int     Set duration of transition in milliseconds (use with --blink) (default 100)
        --led intSlice     Led to manipulate: 0..7. If unspecified, action will be performed on all leds
        --repeats int      Number of repetitions (use with --blink) (default 10)
        --serial string    Select device by serial number. If unspecified, action will be performed on all BlinkSticks Flex
```

### Strip & Square

Idem as flex, but max led is 7.

```bash
$ blink strip --help
  strip <command>

  Usage:
    blink strip [flags]
    blink strip [command]

  Available Commands:
    color       Color a blinkstick strip: blink strip color [<color>] [--brightness=n] [--serial=s] [--duration=n] [--repeats=n] [--blink]
    list        List all blinkstick strip
```


```bash
$ blink strip color --help
  Color a blinkstick strip:

  Set the same color for all leds with 50% brightness :
    blink flex color orange --brightness 50

  Color led 0 and 7
   blink flex color red --led 0 --led 7

  Examples:
   blink flex color powderblue
   blink flex color powderblue --brightness 60 --blink --repeats 1
   blink flex color ghostwhite --led 0 --led 2 --led 3 --led 5 --led 7

  Turn off light:
    blink flex color black

  Usage:
    blink strip color [flags]

  Flags:
        --blink            Blink LED
        --brightness int   Limit the brightness of the color 0..100 (default 10)
        --color string     Color for top and bottom led: blink strip color red
        --duration int     Set duration of transition in milliseconds (use with --blink) (default 100)
        --led intSlice     Led to manipulate: 0..7. If unspecified, action will be performed on all leds
        --repeats int      Number of repetitions (use with --blink) (default 10)
        --serial string    Select device by serial number. If unspecified, action will be performed on all BlinkSticks Strip

```

### Color

```
for i in `./blink color list`; do echo $i && ./blink flex color $i; done;
```

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
