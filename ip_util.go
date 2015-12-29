package goutils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"regexp"
	"strings"
	"time"
)

func GetIpFromAddress(address string) string { // address="ip:port"
	list := strings.Split(address, ":")
	if len(list) == 2 {
		return list[0]
	}
	list = strings.Split(address, "/")
	if len(list) == 2 {
		return list[0]
	}
	return ""
}

// use command :nc ns1.dnspod.net 6666
func GetLocalPublicIpUseDnspod() (string, error) {
	timeout := 10 * time.Second
	conn, err := net.DialTimeout("tcp", "ns1.dnspod.net:6666", timeout)
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Can't get public ip", x)
		}
		if conn != nil {
			conn.Close()
		}
	}()
	if err == nil {
		var bytes []byte
		deadline := time.Now().Add(timeout)
		err = conn.SetDeadline(deadline)
		if err != nil {
			return "", err
		}
		bytes, err = ioutil.ReadAll(conn)
		if err == nil {
			return string(bytes), nil
		}
	}
	return "", err
}

var (
	NO_IP_MATCHE_ERROR = errors.New("no ip matched")
)

func GetLocalPublicIpUseIP138() (ip string, err error) {
	timeout := 10 * 1000 // 10 s
	var data []byte
	data, err = GetHttpResponseAsJson("http://www.ip138.com/ips1388.asp", time.Now(), timeout)
	if err != nil {
		return
	}
	// 您的IP地址是：[124.126.228.111]
	r := regexp.MustCompile(`\[(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\]`)
	results := r.FindStringSubmatch(string(data))
	if len(results) > 1 {
		ip = string(results[1])
		return
	}
	err = NO_IP_MATCHE_ERROR
	return
}

// get ip from https://cgi1.apnic.net/cgi-bin/my-ip.php
func GetLocalPublicIpUseApnic() (ip string, err error) {
	timeout := 10 * 1000 // 10 s
	var data []byte
	data, err = GetHttpResponseAsJson("https://cgi1.apnic.net/cgi-bin/my-ip.php", time.Now(), timeout)
	if err != nil {
		return
	}
	r := regexp.MustCompile(`ip:"(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})"`)
	results := r.FindStringSubmatch(string(data))
	if len(results) > 1 {
		ip = string(results[1])
		return
	}
	err = NO_IP_MATCHE_ERROR
	return
}

func GetLocalIP() (ip string, err error) {
	conn, err := net.Dial("udp", "www.ip138.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		if conn != nil {
			conn.Close()
			conn = nil
		}
	}()
	ip = strings.Split(conn.LocalAddr().String(), ":")[0]
	return
}
