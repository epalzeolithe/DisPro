// Made by SirSAC for Network.
package main
//
import (
	. "encoding/binary"
	. "errors"
	. "fmt"
	"log"
	. "net"
)
//
func server_choice(local_connection Conn) (error) {
	if nWrite, err := local_connection.Write([]byte {SOCKS_VERSION, NO_AUTHENTICATION}); err != nil || nWrite != 2 {
		return New("[!] Servers choice failed")
	}
	return nil
}
//
func client_greeting(local_connection Conn) (byte, []byte, error) {
	buf := make([]byte, 2)
	if nRead, err := local_connection.Read(buf); err != nil || nRead != len(buf) {
		return 0, nil, New("[!] Client greeting failed")
	}
	nauth := buf[1]
	auth := make([]byte, nauth)
	if nRead, err := local_connection.Read(auth); err != nil || nRead != int(nauth) {
		return 0, nil, New("[!] Client greeting failed")
	}
	ver := buf[0]
	return ver, auth, nil
}
//
func client_request(local_connection Conn) (string, error) {
	header := make([]byte, 4)
	if nRead, err := local_connection.Read(header); err != nil || nRead != len(header) {
		local_connection.Write([]byte {5, GENERAL_FAILURE, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		return "", New("[!] Client connection request failed")
	}
	ver := header[0]
	if ver != SOCKS_VERSION {
		local_connection.Write([]byte {5, CONNECTION_NOT_ALLOWED, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		return "", New("[!] Unsupported SOCKS version")
	}
	cmd := header[1]
	if cmd != STREAM_CONNECTION {
		if cmd != PORT_BINDING {
			if cmd != ASSOCIATE_PORT {
				local_connection.Write([]byte {5, COMMAND_NOT_SUPPORTED, 0, 1, 0, 0, 0, 0, 0, 0})
				local_connection.Close()
				return "", New("[!] Unsupported command code")
			}
		}
	}
	dstaddr := header[3]
	dstport := make([]byte, 2)
	var address string
	switch dstaddr {
		case DOMAIN_NAME:
			domain_name_length := make([]byte, 1)
			if nRead, err := local_connection.Read(domain_name_length); err != nil || nRead != len(domain_name_length) {
				local_connection.Write([]byte {5, CONNECTION_REFUSED, 0, 1, 0, 0, 0, 0, 0, 0})
				local_connection.Close()
				return "", New("[!] Client connection request failed")
			}
			domain_name := make([]byte, domain_name_length[0])
			if nRead, err := local_connection.Read(domain_name); err != nil || nRead != len(domain_name) {
				local_connection.Write([]byte {5, GENERAL_FAILURE, 0, 1, 0, 0, 0, 0, 0, 0})
				local_connection.Close()
				return "", New("[!] Client connection request failed")
			}
			if nRead, err := local_connection.Read(dstport); err != nil || nRead != len(dstport) {
				local_connection.Write([]byte {5, GENERAL_FAILURE, 0, 1, 0, 0, 0, 0, 0, 0})
				local_connection.Close()
				return "", New("[!] Client connection request failed")
			}
			address = Sprintf("%s:%d", string(domain_name), BigEndian.Uint16(dstport))
		case IPV6_ADDRESS:
			ipv6_address := make([]byte, 16)
			if nRead, err := local_connection.Read(ipv6_address); err != nil || nRead != len(ipv6_address) {
				local_connection.Write([]byte {5, CONNECTION_REFUSED, 0, 1, 0, 0, 0, 0, 0, 0})
				local_connection.Close()
				return "", New("[!] Client connection request failed")
			}
			if nRead, err := local_connection.Read(dstport); err != nil || nRead != len(dstport) {
				local_connection.Write([]byte {5, GENERAL_FAILURE, 0, 1, 0, 0, 0, 0, 0, 0})
				local_connection.Close()
				return "", New("[!] Client connection request failed")
			}
			address = Sprintf("[%d:%d:%d:%d:%d:%d:%d:%d]:%d", ipv6_address[0], ipv6_address[1], ipv6_address[2], ipv6_address[3], ipv6_address[4], ipv6_address[5], ipv6_address[6], ipv6_address[7], BigEndian.Uint16(dstport))
		case IPV4_ADDRESS:
			ipv4_address := make([]byte, 4)
			if nRead, err := local_connection.Read(ipv4_address); err != nil || nRead != len(ipv4_address) {
				local_connection.Write([]byte {5, CONNECTION_REFUSED, 0, 1, 0, 0, 0, 0, 0, 0})
				local_connection.Close()
				return "", New("[!] Client connection request failed")
			}
			if nRead, err := local_connection.Read(dstport); err != nil || nRead != len(dstport) {
				local_connection.Write([]byte {5, GENERAL_FAILURE, 0, 1, 0, 0, 0, 0, 0, 0})
				local_connection.Close()
				return "", New("[!] Client connection request failed")
			}
			address = Sprintf("%d.%d.%d.%d:%d", ipv4_address[0], ipv4_address[1], ipv4_address[2], ipv4_address[3], BigEndian.Uint16(dstport))
		default:
			local_connection.Write([]byte {5, ADDRESS_TYPE_NOT_SUPPORTED, 0, 1, 0, 0, 0, 0, 0, 0})
			local_connection.Close()
			return "", New("[!] Unsupported address type")
	}
	return address, nil
}
//
func handle_socks(local_connection Conn) (string, error) {
	if err := server_choice(local_connection); err != nil {
		log.Println(string(COLOR_YELLOW), err, string(COLOR_RESET))
		return "", err
	}
	if _, _, err := client_greeting(local_connection); err != nil {
		log.Println(string(COLOR_YELLOW), err, string(COLOR_RESET))
		return "", err
	}
	address, err := client_request(local_connection)
	if err != nil {
		log.Println(string(COLOR_YELLOW), err, string(COLOR_RESET))
		return "", err
	}
	return address, nil
}
