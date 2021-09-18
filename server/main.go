package main

import (
	"context"
	"godis/echo"
	"godis/log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}

func ListenAndServe(listener net.Listener, handler Handler, closeChan <-chan struct{}) {
	// 监听关闭channel
	go func() {
		<-closeChan // TODO:如果需要不同类型的通知。复用同一个chan,根据接收到的值进行switch
		log.Info("shutting down...")
		_ = listener.Close()
		_ = handler.Close()
	}()

	defer func() {
		_ = listener.Close()
		_ = handler.Close()
	}()

	ctx := context.Background()
	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("accept err: %v", err)
			break
		}
		log.Info("accept client")
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			handler.Handle(ctx, conn)
		}()
	}
	wg.Wait()
}

func ListenAndServeWithSignal(address string, handler Handler) {
	closeChan := make(chan struct{})
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigChan
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("recv signal %v, close server", sig)
			closeChan <- struct{}{}
		}
	}()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("listen err: %v", err)
	}
	defer listener.Close()

	log.Info("bind: %s, start listening...", address)
	ListenAndServe(listener, handler, closeChan)
}

func main() {
	ListenAndServeWithSignal(":8088", echo.NewHandler())
}
