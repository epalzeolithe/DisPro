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
func handle_internet(local_connection Conn, address string, processor_thread int, pipe_size int, try_count int, delay_protocol bool, keep_alive bool, serial bool) {
	if serial == true {
		sync_group.Add(processor_thread)
	}
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
				go handle_proxy(local_connection, remote_connection, load_balancer.address, address, pipe_size, delay_protocol, keep_alive)
				if serial == true {
					defer sync_group.Wait()
				}
				return
			}
			try_again++
		}
		log.Println(string(COLOR_YELLOW), "[!]", address, "-=>", load_balancer.address, Sprintf("{%s}", err), string(COLOR_RESET))
		local_connection.Write([]byte {5, NETWORK_UNREACHABLE, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		if serial == true {
			defer sync_group.Wait()
			defer sync_group.Done()
		}
		runtime.Goexit()
	}
	go handle_proxy(local_connection, remote_connection, load_balancer.address, address, pipe_size, delay_protocol, keep_alive)
	if serial == true {
		defer sync_group.Wait()
	}
}
//
func execute_command(option_setting bool) {
}
