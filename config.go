package go_rbac

type ConfigMgr struct {
	Addr string
}

const (
	webAddr = "rbac-web-svc:6802"
)

func NewDefaultConfigMgr() *ConfigMgr {
	return &ConfigMgr{Addr: webAddr}
}
