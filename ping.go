package main

import (
	"log"
	"net"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

var (
	ipchnl = make(chan string)
	ipList = make([]string, 0)
)

func Ping(wg *sync.WaitGroup, ip string) {
	defer wg.Done()
	resp := exec.Command("ping", "-b", "-c1", "-W1", ip)
	_, err := resp.Output()
	if nil == err {
		//log.Println("IP up :", ip)
		ipchnl <- ip
	}
	return
}

func GenerateIP() {
	var wg sync.WaitGroup
	go chnlTrack(&wg)
	wg.Add(1)
	base := "192.168."
	var ip string
	for i := 43; i <= 255; i++ {
		for j := 100; j <= 255; j++ {
			ip = base + strconv.Itoa(i) + "." + strconv.Itoa(j)
			//log.Println("Generated Ip := ", ip)
			go Ping(&wg, ip)
			wg.Add(1)
		}
	}
	wg.Wait()
	log.Println("Final Ip :", ipList)
}

func chnlTrack(wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case ip := <-ipchnl:
		ipList = append(ipList, ip)
		if len(ipList) == 2 {
			close(ipchnl)
			return
		}
		log.Println("Final", ipList)
	case <-time.After(time.Minute * 1):
		log.Println("Total length is :", len(ipList))
		log.Fatal("Done")
		return
	}
}

func main() {
	GenerateIP()
}

func GetLocalIpAdress() {
	ifcases, err := net.Interfaces()
	if err != nil {
		log.Println(err)
	}
	for _, i := range ifcases {
		addrs, err := i.Addrs()
		if err != nil {
			log.Panic(err)
		}
		for _, add := range addrs {
			switch v := add.(type) {
			case *net.IPNet:
				log.Println("first : ", v.IP)
			case *net.IPAddr:
				log.Println("Second ", v.IP)
			}
		}
	}
}
