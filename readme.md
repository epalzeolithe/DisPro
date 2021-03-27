# DisPro
A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS) load balancing proxy to combine multiple internet connections into one.
It can also be used as a transparent proxy to load balance multiple [SSH](https://en.wikipedia.org/wiki/SSH_(Secure_Shell)) tunnels.
Written in pure [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) with no additional dependencies.
Works on [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows), [Darwin](https://en.wikipedia.org/wiki/MacOS) and [Linux](https://en.wikipedia.org/wiki/Linux).

## Rationale
The idea for this project came from [go-dispatch-proxy](https://github.com/extremecoders-re/go-dispatch-proxy) which is written in [Go](https://en.wikipedia.org/wiki/Go_(programming_language)).
I needed something advanced, secure, stable, portable and light, preferably a single binary that is independent by location without polluting the entire drive and without security or stability issues.

## Installation
No installation required.
Just grab the latest binary for your platform from the [releases](https://github.com/SirSAC/DisPro/releases/latest) and rename it `DisPro.bin` then start speeding up your internet connection.

## Informations
DisPro now supports [Linux](https://en.wikipedia.org/wiki/Linux) in both modes, normal and tunnel.
On [Linux](https://en.wikipedia.org/wiki/Linux) normal mode, DisPro uses the `SOL_SOCKET` and `SO_BINDTODEVICE` from [`syscall`](https://golang.org/pkg/syscall/#BindToDevice) package to bind the interface corresponding to the load balancer [IP](https://en.wikipedia.org/wiki/IP_address) addresses.
As a result, the binary must be run with necessary capabilities and with admin or root privilege.
### For to bypass the [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) and [Linux](https://en.wikipedia.org/wiki/Linux) conflicts, the following commands will be automaticaly executed.
On [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) will be.
```
cmd.exe /c net.exe stop /y RemoteAccess
```
On [Linux](https://en.wikipedia.org/wiki/Linux) will be.
```
sh -c sysctl --write net.ipv4.conf.all.rp_filter=0
```
Tunnel mode doesn't require admin or root privilege.

## Usage
For [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) use in front of binary name.
```
.\
```
For [Darwin](https://en.wikipedia.org/wiki/MacOS) and [Linux](https://en.wikipedia.org/wiki/Linux) use in front of binary name.
```
./
```
### For to show all commands, type this.
```
DisPro.bin -help
```
Will show this.
```
Usage of C:\Users\SirSAC\Downloads\DisPro.bin:
	-delay
		Use delay mode (acts a combining a number of small outgoing messages and sending them all at once)
	-host string
		The IP address to listen for SOCKS connection (default "::1")
	-keep
		Use keep mode (sets whether the operating system should send keep-alive messages on the connection)
	-list
		Shows the available addresses for dispatching (non-tunneling mode only)
	-multiply int
		The threads are multiplied by the specific value (default 2)
	-option
		Use option mode (sets the operating system options for maximum potential)
	-pipe int
		The size of buffers in bytes for more throughput (default 8192)
	-port int
		The port number to listen for SOCKS connection (default 1080)
	-secure
		Use secure mode (acts like using secure ports than usual ones)
	-serial
		Use serial mode (acts to serialize access to function get load balancer)
	-try int
		The number of retries for SOCKS connection (default 0)
	-tunnel
		Use tunnel mode (acts as a transparent load balancing proxy)
PS C:\Users\SirSAC\Downloads>
```
### For to show all networks, type this.
```
DisPro.bin -list
```
Will show like this.
```
2021/03/21 17:15:31 [-] Listing the available addresses for dispatching
2021/03/21 17:15:31 [+] 1, IP: fe80::e464:1874:c2d3:ec5
2021/03/21 17:15:31 [+] 1, IP: 192.168.1.2
2021/03/21 17:15:31 [+] 2, IP: fe80::58c1:394f:7f9f:bc65
2021/03/21 17:15:31 [+] 2, IP: 192.168.43.118
2021/03/21 17:15:31 [+] 4, IP: fe80::5562:b18d:553f:c48d
2021/03/21 17:15:31 [+] 4, IP: 169.254.196.141
PS C:\Users\SirSAC\Downloads>
```
The example 1.
```
DisPro.bin -host ::1 -port 1080 -multiply 2 -pipe 8192 -try 1 -option -secure -delay -keep -serial 192.168.1.2 192.168.43.118@1
```
The example 2.
```
DisPro.bin -host 127.0.0.1 -port 8080 -multiply 1 -pipe 4096 -try 2 -tunnel -option -delay -keep -serial 127.0.0.1:443 127.0.0.1:80@2
```

## Credits
- [go-dispatch-proxy](https://github.com/extremecoders-re/go-dispatch-proxy): A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS) load balancing proxy to combine multiple internet connections into one.
- [dispatch-proxy](https://github.com/alexkirsz/dispatch-proxy): A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS)/[HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) proxy that balances traffic between multiple internet connections.

## Conduct
Conduct under the [Contributor Covenant Code of Conduct](https://github.com/SirSAC/DisPro/blob/master/code_of_conduct.md).

## License
Licensed under the [BSD 3-Clause "New" or "Revised" License](https://github.com/SirSAC/DisPro/blob/master/license.md).
