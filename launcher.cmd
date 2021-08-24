:: Made by SirSAC for DisPro.
echo off
:: Clear Screen.
cls
:: Change Directory.
chdir /d %~dp0
:: 0.
dispro.exe -help
netsh.exe interface ipv6 show interfaces
dispro.exe -list
:: Setting.
set /p index_one="Index one>"
set /p index_two="Index two>"
set /p index_three="Index three>"
set /p index_four="Index four>"
set /p application_argument="Application argument>"
:: Task Kill.
taskkill.exe /im DisPro.exe /t /f
:: Start.
start "LAN" /d %cd% /min /realtime DisPro.exe %application_argument%
:: One.
netsh.exe interface ipv6 set interface %index_one% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_one% mtu=1500 store=active
netsh.exe interface ipv6 set dnsservers name=%index_one% source=static address=2606:4700:4700::1113 register=both validate=yes
netsh.exe interface ipv6 add dnsservers name=%index_one% address=2606:4700:4700::1003 index=2 validate=yes
netsh.exe interface ipv4 set interface %index_one% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_one% mtu=1500 store=active
netsh.exe interface ipv4 set dnsservers name=%index_one% source=static address=1.1.1.3 register=both validate=yes
netsh.exe interface ipv4 add dnsservers name=%index_one% address=1.0.0.3 index=2 validate=yes
:: Two.
netsh.exe interface ipv6 set interface %index_two% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_two% mtu=1500 store=active
netsh.exe interface ipv6 set dnsservers name=%index_two% source=static address=2606:4700:4700::1003 register=both validate=yes
netsh.exe interface ipv6 add dnsservers name=%index_two% address=2606:4700:4700::1113 index=2 validate=yes
netsh.exe interface ipv4 set interface %index_two% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_two% mtu=1500 store=active
netsh.exe interface ipv4 set dnsservers name=%index_two% source=static address=1.0.0.3 register=both validate=yes
netsh.exe interface ipv4 add dnsservers name=%index_two% address=1.1.1.3 index=2 validate=yes
:: Three.
netsh.exe interface ipv6 set interface %index_three% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_three% mtu=1500 store=active
netsh.exe interface ipv6 set dnsservers name=%index_three% source=static address=2606:4700:4700::1113 register=both validate=yes
netsh.exe interface ipv6 add dnsservers name=%index_three% address=2606:4700:4700::1003 index=2 validate=yes
netsh.exe interface ipv4 set interface %index_three% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_three% mtu=1500 store=active
netsh.exe interface ipv4 set dnsservers name=%index_three% source=static address=1.1.1.3 register=both validate=yes
netsh.exe interface ipv4 add dnsservers name=%index_three% address=1.0.0.3 index=2 validate=yes
:: Four.
netsh.exe interface ipv6 set interface %index_four% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv6 set subinterface interface=%index_four% mtu=1500 store=active
netsh.exe interface ipv6 set dnsservers name=%index_four% source=static address=2606:4700:4700::1003 register=both validate=yes
netsh.exe interface ipv6 add dnsservers name=%index_four% address=2606:4700:4700::1113 index=2 validate=yes
netsh.exe interface ipv4 set interface %index_four% forwarding=disabled advertise=disabled metric=2 ignoredefaultroutes=disabled advertisedefaultroute=disabled store=active ecncapability=application
netsh.exe interface ipv4 set subinterface interface=%index_four% mtu=1500 store=active
netsh.exe interface ipv4 set dnsservers name=%index_four% source=static address=1.0.0.3 register=both validate=yes
netsh.exe interface ipv4 add dnsservers name=%index_four% address=1.1.1.3 index=2 validate=yes
:: IP Configuration.
ipconfig.exe /registerdns
ipconfig.exe /flushdns
:: Timeout.
timeout.exe /t 1
:: Echoing.
echo:
