:: Made by SirSAC for DisPro.
echo off
:: Clear Screen.
cls
:: Change Directory.
chdir /d %~dp0
:: Help.
dispro.exe -help
netsh.exe interface ipv6 show interfaces
echo Block malware and adult content
echo IPv6 2606:4700:4700::1113 and 2606:4700:4700::1003
echo IPv4 1.1.1.3 and 1.0.0.3
echo Block malware not adult content
echo IPv6 2606:4700:4700::1112 and 2606:4700:4700::1002
echo IPv4 1.1.1.2 and 1.0.0.2
echo Allow malware and adult content
echo IPv6 2606:4700:4700::1111 and 2606:4700:4700::1001
echo IPv4 1.1.1.1 and 1.0.0.1
echo:
dispro.exe -list
echo:
:: Setting.
set /p index_one="Index one>"
set /p index_two="Index two>"
set /p index_three="Index three>"
set /p index_four="Index four>"
set /p ipv6_one="IPv6 one>"
set /p ipv6_two="IPv6 two>"
set /p ipv4_one="IPv4 one>"
set /p ipv4_two="IPv4 two>"
set /p application_argument="Application argument>"
:: Task Kill.
taskkill.exe /im DisPro.exe /t /f
:: Start.
start "LAN" /d %cd% /min /realtime DisPro.exe %application_argument%
:: One.
netsh.exe interface ipv6 set interface %index_one% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_one% mtu=1500 store=active
netsh.exe interface ipv6 set dnsservers name=%index_one% source=static address=%ipv6_one% register=both validate=yes
netsh.exe interface ipv6 add dnsservers name=%index_one% address=%ipv6_two% index=2 validate=yes
netsh.exe interface ipv4 set interface %index_one% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_one% mtu=1500 store=active
netsh.exe interface ipv4 set dnsservers name=%index_one% source=static address=%ipv4_one% register=both validate=yes
netsh.exe interface ipv4 add dnsservers name=%index_one% address=%ipv4_two% index=2 validate=yes
:: Two.
netsh.exe interface ipv6 set interface %index_two% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_two% mtu=1500 store=active
netsh.exe interface ipv6 set dnsservers name=%index_two% source=static address=%ipv6_two% register=both validate=yes
netsh.exe interface ipv6 add dnsservers name=%index_two% address=%ipv6_one% index=2 validate=yes
netsh.exe interface ipv4 set interface %index_two% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_two% mtu=1500 store=active
netsh.exe interface ipv4 set dnsservers name=%index_two% source=static address=%ipv4_two% register=both validate=yes
netsh.exe interface ipv4 add dnsservers name=%index_two% address=%ipv4_one% index=2 validate=yes
:: Three.
netsh.exe interface ipv6 set interface %index_three% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_three% mtu=1500 store=active
netsh.exe interface ipv6 set dnsservers name=%index_three% source=static address=%ipv6_one% register=both validate=yes
netsh.exe interface ipv6 add dnsservers name=%index_three% address=%ipv6_two% index=2 validate=yes
netsh.exe interface ipv4 set interface %index_three% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_three% mtu=1500 store=active
netsh.exe interface ipv4 set dnsservers name=%index_three% source=static address=%ipv4_one% register=both validate=yes
netsh.exe interface ipv4 add dnsservers name=%index_three% address=%ipv4_two% index=2 validate=yes
:: Four.
netsh.exe interface ipv6 set interface %index_four% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_four% mtu=1500 store=active
netsh.exe interface ipv6 set dnsservers name=%index_four% source=static address=%ipv6_two% register=both validate=yes
netsh.exe interface ipv6 add dnsservers name=%index_four% address=%ipv6_one% index=2 validate=yes
netsh.exe interface ipv4 set interface %index_four% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_four% mtu=1500 store=active
netsh.exe interface ipv4 set dnsservers name=%index_four% source=static address=%ipv4_two% register=both validate=yes
netsh.exe interface ipv4 add dnsservers name=%index_four% address=%ipv4_one% index=2 validate=yes
:: IP Configuration.
ipconfig.exe /registerdns
ipconfig.exe /flushdns
