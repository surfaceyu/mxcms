package discovery

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"log"
)

type Address struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var urls []string

func GetEtcdUrls() []string {
	addr := make([]Address, 0)
	if err := config.Get("etcd").Scan(&addr); err != nil {
		log.Panic(err)
	}

	for _, addr := range addr {
		urls = append(urls, fmt.Sprintf("%s:%d", addr.Host, addr.Port))
	}

	return urls
}

