package grace

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server interface {
	Shutdown(ctx context.Context) error
}

type keeper struct {
	name    string
	server  Server
	timeout time.Duration
}

type Manager struct {
	servers []*keeper

	// for testing purposes
	mu           sync.Mutex
	manualSignal chan interface{}
}

type Opt func(*Manager)

func New(opts ...Opt) *Manager {
	f := &Manager{}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func WithServer(name string, server Server, timeout time.Duration) Opt {
	return func(f *Manager) {
		f.servers = append(f.servers, &keeper{
			name:    name,
			server:  server,
			timeout: timeout,
		})
	}
}

func (m *Manager) Wait() {
	stop := make(chan os.Signal, 2)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
	case <-m.ManualSignal():
	}

	for _, keeper := range m.servers {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), keeper.timeout)
			defer cancel()

			_ = keeper.server.Shutdown(ctx)
		}()
	}
}

func (m *Manager) ManualSignal() chan interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.manualSignal == nil {
		m.manualSignal = make(chan interface{}, 1)
	}

	return m.manualSignal
}
