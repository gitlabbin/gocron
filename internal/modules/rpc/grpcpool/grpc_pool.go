package grpcpool

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/ouqiang/gocron/internal/modules/app"
	"github.com/ouqiang/gocron/internal/modules/rpc/auth"
	"github.com/ouqiang/gocron/internal/modules/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	backOffMaxDelay = 3 * time.Second
	dialTimeout     = 2 * time.Second
)

var (
	Pool = &GRPCPool{
		conns: make(map[string]*entry),
	}

	keepAliveParams = keepalive.ClientParameters{
		Time:                20 * time.Second,
		Timeout:             3 * time.Second,
		PermitWithoutStream: true,
	}
)

type Client struct {
	conn      *grpc.ClientConn
	rpcClient rpc.TaskClient
}

type GRPCPool struct {
	// map key格式 ip:port
	conns map[string]*entry
	mu    sync.Mutex
}

type entry struct {
	res   *Client
	ready chan struct{} // closed when res is ready
}

func (p *GRPCPool) Get(addr string) (rpc.TaskClient, error) {
	var err error
	p.mu.Lock()
	e, _ := p.conns[addr]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		p.conns[addr] = e
		p.mu.Unlock()
		e.res, err = p.factory(addr)
		close(e.ready) // broadcast ready condition
	} else {
		p.mu.Unlock()
		<-e.ready // wait for ready condition
	}

	if err != nil {
		return nil, err
	}
	return e.res.rpcClient, nil
}

// 释放连接
func (p *GRPCPool) Release(addr string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	client, ok := p.conns[addr]
	if !ok {
		return
	}
	delete(p.conns, addr)
	client.res.conn.Close()
}

// 创建连接
func (p *GRPCPool) factory(addr string) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithKeepaliveParams(keepAliveParams),
		grpc.WithBackoffMaxDelay(backOffMaxDelay),
	}

	if !app.Setting.EnableTLS {
		opts = append(opts, grpc.WithInsecure())
	} else {
		server := strings.Split(addr, ":")
		certificate := auth.Certificate{
			CAFile:     app.Setting.CAFile,
			CertFile:   app.Setting.CertFile,
			KeyFile:    app.Setting.KeyFile,
			ServerName: server[0],
		}

		transportCreds, err := certificate.GetTransportCredsForClient()
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(transportCreds))
	}

	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:      conn,
		rpcClient: rpc.NewTaskClient(conn),
	}

	return client, nil
}
