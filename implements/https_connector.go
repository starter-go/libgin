package implements

import (
	"net/http"
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
)

// HTTPSConnector 是默认的 HTTPS 连接器
type HTTPSConnector struct {
	//starter:component
	_as func(libgin.ConnectorRegistry) //starter:as(".")

	Context  libgin.Context //starter:inject("#")
	Enabled  bool           //starter:inject("${server.https.enabled}")
	Host     string         //starter:inject("${server.https.host}")
	Port     int            //starter:inject("${server.https.port}")
	KeyFile  string         //starter:inject("${server.https.key-file}")
	CertFile string         //starter:inject("${server.https.certificate-file}")

	connector connector
}

func (inst *HTTPSConnector) _impl() libgin.ConnectorRegistry {
	return inst
}

// ListRegistrations ...
func (inst *HTTPSConnector) ListRegistrations() []*libgin.ConnectorRegistration {
	r1 := &libgin.ConnectorRegistration{
		Name:      "https",
		Connector: inst,
	}
	return []*libgin.ConnectorRegistration{r1}
}

// Life ...
func (inst *HTTPSConnector) Life() *application.Life {
	return &application.Life{
		OnCreate:   inst.init,
		OnStart:    inst.start,
		OnLoop:     inst.loop,
		OnStop:     inst.stop,
		OnStopPost: inst.stopPost,
	}
}

// SetHandler ...
func (inst *HTTPSConnector) SetHandler(h http.Handler) {
	inst.connector.handler = h
}

func (inst *HTTPSConnector) init() error {
	router, err := inst.Context.GetRouterByName(theNameDefault)
	if err != nil {
		return err
	}
	inst.connector.enabled = inst.Enabled
	inst.connector.useHTTPS = true
	inst.connector.keyFile = inst.KeyFile
	inst.connector.certFile = inst.CertFile
	inst.connector.handler = router.Engine()
	return inst.connector.setAddress(inst.Host, inst.Port)
}

func (inst *HTTPSConnector) start() error {
	return inst.connector.start()
}

func (inst *HTTPSConnector) stop() error {
	return inst.connector.stop()
}

func (inst *HTTPSConnector) stopPost() error {
	timeout := time.Second * 15
	return inst.connector.waitForStopped(timeout)
}

func (inst *HTTPSConnector) loop() error {
	inst.connector.waitForStopping()
	return nil
}
