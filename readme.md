# DisPro
A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS) load balancing proxy to combine multiple internet connections into one.
It can also be used as a transparent proxy to load balance multiple [SSH](https://en.wikipedia.org/wiki/SSH_(Secure_Shell)) tunnels.
Written in pure [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) with no additional dependencies.
Works on [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows), [Darwin](https://en.wikipedia.org/wiki/MacOS) and [Linux](https://en.wikipedia.org/wiki/Linux).

## Rationale
The idea for this project came from [go-dispatch-proxy](https://github.com/extremecoders-re/go-dispatch-proxy) which is written in [Go](https://en.wikipedia.org/wiki/Go_(programming_language)).
I needed something advanced, portable, stable and light, preferably a single binary without polluting the entire drive and without random issues.

## Installation
No installation required.
Grab the latest binary for your platform from the [releases](https://github.com/SirSAC/DisPro/releases) and start speeding up your internet connection.

## Usage
For [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) use in front of binary name.
```
.\
```
For [Darwin](https://en.wikipedia.org/wiki/MacOS) and [Linux](https://en.wikipedia.org/wiki/Linux) use in front of binary name.
```
./
```
For to show all commands, type this.
```
GoDispatchProxy.bin -help
```
For to show all networks, type this.
```
GoDispatchProxy.bin -list
```
The example 1.
```
GoDispatchProxy.bin -host ::1 -port 1080 -size 4096 -try 2 -delay -alive -serial 192.168.0.2@2 192.168.1.2
```
The example 2.
```
GoDispatchProxy.bin -host ::1 -port 1080 -size 4096 -try 2 -tunnel -delay -alive -serial 192.168.0.2:443@2 192.168.1.2:80
```

## Credits
- [go-dispatch-proxy](https://github.com/extremecoders-re/go-dispatch-proxy): A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS) load balancing proxy to combine multiple internet connections into one.
- [dispatch-proxy](https://github.com/alexkirsz/dispatch-proxy): A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS)/[HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) proxy that balances traffic between multiple internet connections.

## License
Licensed under [MIT](https://github.com/SirSAC/DisPro/blob/master/license).
