// Made by SirSAC for Network.
package main
//
import (
	. "fmt"
	"log"
	. "net"
	"runtime"
)
//
func handle_internet(local_connection Conn, address string, pipe_size int, try_count int, no_delay bool, keep_alive bool, serial bool) {
	remote_address, _ := ResolveTCPAddr("tcp", address)
	load_balancer := get_load_balancer(serial)
	local_address, _ := ResolveTCPAddr("tcp", load_balancer.address)
	remote_connection, err := DialTCP("tcp", local_address, remote_address)
	if err != nil {
		try_again := 0
		for try_again < try_count {
			load_balancer := get_load_balancer(serial)
			local_address, _ := ResolveTCPAddr("tcp", load_balancer.address)
			remote_connection, err := DialTCP("tcp", local_address, remote_address)
			if err == nil {
				go handle_proxy(local_connection, remote_connection, load_balancer.address, address, pipe_size, no_delay, keep_alive)
				return
			}
			try_again++
		}
		log.Println(string(COLOR_YELLOW), "[!]", address, "-=>", load_balancer.address, Sprintf("{%s}", err), string(COLOR_RESET))
		local_connection.Write([]byte {5, NETWORK_UNREACHABLE, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		runtime.Goexit()
	}
	go handle_proxy(local_connection, remote_connection, load_balancer.address, address, pipe_size, no_delay, keep_alive)
}
//
func execute_command() {
	if runtime.GOOS == "darwin" {
	}
}
