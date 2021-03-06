package cron

import "time"

type ICron interface {
	Start()
}

type Cron struct {
	cronList []cron
}

type cron struct {
	Cron      ICron
	Schedule  time.Duration
	LastStart time.Time
}

func (c *Cron) Run() {
	go func() {
		for {
			for key, item := range c.cronList {
				if time.Now().Sub(item.LastStart) > item.Schedule {
					go item.Cron.Start()
					c.cronList[key].LastStart = time.Now()
				}
			}
			time.Sleep(time.Millisecond * 200)
		}
	}()
}

func (c *Cron) Add(iCron ICron, duration time.Duration) {
	c.cronList = append(c.cronList, cron{Cron: iCron, Schedule: duration})
}
