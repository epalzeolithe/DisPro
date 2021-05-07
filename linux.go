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
func handle_bind(local_connection Conn, remote_connection Conn, balancer_address string, target_address string, pipe_size int, delay_protocol bool, keep_alive bool) {
	local_connection.Write([]byte {5, REQUEST_GRANTED, 0, 1, 0, 0, 0, 0, 0, 0})
	log.Println(string(COLOR_BLUE), "[*]", target_address, "-=>", balancer_address, string(COLOR_RESET))
	go handle_pipe(local_connection, remote_connection, pipe_size, keep_alive)
	go handle_pipe(remote_connection, local_connection, pipe_size, keep_alive)
}
//
func handle_internet(local_connection Conn, target_address string, processor_thread int, pipe_size int, try_count int, delay_protocol bool, keep_alive bool, serial_order bool) {
	if serial_order == true {
		sync_group.Add(processor_thread)
	}
	load_balancer := get_load_balancer(serial_order)
	local_address, _ := ResolveTCPAddr("tcp", load_balancer.address)
	network_dialer := Dialer {
		LocalAddr: local_address, Control: func(network_protocol string, remote_address string, c syscall.RawConn) (err error) {
			return c.Control(func(fd uintptr) {
				err := syscall.BindToDevice(int(fd), load_balancer.interface_name)
				if err != nil {
					log.Println(string(COLOR_YELLOW), "[!] Couldn't bind to interface", load_balancer.interface_name, string(COLOR_RESET))
				}
			})
		},
	}
	remote_connection, err := network_dialer.Dial("tcp", target_address)
	if err != nil {
		try_again := 0
		for try_again < try_count {
			load_balancer := get_load_balancer(serial_order)
			local_address, _ := ResolveTCPAddr("tcp", load_balancer.address)
			network_dialer := Dialer {
				LocalAddr: local_address, Control: func(network_protocol string, remote_address string, c syscall.RawConn) (err error) {
					return c.Control(func(fd uintptr) {
						err := syscall.BindToDevice(int(fd), load_balancer.interface_name)
						if err != nil {
							log.Println(string(COLOR_YELLOW), "[!] Couldn't bind to interface", load_balancer.interface_name, string(COLOR_RESET))
						}
					})
				},
			}
			remote_connection, err := network_dialer.Dial("tcp", target_address)
			if err == nil {
				go handle_bind(local_connection, remote_connection, load_balancer.address, target_address, pipe_size, delay_protocol, keep_alive)
				if serial_order == true {
					defer sync_group.Wait()
				}
				return
			}
			try_again++
		}
		log.Println(string(COLOR_YELLOW), "[!]", target_address, "-=>", load_balancer.address, Sprintf("{%s}", err), string(COLOR_RESET))
		local_connection.Write([]byte {5, NETWORK_UNREACHABLE, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		if serial_order == true {
			defer sync_group.Wait()
			defer sync_group.Done()
		}
		runtime.Goexit()
	}
	go handle_bind(local_connection, remote_connection, load_balancer.address, target_address, pipe_size, delay_protocol, keep_alive)
	if serial_order == true {
		defer sync_group.Wait()
	}
}
//
func execute_command(mtu_size string, option_setting bool) {
	exec.Command("sh", "-c", "chmod --verbose 0755 ./DisPro.bin").Run()
	exec.Command("sh", "-c", "chown -c root:daemon ./DisPro.bin").Run()
	exec.Command("sh", "-c", "setcap cap_net_admin,cap_net_raw=eip ./DisPro.bin").Run()
	exec.Command("sh", "-c", "ifconfig -a lo add 127.0.0.1 netmask 255.255.255.255 mtu", mtu_size, "arp allmulti multicast dynamic up").Run()
	if option_setting == true {
		exec.Command("sh", "-c", "sysctl --write net.ipv4.conf.all.accept_local=1").Run()
	}
	exec.Command("sh", "-c", "sysctl --write net.ipv4.conf.all.rp_filter=0").Run()
}
