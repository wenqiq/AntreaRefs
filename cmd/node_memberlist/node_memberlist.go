package main

import (
	"github.com/wenqiq/AntreaRefs/pkg/memberlist"
	"k8s.io/klog"
	"os"
	"os/signal"
	"syscall"
)

var (
	capturedSignals = []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT}
)

func main() {
	notifyCh := make(chan os.Signal, 2)
	stopCh := make(chan struct{})

	go func() {
		<-notifyCh
		close(stopCh)
		<-notifyCh
		klog.Warning("Received second signal, will force exit")
		klog.Flush()
		os.Exit(1)
	}()

	signal.Notify(notifyCh, capturedSignals...)
	memberlist.Run(stopCh)
}
