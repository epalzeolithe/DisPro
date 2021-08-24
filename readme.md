![GoLang](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1024px-Go_Logo_Blue.svg.png)

# DisPro
A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS) load balancing proxy to combine multiple internet connections into one. <br>
It can also be used as a transparent proxy to load balance multiple [SSH](https://en.wikipedia.org/wiki/SSH_(Secure_Shell)) tunnels. <br>
Written in pure [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) with no additional dependencies. <br>
Works on [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows), [Darwin](https://en.wikipedia.org/wiki/MacOS) and [Linux](https://en.wikipedia.org/wiki/Linux).

## Rationale
The idea for this project came from [go-dispatch-proxy](https://github.com/extremecoders-re/go-dispatch-proxy) which is written in [Go](https://en.wikipedia.org/wiki/Go_(programming_language)). <br>
I needed something advanced, secure, fast, stable, portable and light, preferably a single binary that is independent by location without polluting the entire drive and without security or stability issues.

## Announces
Up to date.

## Installation
No installation required. <br>
Just grab the latest binary for your platform from the [releases](https://github.com/SirSAC/DisPro/releases/latest) and rename it `DisPro.exe` or `DisPro.bin` if you are not on [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) then start speeding up your internet connection. <br>
For `-option` on [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) will require the [Privoxy](https://en.wikipedia.org/wiki/Privoxy) software and this can be downloaded from [release](https://sourceforge.net/projects/ijbswa/files/Win32/3.0.32%20%28stable%29) and make sure it is listening on [IP](https://en.wikipedia.org/wiki/IP_address) `::1` and port `8118`. <br>
If you want to use [`cmd.exe`](https://en.wikipedia.org/wiki/Cmd.exe) on [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) this will require the [`powershell.exe`](https://en.wikipedia.org/wiki/PowerShell) to be already installed for fixing text colors but the program can work without it if you don't want text colors anymore. <br>
This program can be used with `launcher.cmd` for insane speed and workaround of any bugs on [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows).

## Informations
DisPro now supports [Linux](https://en.wikipedia.org/wiki/Linux) in both modes, normal and tunnel. <br>
On [Linux](https://en.wikipedia.org/wiki/Linux) normal mode, DisPro uses the `SOL_SOCKET` and `SO_BINDTODEVICE` from [`syscall`](https://golang.org/pkg/syscall/#BindToDevice) package to bind the interface corresponding to the load balancer [IP](https://en.wikipedia.org/wiki/IP_address) addresses. <br>
As a result, the binary must be run with necessary capabilities and with admin or root privilege. <br>
The adaptive [MTU](https://en.wikipedia.org/wiki/Maximum_transmission_unit) changing is automaticaly and increases or decreases for each load balancer address, this function is named as Jumbo in program, these changes are not permanent and will be reset them to the default ones when you restart or reboot but not when you hibernate or sleep your computer, im not sure about shotdown. <br>
The `-secure` is smarter now that means if a website or address support to be forwarded on secure port that will use it, if is not supported that will be on unsecure port for the website or address to can be accessed. <br>
The option `-try` is removed on new updates because now is automatic and is essential to be always enabled for preventing the issues if an load balancer have some troubles or if is down.
### For to bypass the [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) and [Linux](https://en.wikipedia.org/wiki/Linux) conflicts, the following commands will be automaticaly executed.
On [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) will be.
```
cmd.exe /c wmic.exe process where 'name='DisPro.exe'' call setpriority realtime
cmd.exe /c net.exe stop /y RemoteAccess
cmd.exe /c netsh.exe interface ipv6 set interface 1 forwarding=enabled advertise=disabled metric=1 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
cmd.exe /c netsh.exe interface ipv6 set subinterface interface=1 mtu=(Adaptive MTU) store=active
cmd.exe /c netsh.exe interface ipv6 set address interface=1 type=anycast store=active
cmd.exe /c netsh.exe interface ipv6 set address interface=1 address=::1 type=unicast validlifetime=infinite preferredlifetime=infinite store=active
cmd.exe /c netsh.exe interface ipv4 set interface 1 forwarding=enabled advertise=disabled metric=1 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
cmd.exe /c netsh.exe interface ipv4 set subinterface interface=1 mtu=(Adaptive MTU) store=active
cmd.exe /c netsh.exe interface ipv4 set address name=1 source=dhcp type=anycast store=active
cmd.exe /c netsh.exe interface ipv4 set address name=1 source=static address=127.0.0.1 mask=255.255.255.255 gwmetric=1 type=unicast store=active
```
On [Linux](https://en.wikipedia.org/wiki/Linux) will be.
```
sh -c chmod --verbose 0755 ./DisPro.bin
sh -c chown -c root:daemon ./DisPro.bin
sh -c setcap cap_net_admin,cap_net_raw=eip ./DisPro.bin
sh -c ifconfig -a lo add 127.0.0.1 netmask 255.255.255.255 mtu (Adaptive MTU) arp allmulti multicast dynamic up
sh -c sysctl --write net.ipv4.conf.all.rp_filter=0
```
Tunnel mode doesn't require admin or root privilege. <br>
Ignore the `-option` for [Darwin](https://en.wikipedia.org/wiki/MacOS), it will do nothing.
### For the `launcher.cmd` to work, the following commands will be automaticaly executed.
On [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) will be.
```
:: One.
netsh.exe interface ipv6 set interface %index_one% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_one% mtu=1500 store=active
netsh.exe interface ipv4 set interface %index_one% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_one% mtu=1500 store=active
:: Two.
netsh.exe interface ipv6 set interface %index_two% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_two% mtu=1500 store=active
netsh.exe interface ipv4 set interface %index_two% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_two% mtu=1500 store=active
:: Three.
netsh.exe interface ipv6 set interface %index_three% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_three% mtu=1500 store=active
netsh.exe interface ipv4 set interface %index_three% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_three% mtu=1500 store=active
:: Four.
netsh.exe interface ipv6 set interface %index_four% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_four% mtu=1500 store=active
netsh.exe interface ipv4 set interface %index_four% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_four% mtu=1500 store=active
```
The rest of code is not important to be listed here.

## Warning
The `-option` and `launcher.cmd` will change the operating system settings and some of these can be permanent but can be reset them to the default values manually if you know them.
### For [Windows](https://en.wikipedia.org/wiki/Microsoft_Windows) the settings will be these.
<https://github.com/SirSAC/DisPro/blob/master/windows.go#L63-L80> <br>
Changing the [DNS](https://en.wikipedia.org/wiki/Domain_Name_System) servers from the `launcher.cmd` will modify the system settings for each selected internet adapters.
### For [Linux](https://en.wikipedia.org/wiki/Linux) the settings will be these.
<https://github.com/SirSAC/DisPro/blob/master/linux.go#L83>
### For noobs please do not use the `-option` command.
<https://github.com/SirSAC/DisPro/blob/master/readme.md#L19>

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
DisPro.(exe/bin) -help
```
Will show this.
```
Usage of C:\Users\SirSAC\Downloads\DisPro.(exe/bin):
	-host string
		The IP address to listen for SOCKS connection (default "::1")
	-keep
		Use keep mode (sets whether the program should keep the connection alive even if it is done)
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
	-tunnel
		Use tunnel mode (acts as a transparent load balancing proxy)
PS C:\Users\SirSAC\Downloads>
```
### For to show all networks, type this.
```
DisPro.(exe/bin) -list
```
Will show like this.
```
2021/03/21 17:15:31 [-] Listing the available addresses for dispatching
2021/03/21 17:15:31 [+] 1, IP: fe80::e464:1874:c2d3:ec5
2021/03/21 17:15:31 [+] 1, IP: 192.168.1.2
2021/03/21 17:15:31 [+] 2, IP: fe80::4843:2f14:2f6a:81a8
2021/03/21 17:15:31 [+] 2, IP: 192.168.43.4
PS C:\Users\SirSAC\Downloads>
```
The example 1.
```
DisPro.(exe/bin) -host ::1 -port 1080 -multiply 2 -pipe 8192 -option -secure -delay -keep -serial 192.168.1.2 192.168.43.4@1
```
The example 2.
```
DisPro.(exe/bin) -host 127.0.0.1 -port 8080 -multiply 1 -pipe 4096 -tunnel -option -delay -keep -serial 127.0.0.1:443 127.0.0.1:80@2
```
Please do not use the examples, is just informative.

## Credits
- [go-dispatch-proxy](https://github.com/extremecoders-re/go-dispatch-proxy): A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS) load balancing proxy to combine multiple internet connections into one.
- [dispatch-proxy](https://github.com/alexkirsz/dispatch-proxy): A [SOCKS5](https://en.wikipedia.org/wiki/SOCKS)/[HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) proxy that balances traffic between multiple internet connections.

## Conduct
Conduct under the [Contributor Covenant Code of Conduct](https://github.com/SirSAC/DisPro/blob/master/code_of_conduct.md).

## License
Licensed under the [BSD 3-Clause "New" or "Revised" License](https://github.com/SirSAC/DisPro/blob/master/license.md).
