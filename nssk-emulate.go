package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/zhengxiaoyao0716/nssk-emulate/server"
	"github.com/zhengxiaoyao0716/util/bat"
	"github.com/zhengxiaoyao0716/util/requests"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	signal := flag.String("s", "run", "信号 run | scan")
	host := flag.String("host", "localhost", "指定主机IP或域名")
	port := flag.Int("port", -1, "指定端口号（-1表示自动检索）")
	master := flag.String("master", "http://localhost:5000", "中央服务主机通讯地址")
	flag.Parse()

	switch *signal {
	case "scan":
		log.Println("正在扫描本机可用IPv4地址")
		scanIPv4()
		log.Println("完成")
		os.Exit(0)
	case "run":
		address := "http://" + *host + ":"
		if *port == -1 {
			*port = 5000
			for true {
				if checkAddress(address + strconv.Itoa(*port)) {
					break
				}
				*port++
			}
		}
		address += strconv.Itoa(*port)
		if !checkAddress(address) {
			log.Println("检查到连接不可用")
			os.Exit(0)
		}
		if !run(address, *master) {
			log.Println("初始化失败")
			os.Exit(0)
		}
		server.Run(*host, *port)
	default:
		flag.PrintDefaults()
	}
}
func scanIPv4() {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln(err)
	}

	for _, address := range addresses {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println(ipnet.IP.String())
			}
		}
	}
}
func checkAddress(address string) bool {
	_, err := http.Get(address)
	if err == nil || !strings.Contains(err.Error(), "refused") {
		return false
	}
	return true
}
func run(address string, master string) bool {
	if master != address {
		server.PushCache("master", master)
		resp, err := requests.Post(
			master+"/api/s/user/join",
			map[string]interface{}{"address": address},
		)
		if err != nil {
			return false
		}
		secret := resp.JSON()["body"].(string)
		server.PushCache("secret", secret)
	}
	server.PushCache("address", address)

	if err := bat.Exec("start " + address + "/view"); err != nil {
		log.Println("* Start browser failed:", err)
		log.Println("* Please open this page in your browser manually:")
		log.Println("* " + address)
	}
	return true
}
