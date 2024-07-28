package implements

import (
	"fmt"
	"strings"
	"sync"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/mimetypes"
)

// DefaultContentTypeManager ...
type DefaultContentTypeManager struct {

	//starter:component

	_as func(libgin.ContentTypeManager) //starter:as("#")

	Source mimetypes.Manager //starter:inject("#")

	mutex sync.Mutex
	table map[string]*mimetypes.Info
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

	// get in cache
	key := inst.keyForSuffix(suffix)
	table := inst.getTable()
	item := table[key]
	if item != nil {
		return item.Type.String(), nil
	}

	// load from source
	info, err := inst.Source.FindBySuffix(suffix, nil)
	if err == nil && info != nil {
		table[key] = info
		return info.Type.String(), nil
	}

	// not found
	return "", fmt.Errorf("no type info for suffix:'%s'", suffix)
}

func (inst *DefaultContentTypeManager) getTable() map[string]*mimetypes.Info {
	t := inst.table
	if t == nil {
		t = make(map[string]*mimetypes.Info)
		inst.table = t
	}
	return t
}
