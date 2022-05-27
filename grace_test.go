package grace

import (
	"context"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

type testServer struct {
	shutdown bool
	wait     time.Duration
}

func (t *testServer) Shutdown(ctx context.Context) error {
	wait := time.NewTimer(t.wait)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-wait.C:
		// server finished fake busy work
	}

	t.shutdown = true

	return nil
}

func TestManager_Wait(t *testing.T) {
	srv := &testServer{wait: time.Second}

	m := New(
		WithServer("", srv, time.Second*10),
	)

	go func() {
		m.ManualSignal() <- nil
	}()

	assert.Assert(t, !srv.shutdown, "expected server not to be shutdown yet")

	m.Wait()

	assert.Assert(t, srv.shutdown, "expected server to be shutdown")
}

func TestManager_ManualSignal(t *testing.T) {
	m := New()
	s1 := m.ManualSignal()
	s2 := m.ManualSignal()

	assert.Assert(t, s1 == s2, "expected signal channels to be the same")
}

func TestWithServer(t *testing.T) {
	srv := &testServer{wait: time.Second}

	m := New(
		WithServer("cooper", srv, time.Second*10),
	)

	assert.Assert(t, len(m.servers) == 1, "expected 1 server")
	assert.Assert(t, m.servers[0].name == "cooper", "expected default server name")
	assert.Assert(t, m.servers[0].server == srv, "expected default server")
	assert.Assert(t, m.servers[0].timeout == time.Second*10, "expected default server timeout")
}

func TestNew(t *testing.T) {
	m := New()

	assert.Assert(t, len(m.servers) == 0, "expected 0 servers")
}
