package implements

import (
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/mimetypes"
)

// ContentTypeResourceLoader 负责从指定的资源文件中加载类型信息
type ContentTypeResourceLoader struct {

	//starter:component

	_as func(mimetypes.Registry) //starter:as(".")

	Context    application.Context //starter:inject("context")
	TypesProps string              //starter:inject("${web.content-types.properties}")

	cached []*mimetypes.Registration
}

func (inst *ContentTypeResourceLoader) _impl() mimetypes.Registry {
	return inst
}

// ListRegistrations ...
func (inst *ContentTypeResourceLoader) ListRegistrations() []*mimetypes.Registration {
	list := inst.cached
	if list == nil {
		list = inst.load()
		inst.cached = list
	}
	return list
}

func (inst *ContentTypeResourceLoader) load() []*mimetypes.Registration {

	dst := make([]*mimetypes.Registration, 0)

	res, err := inst.Context.GetResources().GetResource(inst.TypesProps)
	if err != nil {
		return dst
	}

	text, err := res.ReadText()
	if err != nil {
		return dst
	}

	tab, err := properties.Parse(text, safe.Fast())
	if err != nil {
		return dst
	}

	kv := tab.Export(nil)
	for k, v := range kv {
		r, err := inst.parseItem(k, v)
		if err == nil {
			dst = append(dst, r)
		}
	}

	return dst
}

func (inst *ContentTypeResourceLoader) parseItem(k, v string) (*mimetypes.Registration, error) {

	const dot = "."
	item := &mimetypes.Registration{}
	namelist := strings.Split(v, ",")

	// item.ContentType = strings.TrimSpace(k)
	tpName := strings.TrimSpace(k)
	item.Name = tpName
	item.Info.Type = mimetypes.Type(tpName)
	item.Info.Label = tpName

	for _, name := range namelist {
		name = strings.TrimSpace(name)
		name = strings.ToLower(name)
		if !strings.HasPrefix(name, dot) {
			name = dot + name
		}
		item.Suffixes = append(item.Suffixes, name)
	}

	return item, nil
}
