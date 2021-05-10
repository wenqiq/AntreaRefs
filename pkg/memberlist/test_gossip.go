package memberlist

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"k8s.io/klog"

	"github.com/hashicorp/memberlist"
)

var (
	bindPort = flag.Int("port", 8002, "gossip port")
)

func Run() {
	flag.Parse()
	hostname, _ := os.Hostname()

	config := memberlist.DefaultLocalConfig()
	config.Name = hostname + "-" + strconv.Itoa(*bindPort)

	config.BindPort = *bindPort
	config.AdvertisePort = *bindPort

	klog.Infof("Configs: %+v\n", config)

	/* Create the initial memberlist from a safe configuration.
	   Please reference the godoc for other default config types.
	   http://godoc.org/github.com/hashicorp/memberlist#Config
	*/
	list, err := memberlist.Create(config)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	// Join an existing cluster by specifying at least one known member.
	n, err := list.Join([]string{"127.0.0.1:8002"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	klog.Infof("join cluster: %v\n", n)

	// Ask for members of the cluster
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}

	// Continue doing whatever you need, memberlist will maintain membership
	// information in the background. Delegates can be used for receiving
	// events when members join or leave.
	for {
		for _, member := range list.Members() {
			fmt.Printf("Member: %s %s, %#v\n", member.Name, member.Addr, member.State)
		}
		time.Sleep(time.Second * 3)
	}
}
