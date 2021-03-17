// Made by SirSAC for Network.
package main
//
import (
	. "fmt"
	"log"
	. "net"
	"os/exec"
	"runtime"
	"syscall"
)
//
func handle_bind(local_connection Conn, remote_connection Conn, load_balancer_addr string, address string, pipe_size int, no_delay bool, keep_alive bool) {
	local_connection.Write([]byte {5, REQUEST_GRANTED, 0, 1, 0, 0, 0, 0, 0, 0})
	log.Println(string(COLOR_BLUE), "[*]", address, "-=>", load_balancer_addr, string(COLOR_RESET))
	go handle_pipe(local_connection, remote_connection, pipe_size, keep_alive)
	go handle_pipe(remote_connection, local_connection, pipe_size, keep_alive)
}
//
func handle_internet(local_connection Conn, remote_address string, pipe_size int, try_count int, no_delay bool, keep_alive bool, serial bool) {
	load_balancer := get_load_balancer(serial)
	local_address, _ := ResolveTCPAddr("tcp", load_balancer.address)
	dialer := Dialer {
		LocalAddr: local_address, Control: func(network, address string, c syscall.RawConn) (error) {
			return c.Control(func(fd uintptr) {
				if err := syscall.BindToDevice(int(fd), load_balancer.iface); err != nil {
					log.Println(string(COLOR_YELLOW), "[!] Couldn't bind to interface", load_balancer.iface, string(COLOR_RESET))
				}
			})
		},
	}
	remote_connection, err := dialer.Dial("tcp", remote_address)
	if err != nil {
		try_again := 0
		for try_again < try_count {
			load_balancer := get_load_balancer(serial)
			local_address, _ := ResolveTCPAddr("tcp", load_balancer.address)
			dialer := Dialer {
				LocalAddr: local_address, Control: func(network, address string, c syscall.RawConn) (error) {
					return c.Control(func(fd uintptr) {
						if err := syscall.BindToDevice(int(fd), load_balancer.iface); err != nil {
							log.Println(string(COLOR_YELLOW), "[!] Couldn't bind to interface", load_balancer.iface, string(COLOR_RESET))
						}
					})
				},
			}
			remote_connection, err := dialer.Dial("tcp", remote_address)
			if err == nil {
				go handle_bind(local_connection, remote_connection, load_balancer.address, remote_address, pipe_size, no_delay, keep_alive)
				return
			}
			try_again++
		}
		log.Println(string(COLOR_YELLOW), "[!]", remote_address, "-=>", load_balancer.address, Sprintf("{%s}", err), string(COLOR_RESET))
		local_connection.Write([]byte {5, NETWORK_UNREACHABLE, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		runtime.Goexit()
	}
	go handle_bind(local_connection, remote_connection, load_balancer.address, remote_address, pipe_size, no_delay, keep_alive)
}
//
func execute_command() {
	if runtime.GOOS == "linux" {
		exec.Command("sudo", `setcap cap_net_raw="eip" "./GoDispatchProxy.bin"`).Run()
		exec.Command("sudo", `sysctl -w net.ipv6.conf.all.accept_local="1"`).Run()
		exec.Command("sudo", `sysctl -w net.ipv6.conf.all.rp_filter="0"`).Run()
		exec.Command("sudo", `sysctl -w net.ipv4.conf.all.accept_local="1"`).Run()
		exec.Command("sudo", `sysctl -w net.ipv4.conf.all.rp_filter="0"`).Run()
		exec.Command("sudo", `ifconfig lo mtu 1280 arp multicast up`).Run()
	}
}
