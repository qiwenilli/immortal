package immortal

import (
	"log"
	"os"
	"os/user"
)

type Immortalizer interface {
	Daemon()
	Run()
	Supervise()
	WatchPid()
}

type Immortal struct {
	command []string
	*Config
}

func New() *Immortal {
	return &Immortal{}
}

//		Users:  &Users{},

func (self *Immortal) Daemon()    {}
func (self *Immortal) Run()       {}
func (self *Immortal) Supervise() {}
func (self *Immortal) WatchPid()  {}

type Daemon struct {
	command     []string
	count       uint32
	count_defer uint32
	ctrl        Ctrl
	log         bool
	logger      *log.Logger
	owner       *user.User
	process     *os.Process
	run         Run
}

type Run struct {
	Command   string
	Ctrl      bool
	Cwd       string
	Env       map[string]string
	Logfile   string
	Logger    string
	User      string
	ParentPid string
	ChildPid  string
	FollowPid string
}

type Ctrl struct {
	fifo         chan Return
	quit         chan struct{}
	state        chan error
	control_fifo *os.File
	status_fifo  *os.File
}

type Return struct {
	err error
	msg string
}
