# GoDispatchProxy
Made by SirSAC for Network.

## Rationale
The idea for this project came from [go-dispatch-proxy](https://github.com/extremecoders-re/go-dispatch-proxy) which is written in Go.
I needed something advanced, portable, stable & light, preferably a single binary without polluting the entire drive and without random issues.

## Installation
No installation required.
Grab the latest binary for your platform from the [releases](https://github.com/SirSAC/GoDispatchProxy/releases) and start speeding up your internet connection.

## Usage
For Windows use.
```
.\GoDispatchProxy.bin
```
For Darwin and Linux use.
```
./GoDispatchProxy.bin
```
For to show all commands type this.
```
GoDispatchProxy.bin -help
```
For to show all networks type this.
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
- [go-dispatch-proxy](https://github.com/extremecoders-re/go-dispatch-proxy): A SOCKS5 load balancing proxy to combine multiple internet connections into one.
- [dispatch-proxy](https://github.com/alexkirsz/dispatch-proxy): A SOCKS5/HTTP proxy that balances traffic between multiple internet connections.

## License
Licensed under [MIT](https://github.com/SirSAC/GoDispatchProxy/blob/main/LICENSE).
