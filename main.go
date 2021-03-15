// Made by SirSAC for Network.
package main
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
var mutex *RWMutex
var lb_list []load_balancer
var lb_index int = 0
var buffer_byte bytes.Buffer
//
func get_load_balancer(serial bool) (*load_balancer) {
	if serial == true {
		mutex.Lock()
	}
	lb := &lb_list[lb_index]
	lb.current_connections += 1
	if lb.current_connections == lb.contention_ratio {
		lb.current_connections = 0
		lb_index += 1
		if lb_index == len(lb_list) {
			lb_index = 0
		}
	}
	if serial == true {
		defer mutex.Unlock()
	}
	return lb
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
func handle_proxy(local_connection Conn, remote_connection *TCPConn, load_balancer_addr string, address string, pipe_size int, no_delay bool, keep_alive bool) {
	remote_connection.SetNoDelay(no_delay)
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
func handle_tunnel(local_connection Conn, pipe_size int, try_count int, no_delay bool, keep_alive bool, serial bool) {
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
				go handle_proxy(local_connection, remote_connection, load_balancer.address, print_message, pipe_size, no_delay, keep_alive)
				return
			}
			try_again++
		}
		log.Println(string(COLOR_YELLOW), "[!]", load_balancer.address, Sprintf("{%s}", err), string(COLOR_RESET))
		local_connection.Write([]byte {5, HOST_UNREACHABLE, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		runtime.Goexit()
	}
	go handle_proxy(local_connection, remote_connection, load_balancer.address, print_message, pipe_size, no_delay, keep_alive)
}
//
func detect_interfaces(ip string, list bool) (string) {
	if list == true {
		log.Println(string(COLOR_CYAN), "[-] Listing the available addresses for dispatching", string(COLOR_RESET))
	}
	ifaces, _ := Interfaces()
	for _, iface := range ifaces {
		if (iface.Flags&FlagUp == FlagUp) && (iface.Flags&FlagLoopback != FlagLoopback) {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				if ipnet, ok := addr.(*IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To16() != nil {
						if list == false {
							if ipnet.IP.String() == ip {
								return iface.Name + "\x00"
							}
						}
						if list == true {
							log.Printf("%s [+] %s, IP: %s %s\n", string(COLOR_PURPLE), iface.Name, ipnet.IP.String(), string(COLOR_RESET))
						}
					}
				}
			}
		}
	}
	return ""
}
//
func parse_load_balancers(args []string, tunnel bool, list bool) {
	if len(args) == 0 {
		log.Fatalln(string(COLOR_RED), "[x] Please specify one or more load balancers", string(COLOR_RESET))
	}
	lb_list = make([]load_balancer, NArg())
	for idx, a := range args {
		splitted := strings.Split(a, "@")
		var lb_ip string
		var lb_port int
		var err error
		if tunnel == true {
			ip_port := strings.Split(splitted[0], ":")
			if len(ip_port) != 2 {
				log.Fatalln(string(COLOR_RED), "[x] Invalid address specification", splitted[0], string(COLOR_RESET))
			}
			lb_ip = ip_port[0]
			lb_port, err = Atoi(ip_port[1])
			if err != nil || lb_port < 0 || lb_port > 65535 {
				log.Fatalln(string(COLOR_RED), "[x] Invalid port", splitted[0], string(COLOR_RESET))
			}
		}
		if tunnel == false {
			lb_ip = splitted[0]
			lb_port = 0
		}
		if ParseIP(lb_ip).To16() == nil {
			log.Fatalln(string(COLOR_RED), "[x] Invalid address", lb_ip, string(COLOR_RESET))
		}
		var cont_ratio int = 1
		if len(splitted) > 1 {
			cont_ratio, err = Atoi(splitted[1])
			if err != nil || cont_ratio <= 0 {
				log.Fatalln(string(COLOR_RED), "[x] Invalid contention ratio for", lb_ip, string(COLOR_RESET))
			}
		}
		iface := detect_interfaces(lb_ip, list)
		if iface == "" {
			log.Fatalln(string(COLOR_RED), "[x] IP address not associated with an interface", lb_ip, string(COLOR_RESET))
		}
		log.Printf("%s [i] Load balancer %d: %s, contention ratio: %d %s\n", string(COLOR_GREEN), idx + 1, lb_ip, cont_ratio, string(COLOR_RESET))
		lb_list[idx] = load_balancer {address: Sprintf("%s:%d", lb_ip, lb_port), iface: iface, contention_ratio: cont_ratio, current_connections: 0}
	}
}
//
func handle_network(local_connection Conn, pipe_size int, try_count int, tunnel bool, no_delay bool, keep_alive bool, serial bool) {
	if tunnel == false {
		if address, err := handle_socks(local_connection); err == nil {
			go handle_internet(local_connection, address, pipe_size, try_count, no_delay, keep_alive, serial)
		}
	}
	if tunnel == true {
		go handle_tunnel(local_connection, pipe_size, try_count, no_delay, keep_alive, serial)
	}
}
//
func main() {
	processor_thread := runtime.NumCPU()
	runtime.GOMAXPROCS(processor_thread)
	var host = String("host", "::1", "The host to listen for SOCKS connection")
	var port = Int("port", 1080, "The port to listen for SOCKS connection")
	var size = Int("size", 4096, "The size of buffers in bytes for more power (default 4096)")
	var try = Int("try", 0, "The number of retries for SOCKS connection (default 0)")
	var tunnel = Bool("tunnel", false, "Use tunneling mode (acts as a transparent load balancing proxy)")
	var delay = Bool("delay", false, "Use delay mode (acts a combining a number of small outgoing messages and sending them all at once)")
	var alive = Bool("alive", false, "Use alive mode (sets whether the operating system should send keep-alive messages on the connection)")
	var serial = Bool("serial", false, "Use serial mode (acts to serialize access to function get load balancer)")
	var list = Bool("list", false, "Shows the available addresses for dispatching (non-tunneling mode only)")
	Parse()
	if runtime.GOOS == "windows" {
		exec.Command("powershell.exe", `(cls)`).Run()
	}
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
	parse_load_balancers(Args(), *tunnel, *list)
	execute_command()
	log.Println(string(COLOR_GREEN), "[i] Local server started on", host_port, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Pipe size is", *size, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Try count is", *try, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Tunnel is", *tunnel, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Protocol delay is", *delay, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Keep alive is", *alive, string(COLOR_RESET))
	log.Println(string(COLOR_GREEN), "[i] Serialize is", *serial, string(COLOR_RESET))
	if *delay == false {
		*delay = true
	}
	if *serial == true {
		mutex = &RWMutex {}
	}
	for {
		local_connection, _ := local_host.Accept()
		go handle_network(local_connection, *size, *try, *tunnel, *delay, *alive, *serial)
		if *serial == false {
			continue
		}
	}
}
