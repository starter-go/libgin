package implements

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/application/resources"
	"github.com/starter-go/libgin"
)

// StaticController 是静态资源控制器，它负责把静态的web资源注册到 static-group
type StaticController struct {
	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Context    application.Context       //starter:inject("context")
	ResPath    string                    //starter:inject("${web-group.static.resources}")
	IndexNames string                    //starter:inject("${web-group.static.index-names}")
	Types      libgin.ContentTypeManager //starter:inject("#")

	indexNameList []string // cached for IndexNames
}

func (inst *StaticController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *StaticController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{libgin.GroupStatic},
		Route:  inst.r,
	}
}

func (inst *StaticController) r(g *gin.RouterGroup) error {
	reslist := inst.listStaticResources()
	for _, h := range reslist {
		inst.makeHandler(g, h)
	}
	return nil
}

func (inst *StaticController) makeHandler(g *gin.RouterGroup, h *myStaticResHolder) {
	path := h.webPath
	path2 := ""
	path3 := ""

	if inst.isIndexName(h.simpleName) {
		i := strings.LastIndex(path, "/")
		if i >= 0 {
			path2 = path[0:i]
			path3 = path[0 : i+1]
		}
	}

	g.GET(path, func(c *gin.Context) {
		data := h.getData()
		c.Data(http.StatusOK, h.contentType, data)
	})

	if path3 != "" {
		g.GET(path3, func(c *gin.Context) {
			data := h.getData()
			c.Data(http.StatusOK, h.contentType, data)
		})
	}

	if path2 != "" {
		g.GET(path2, func(c *gin.Context) {
			data := "redirect to " + path3
			c.Header("Location", path3)
			c.Data(http.StatusTemporaryRedirect, "text/plain", []byte(data))
		})
	}
}

func normalizePath(path string) string {
	const (
		sep1 = "\\"
		sep2 = "/"
	)
	builder := &strings.Builder{}
	path = strings.ReplaceAll(path, sep1, sep2)
	elements := strings.Split(path, sep2)
	for _, el := range elements {
		el = strings.TrimSpace(el)
		if el == "" {
			continue
		}
		builder.WriteString("/")
		builder.WriteString(el)
	}
	return builder.String()
}

func (inst *StaticController) listStaticResources() []*myStaticResHolder {
	prefix := normalizePath(inst.ResPath) + "/"
	all := inst.Context.GetResources().Export(nil)
	dst := make([]*myStaticResHolder, 0)
	for _, res := range all {
		h := &myStaticResHolder{}
		h.init(inst, res, prefix)
		if h.isStaticRes {
			dst = append(dst, h)
		}
	}
	return dst
}

func (inst *StaticController) loadIndexNameList() []string {
	src := strings.Split(inst.IndexNames, ",")
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

func (inst *StaticController) getIndexNameList() []string {
	list := inst.indexNameList
	if list == nil {
		list = inst.loadIndexNameList()
		inst.indexNameList = list
	}
	return list
}

func (inst *StaticController) isIndexName(name string) bool {
	all := inst.getIndexNameList()
	for _, index := range all {
		if index == name {
			return true
		}
	}
	return false
}

func (inst *StaticController) findType(suffix string) string {
	t, err := inst.Types.FindTypeBySuffix(suffix)
	if err == nil {
		return t
	}
	return "application/octet-stream"
}

////////////////////////////////////////////////////////////////////////////////

type myStaticResHolder struct {
	webPath  string
	resPath  string
	basePath string

	contentType string
	simpleName  string
	suffix      string

	isStaticRes bool
	res         resources.Resource
	cachedData  []byte
}

func (inst *myStaticResHolder) init(parent *StaticController, res resources.Resource, basepath string) {

	if !strings.HasSuffix(basepath, "/") {
		basepath = basepath + "/"
	}

	// res path
	respath := normalizePath(res.Path())

	// simple name
	simpleName := ""
	i := strings.LastIndex(respath, "/")
	if i >= 0 {
		simpleName = respath[i+1:]
	}

	// suffix
	suffix := ""
	i = strings.LastIndex(simpleName, ".")
	if i >= 0 {
		suffix = simpleName[i:]
	}

	// content type
	ctype := parent.findType(suffix)

	if strings.HasPrefix(respath, basepath) {
		inst.isStaticRes = true
		inst.resPath = respath
		inst.basePath = basepath
		inst.webPath = "/" + respath[len(basepath):]
		inst.simpleName = simpleName
		inst.suffix = suffix
		inst.contentType = ctype
		inst.res = res
	} else {
		inst.isStaticRes = false
	}
}

func (inst *myStaticResHolder) getData() []byte {
	data := inst.cachedData
	if data != nil {
		return data
	}
	content, err := inst.res.ReadBinary()
	if err == nil {
		data = content
	} else {
		data = make([]byte, 0)
	}
	inst.cachedData = data
	return data
}
