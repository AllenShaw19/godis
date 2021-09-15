package echo

import (
	"bufio"
	"context"
	"errors"
	"godis/log"
	"io"
	"net"
	"sync"
	"sync/atomic"
)

type Client struct {
	Conn net.Conn
}

func (c *Client) Close() error {
	_ = c.Conn.Close()
	return nil
}

type Handler struct {
	activeConns sync.Map
	closing     atomic.Value
}

func NewHandler() *Handler {
	h := &Handler{}
	h.closing.Store(false)
	return h
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing.Load().(bool) {
		_ = conn.Close()
	}

	client := &Client{Conn: conn}
	h.activeConns.Store(client, struct{}{})

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Info("connection close")
				h.activeConns.Delete(conn)
			} else {
				log.Error("read string from conn err: %v", err)
			}
			return
		}
		conn.Write([]byte(msg))
	}
}

func (h *Handler) Close() error {
	log.Info("handler shutting down...")
	h.closing.Store(true)
	h.activeConns.Range(func(key, value interface{}) bool {
		client := key.(*Client)
		client.Close()
		return true
	})
	return nil
}
