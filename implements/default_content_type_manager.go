package implements

import (
	"fmt"
	"strings"
	"sync"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
)

// DefaultContentTypeManager ...
type DefaultContentTypeManager struct {

	//starter:component
	_as func(libgin.ContentTypeManager) //starter:as("#")

	All []libgin.ContentTypeRegistry //starter:inject(".")

	mutex sync.Mutex
	table map[string]*libgin.ContentTypeRegistration
}

func (inst *DefaultContentTypeManager) _impl() (libgin.ContentTypeManager, application.Lifecycle) {
	return inst, inst
}

// Life ...
func (inst *DefaultContentTypeManager) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *DefaultContentTypeManager) init() error {
	inst.FindTypeBySuffix(".html")
	return nil
}

func (inst *DefaultContentTypeManager) keyForSuffix(suffix string) string {
	const dot = "."
	suffix = strings.TrimSpace(suffix)
	suffix = strings.ToLower(suffix)
	if !strings.HasPrefix(suffix, dot) {
		suffix = dot + suffix
	}
	return suffix
}

// FindTypeBySuffix ...
func (inst *DefaultContentTypeManager) FindTypeBySuffix(suffix string) (string, error) {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	key := inst.keyForSuffix(suffix)
	table := inst.getTable()
	item := table[key]

	if item == nil {
		return "", fmt.Errorf("no type info for suffix:'%s'", suffix)
	}

	return item.ContentType, nil
}

func (inst *DefaultContentTypeManager) getTable() map[string]*libgin.ContentTypeRegistration {
	t := inst.table
	if t == nil {
		t = inst.loadTable()
		inst.table = t
	}
	return t
}

func (inst *DefaultContentTypeManager) loadTable() map[string]*libgin.ContentTypeRegistration {
	src := inst.All
	dst := make(map[string]*libgin.ContentTypeRegistration)
	for _, r1 := range src {
		r2list := r1.ListRegistrations()
		for _, r2 := range r2list {
			for _, suffix := range r2.Suffixes {
				key := inst.keyForSuffix(suffix)
				dst[key] = r2
			}
		}
	}
	return dst
}
