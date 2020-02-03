package balance

import "fmt"

type BalanceMgr struct {
	allBalance map[string]Balance
}

var mgr = BalanceMgr{
	allBalance: make(map[string]Balance),
}

func (p *BalanceMgr) registerBalance(name string, b Balance) {
	p.allBalance[name] = b
}

func RegisterBalance(name string, b Balance) {
	mgr.registerBalance(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balance, ok := mgr.allBalance[name]
	if !ok {
		err = fmt.Errorf("not fount %s", name)
		fmt.Println("not found ", name)
		return
	}
	inst, err = balance.DoBalance(insts)
	if err != nil {
		err = fmt.Errorf(" %s erros", name)
		return
	}
	return
}

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (p *Instance) GetHost() string {
	return p.host
}

func (p *Instance) GetPort() int {
	return p.port
}

type Balance interface {
	/**
	 *负载均衡算法
	 */
	DoBalance([]*Instance, ...string) (*Instance, error)
}
