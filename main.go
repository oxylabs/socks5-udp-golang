package main

import (
	"flag"
	"log"

	"github.com/txthinking/socks5"
)

const defaultProxyURL = "socks.pr.oxylabs.io:7777"
const defaultProxyUsername = "customer-USERNAME"
const defaultProxyPassword = "PASSWORD"
const defaultTarget = "echo-udp.oxylabs.io:42000"
const defaultContent = "Hello UDP!"

func SOCKS5UDPRequest(proxyAddress, proxyUsername, proxyPassword string, target string, payload string) (string, error) {
	client, err := socks5.NewClient(proxyAddress, proxyUsername, proxyPassword, 10, 10)
	if err != nil {
		return "", err
	}

	conn, err := client.Dial("udp", target)
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(payload))
	if err != nil {
		return "", err
	}

	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func main() {
	proxyUrl := flag.String("proxy", defaultProxyURL, "proxy URL (hostname:port")
	proxyUsername := flag.String("u", defaultProxyUsername, "proxy username")
	proxyPassword := flag.String("p", defaultProxyPassword, "proxy password")
	target := flag.String("t", defaultTarget, "target (hostname:port)")
	content := flag.String("c", defaultContent, "content to send to target")
	flag.Parse()

	res, err := SOCKS5UDPRequest(*proxyUrl, *proxyUsername, *proxyPassword, *target, *content)
	if err != nil {
		log.Fatalln("Request failed due to: ", err)
	}
	log.Println("Response: ", res)
}
