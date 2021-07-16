package flag

import (
	"errors"
	"log"
)

func init() {
}

type callback func() error

type Module struct {
	Name     string   // 模块名
	DepNames []string // 依赖的模块
	Initor   func() error
	Fin      func() error

	has_init_flag bool // 已经调用了 Initor
	has_fin_flag  bool // 已经调用了 Fin
	Deps          []*Module
}

type ModuleMgr struct {
	module_list []*Module
	module_map  map[string]*Module
	fin_list    []*Module
	inited      bool
}

var g_module_mgr *ModuleMgr

func GetModuleMgr() *ModuleMgr {
	if g_module_mgr == nil {
		g_module_mgr = &ModuleMgr{
			module_list: make([]*Module, 0, 1000),
			module_map:  make(map[string]*Module),
			fin_list:    make([]*Module, 0, 1000),
		}
	}
	return g_module_mgr
}

func (self *ModuleMgr) CallInit(m *Module) error {
	if m.has_init_flag {
		return nil
	}
	var err error

	for _, o := range m.Deps {
		err = self.CallInit(o)
		if err != nil {
			log.Printf("init module deps failed, [module=%s] [err=%v]", m.Name, err)
			return err
		}
	}

	if m.Initor != nil {
		err = (m.Initor)()
		if err != nil {
			log.Printf("init module failed, [module=%s] [err=%v]", m.Name, err)
			return err
		}
	}

	self.fin_list = append(self.fin_list, m)

	log.Printf("init module success, [module=%s]", m.Name)
	m.has_init_flag = true

	return err
}

func (self *ModuleMgr) InitAllModule() error {
	var err error
	var cnt int
	for _, m := range self.module_list {
		if !m.has_init_flag {
			err = self.CallInit(m)
			if err != nil {
				return err
			}
			cnt += 1
		}
	}

	log.Printf("init all module success, [total=%d]", cnt)
	return nil
}

func CallFin(m *Module) error {
	if m.has_fin_flag {
		return nil
	}

	if m.Fin == nil {
		return nil
	}

	var err error
	err = (m.Fin)()
	if err != nil {
		log.Printf("fin [module=%s] failed, [err=%v]", m.Name, err)
		return err
	}

	log.Printf("fin [module=%s] success", m.Name)
	m.has_fin_flag = true
	return nil
}

func (self *ModuleMgr) FinAllModule() error {
	var err error
	for _, m := range self.fin_list {
		if m == nil {
			continue
		}
		if !m.has_fin_flag {
			err = CallFin(m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *ModuleMgr) Init() error {
	for _, m := range self.module_list {
		for _, o := range m.DepNames {
			if m.Name == o {
				log.Fatalf(" module depend self, [module=%s]", m.Name)
				return errors.New("module depend self")
			}
			other := self.module_map[o]
			m.Deps = append(m.Deps, other)
		}
	}
	return nil
}

func (self *ModuleMgr) Registry(name string) *Module {
	if _, ok := self.module_map[name]; ok {
		log.Fatalf("module duplicated registry! [module=%s]", name)
		return nil
	}

	var m = &Module{Name: name, DepNames: make([]string, 0, 10)}
	self.module_map[name] = m
	self.module_list = append(self.module_list, m)
	return m
}

func (self *ModuleMgr) Get(name string) *Module {
	m, ok := self.module_map[name]
	if !ok {
		return nil
	}
	return m
}

func InitAllModule() error {
	var err error
	mgr := GetModuleMgr()
	if mgr.inited {
		return nil
	}

	err = mgr.Init()
	if err != nil {
		log.Fatalf("module mgr init failed! [err=%v]", err)
		return err
	}

	err = mgr.InitAllModule()
	if err != nil {
		log.Fatalf("init all modules failed! [err=%v]", err)
		return err
	}

	var list = make([]*Module, len(mgr.fin_list))
	for i := len(mgr.fin_list) - 1; i >= 0; i-- {
		if mgr.fin_list[i] != nil {
			list = append(list, mgr.fin_list[i])
		}
	}
	mgr.fin_list = list
	mgr.inited = true
	return nil
}

func FinAllModule() error {
	var err error
	mgr := GetModuleMgr()
	err = mgr.FinAllModule()
	if err != nil {
		log.Fatalf("fin all modules failed! [err=%v]", err)
		return err
	}
	return nil
}

func Define(name string, initor func() error) *Module {
	mgr := GetModuleMgr()
	m := mgr.Registry(name)
	m.Fin = initor
	return m
}

func (self *Module) SetFin(f callback) *Module {
	self.Fin = f
	return self
}
