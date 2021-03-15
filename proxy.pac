// Made by SirSAC for Firefox.
const firefox_socks = "SOCKS [::1]:9050; SOCKS [::1]:9050";
const firefox_proxy = "PROXY [::1]:8118; PROXY [::1]:8118";
//
function FindProxyForURL (url, host) {
	if (shExpMatch (host, "*.onion")) {
		return firefox_socks;
	}
	return firefox_proxy;
}
