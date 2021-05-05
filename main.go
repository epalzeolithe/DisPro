// Made by SirSAC for Network.
package main
//
import "C"
//
import (
	"bufio"
	"bytes"
	. "flag"
	. "fmt"
	"io"
	"log"
	. "net"
	"os"
	"os/exec"
	"runtime"
	. "runtime/debug"
	. "strconv"
	"strings"
	. "sync"
)
//
type load_balancer struct {
	address string
	iface string
	contention_ratio int
	current_connections int
}
//
var (
	sync_group *WaitGroup
	sync_mutex *Mutex
	lb_list []load_balancer
	lb_index int = 0
	buffer_byte bytes.Buffer
)
//
func get_load_balancer(serial bool) (lb *load_balancer) {
	if serial == true {
		sync_mutex.Lock()
	}
	lb = &lb_list[lb_index]
	lb.current_connections += 1
	if lb.current_connections == lb.contention_ratio {
		lb.current_connections = 0
		lb_index += 1
		if lb_index == len(lb_list) {
			lb_index = 0
		}
	}
	if serial == true {
		defer sync_group.Done()
		defer sync_mutex.Unlock()
	}
	return
}
//
func handle_pipe(source_packet Conn, destination_packet Conn, buffer_size int, keep_alive bool) {
	if keep_alive == false {
		defer source_packet.Close()
		defer destination_packet.Close()
	}
	buffer_packet := buffer_byte.Bytes()
	buffer_source := bufio.NewReaderSize(source_packet, buffer_size)
	buffer_destination := bufio.NewWriterSize(destination_packet, buffer_size)
	_, err := io.CopyBuffer(buffer_destination, buffer_source, buffer_packet)
	if err != nil {
		runtime.Goexit()
	}
}
//
func handle_proxy(local_connection Conn, remote_connection *TCPConn, load_balancer_addr string, address string, pipe_size int, delay_protocol bool, keep_alive bool) {
	remote_connection.SetNoDelay(delay_protocol)
	remote_connection.SetKeepAlive(keep_alive)
	if keep_alive == false {
		remote_connection.SetLinger(0)
	}
	local_connection.Write([]byte {5, REQUEST_GRANTED, 0, 1, 0, 0, 0, 0, 0, 0})
	log.Println(string(COLOR_BLUE), "[*]", address, "-=>", load_balancer_addr, string(COLOR_RESET))
	go handle_pipe(local_connection, remote_connection, pipe_size, keep_alive)
	go handle_pipe(remote_connection, local_connection, pipe_size, keep_alive)
}
//
func handle_tunnel(local_connection Conn, processor_thread int, pipe_size int, try_count int, delay_protocol bool, keep_alive bool, serial bool) {
	if serial == true {
		sync_group.Add(processor_thread)
	}
	print_message := string("Tunnelled")
	load_balancer := get_load_balancer(serial)
	remote_address, _ := ResolveTCPAddr("tcp", load_balancer.address)
	remote_connection, err := DialTCP("tcp", nil, remote_address)
	if err != nil {
		try_again := 0
		for try_again < try_count {
			load_balancer := get_load_balancer(serial)
			remote_address, _ := ResolveTCPAddr("tcp", load_balancer.address)
			remote_connection, err := DialTCP("tcp", nil, remote_address)
			if err == nil {
				go handle_proxy(local_connection, remote_connection, load_balancer.address, print_message, pipe_size, delay_protocol, keep_alive)
				if serial == true {
					defer sync_group.Wait()
				}
				return
			}
			try_again++
		}
		log.Println(string(COLOR_YELLOW), "[!]", load_balancer.address, Sprintf("{%s}", err), string(COLOR_RESET))
		local_connection.Write([]byte {5, HOST_UNREACHABLE, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		if serial == true {
			defer sync_group.Wait()
			defer sync_group.Done()
		}
		runtime.Goexit()
	}
	go handle_proxy(local_connection, remote_connection, load_balancer.address, print_message, pipe_size, delay_protocol, keep_alive)
	if serial == true {
		defer sync_group.Wait()
	}
}
//
func detect_interfaces(address_network string, list_network bool) (string) {
	if list_network == true {
		log.Println(string(COLOR_CYAN), "[-] Listing the available addresses for dispatching", string(COLOR_RESET))
	}
	ifaces, _ := Interfaces()
	for _, iface := range ifaces {
		if (iface.Flags&FlagUp == FlagUp) && (iface.Flags&FlagLoopback != FlagLoopback) {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				if ipnet, ok := addr.(*IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To16() != nil {
						if list_network == false {
							if ipnet.IP.String() == address_network {
								return iface.Name
							}
						}
						if list_network == true {
							log.Printf("%s [+] %s, IP: %s %s\n", string(COLOR_MAGENTA), iface.Name, ipnet.IP.String(), string(COLOR_RESET))
						}
					}
				}
			}
		}
	}
	return ""
}
//
func parse_network(argument_network []string, tunnel bool, list_network bool) (mtu_standard int, try_standard int) {
	if len(argument_network) == 0 {
		log.Fatalln(string(COLOR_RED), "[x] Please specify one or more network addresses", string(COLOR_RESET))
	}
	lb_list = make([]load_balancer, NArg())
	for idx, a := range argument_network {
		splitted := strings.Split(a, "@")
		var (
			address_network string
			port_network int
			err error
			cont_ratio int = 1
		)
		if tunnel == true {
			ip_port := strings.Split(splitted[0], ":")
			if len(ip_port) != 2 {
				log.Fatalln(string(COLOR_RED), "[x] Invalid address specification", splitted[0], string(COLOR_RESET))
			}
			address_network = ip_port[0]
			port_network, err = Atoi(ip_port[1])
			if err != nil || port_network < 0 || port_network > 65535 {
				log.Fatalln(string(COLOR_RED), "[x] Invalid port", splitted[0], string(COLOR_RESET))
			}
		}
		if tunnel == false {
			address_network = splitted[0]
			port_network = 0
		}
		if ParseIP(address_network).To16() == nil {
			log.Fatalln(string(COLOR_RED), "[x] Invalid address", address_network, string(COLOR_RESET))
		}
		if len(splitted) > 1 {
			cont_ratio, err = Atoi(splitted[1])
			if err != nil || cont_ratio <= 0 {
				log.Fatalln(string(COLOR_RED), "[x] Invalid contention ratio for", address_network, string(COLOR_RESET))
			}
		}
		iface := detect_interfaces(address_network, list_network)
		if iface == "" {
			log.Fatalln(string(COLOR_RED), "[x] IP address not associated with an interface", address_network, string(COLOR_RESET))
		}
		log.Printf("%s [i] Load balancer %s: %s, contention ratio: %d %s\n", string(COLOR_GREEN), iface, address_network, cont_ratio, string(COLOR_RESET))
		lb_list[idx] = load_balancer {address: Sprintf("%s:%d", address_network, port_network), iface: iface, contention_ratio: cont_ratio, current_connections: 0}
		mtu_standard = mtu_standard + 1500
		try_standard = try_standard + cont_ratio
	}
	return
}
//
func handle_network(local_connection Conn, processor_thread int, pipe_size int, try_count int, tunnel bool, secure_connection bool, delay_protocol bool, keep_alive bool, serial bool) {
	if tunnel == false {
		address, err := handle_socks(local_connection, secure_connection)
		if err == nil {
			go handle_internet(local_connection, address, processor_thread, pipe_size, try_count, delay_protocol, keep_alive, serial)
		}
	}
	if tunnel == true {
		go handle_tunnel(local_connection, processor_thread, pipe_size, try_count, delay_protocol, keep_alive, serial)
	}
}
//
func main() {
	if runtime.GOOS == "windows" {
		exec.Command("powershell.exe").Run()
	}
	processor_thread := runtime.NumCPU()
	var (
		host = String("host", "::1", "The IP address to listen for SOCKS connection")
		port = Int("port", 1080, "The port number to listen for SOCKS connection")
		multiply = Int("multiply", 2, "The threads are multiplied by the specific value")
		percent = Int("percent", 200, "The value in percent for garbage collection")
		pipe = Int("pipe", 8192, "The size of buffers in bytes for more throughput")
		tunnel = Bool("tunnel", false, "Use tunnel mode (acts as a transparent load balancing proxy)")
		option = Bool("option", false, "Use option mode (sets the operating system options for maximum potential)")
		secure = Bool("secure", false, "Use secure mode (acts like using secure ports than usual ones)")
		delay = Bool("delay", false, "Use delay mode (acts a combining a number of small outgoing messages and sending them all at once)")
		keep = Bool("keep", false, "Use keep mode (sets whether the operating system should send keep-alive messages on the connection)")
		serial = Bool("serial", false, "Use serial mode (acts to serialize access to function get load balancer)")
		list = Bool("list", false, "Shows the available addresses for dispatching (non-tunneling mode only)")
	)
	runtime.GOMAXPROCS(processor_thread * *multiply)
	SetGCPercent(*percent)
	Parse()
	if *list == true {
		detect_interfaces("", *list)
		os.Exit(0)
	}
	if ParseIP(*host).To16() == nil {
		log.Fatalln(string(COLOR_RED), "[x] Invalid host", *host, string(COLOR_RESET))
	}
	if *port < 1 || *port > 65535 {
		log.Fatalln(string(COLOR_RED), "[x] Invalid port", *port, string(COLOR_RESET))
	}
	local_port := Itoa(*port)
	host_port := JoinHostPort(*host, local_port)
	local_host, err := Listen("tcp", host_port)
	if err != nil {
		log.Fatalln(string(COLOR_RED), "[x] Could not start local server on", host_port, string(COLOR_RESET))
	}
	mtu_jumbo, try_number := parse_network(Args(), *tunnel, *list)
	mtu_size := Itoa(mtu_jumbo)
	execute_command(mtu_size, *option)
	try_count := try_number - 1
	log.Println(string(COLOR_GREEN), "[i] Local server started on", host_port, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Jumbo size is", mtu_jumbo, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Multiply thread is", *multiply, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Percent ratio is", *percent, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Pipe size is", *pipe, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Try count is", try_count, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Tunnel is", *tunnel, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Option setting is", *option, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Secure connection is", *secure, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Delay protocol is", *delay, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Keep alive is", *keep, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Serial order is", *serial, string(COLOR_RESET))
	if *delay == false {
		*delay = true
	}
	if *serial == true {
		sync_group = &WaitGroup {}
		sync_mutex = &Mutex {}
	}
	for {
		local_connection, _ := local_host.Accept()
		go handle_network(local_connection, processor_thread, *pipe, try_count, *tunnel, *secure, *delay, *keep, *serial)
		if *serial == false {
			continue
		}
	}
}
