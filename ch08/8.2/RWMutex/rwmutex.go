package rwmutex

import (
	"fmt"
	"sync"
	"time"
)

type pass struct {
	RWM sync.RWMutex
	pwd string
}
var RoomPass = pass{pwd:"initPass"}

func Change(p *pass,pwd string)  {
	p.RWM.Lock()
	fmt.Println()
	time.Sleep(5*time.Second)
	p.pwd = pwd
	p.RWM.Unlock()
}

func getPWD(p *pass) string {
	p.RWM.RLock()
	fmt.Println("read pwd")
	time.Sleep(time.Second)
	defer p.RWM.RUnlock()
	return p.pwd
}




