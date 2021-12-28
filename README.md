# go-scheduled-cron

#Install
```
go get github.com/onuragtas/go-scheduled-cron
```
#Usage

```
package cron

type Cron1 struct {

}

func (p *Cron1) Start()  {
	println("Cron1 started")
}
```
```
package cron

type Cron2 struct {

}

func (p *Cron2) Start()  {
	println("Cron2 started")
}
```

```
import (
	"github.com/onuragtas/go-scheduled-cron"
	"time"
)

var cronObject = cron.Cron{}

func init() {
	runCron()
}


func runCron () {
	cronObject.Add(&cron.Cron1{}, time.Millisecond * 1000 * 5)
	cronObject.Add(&cron.Cron2{}, time.Second * 10)
	cronObject.Run()
}
```