package implements

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/starter-go/vlog"
)

type connector struct {
	address  string // like 'host:port'
	handler  http.Handler
	starting bool
	stopping bool
	stopped  bool
	started  bool
	useHTTPS bool
	enabled  bool
	keyFile  string
	certFile string
}

func (inst *connector) setAddress(host string, port int) error {
	if port < 1 {
		if inst.useHTTPS {
			port = 8443
		} else {
			port = 8080
		}
	}
	inst.address = fmt.Sprintf("%s:%d", host, port)
	return nil
}

func (inst *connector) handleError(err error) {
	if err == nil {
		return
	}
	vlog.Error("%v", err)
}

func (inst *connector) start() error {
	if inst.handler == nil {
		return fmt.Errorf("libgin.connector.handler is nil")
	}
	if inst.address == "" {
		return fmt.Errorf("bad libgin.connector.address: '%s'", inst.address)
	}
	inst.starting = true
	go func() {
		err := inst.run()
		inst.handleError(err)
	}()
	return nil
}

func (inst *connector) stop() error {
	inst.stopping = true
	return nil
}

func (inst *connector) run() error {

	if !inst.enabled {
		return nil
	}

	server := &http.Server{
		Addr:    inst.address,
		Handler: inst.handler,
	}

	if inst.useHTTPS {
		inst.serveAsHTTPS(server)
	} else {
		inst.serveAsHTTP(server)
	}

	inst.started = true
	inst.waitForStopping()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		vlog.Fatal("Server forced to shutdown: ", err)
		return err
	}

	return nil
}

func (inst *connector) serveAsHTTP(server *http.Server) {
	go func() {
		vlog.Info("serve HTTP at %s", inst.address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			vlog.Error("listen: %s\n", err)
			inst.stopping = true
		}
	}()
}

func (inst *connector) serveAsHTTPS(server *http.Server) {
	cert := inst.certFile
	key := inst.keyFile
	go func() {
		vlog.Info("serve HTTPS at %s", inst.address)
		if err := server.ListenAndServeTLS(cert, key); err != nil && err != http.ErrServerClosed {
			vlog.Error("listen: %s\n", err)
			inst.stopping = true
		}
	}()
}

func (inst *connector) waitForStopping() {
	if !inst.enabled {
		return
	}
	for {
		if inst.stopping {
			break
		}
		time.Sleep(time.Second * 3)
	}
}

func (inst *connector) waitForStopped(timeout time.Duration) error {
	if !inst.enabled {
		return nil
	}
	ttl := timeout
	step := time.Millisecond * 300
	for ; ttl > 0; ttl -= step {
		if inst.stopped {
			return nil
		}
		time.Sleep(step)
	}
	return fmt.Errorf("timeout: in waitForStopped()")
}
