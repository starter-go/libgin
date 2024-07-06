package implements

import (
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/stopper"
)

// HTTPConnector 是默认的 HTTP 连接器
type HTTPConnector struct {

	//starter:component
	_as func(libgin.ConnectorRegistry) //starter:as(".")

	AppContext application.Context //starter:inject("context")
	Context    libgin.Context      //starter:inject("#")
	Stopper    stopper.Service     //starter:inject("#")

	Host    string //starter:inject("${server.http.host}")
	Port    int    //starter:inject("${server.http.port}")
	Enabled bool   //starter:inject("${server.http.enabled}")

	connector connector
}

func (inst *HTTPConnector) _impl() libgin.ConnectorRegistry {
	return inst
}

// ListRegistrations ...
func (inst *HTTPConnector) ListRegistrations() []*libgin.ConnectorRegistration {
	r1 := &libgin.ConnectorRegistration{
		Name:      "http",
		Connector: inst,
	}
	return []*libgin.ConnectorRegistration{r1}
}

// Life ...
func (inst *HTTPConnector) Life() *application.Life {

	action := stopper.GetAction(inst.AppContext)
	if action == stopper.ActionStop {
		return &application.Life{} // 如果正在执行停止动作， 则返回空的 life
	}

	return &application.Life{
		OnCreate:   inst.init,
		OnStart:    inst.start,
		OnLoop:     inst.loop,
		OnStop:     inst.stop,
		OnStopPost: inst.stopPost,
	}
}

func (inst *HTTPConnector) init() error {
	router, err := inst.Context.GetRouterByName(theNameDefault)
	if err != nil {
		return err
	}
	inst.connector.enabled = inst.Enabled
	inst.connector.handler = router.Engine()
	return inst.connector.setAddress(inst.Host, inst.Port)
}

func (inst *HTTPConnector) start() error {
	return inst.connector.start()
}

func (inst *HTTPConnector) stop() error {
	return inst.connector.stop()
}

func (inst *HTTPConnector) stopPost() error {
	timeout := time.Second * 15
	return inst.connector.waitForStopped(timeout)
}

func (inst *HTTPConnector) loop() error {
	// 这里不需要等待了，用 stopper 的 loop 代替
	// inst.connector.waitForStopping()
	return nil
}
