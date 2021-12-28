package cron

import "time"

type ICron interface {
	Start()
}

type Cron struct {
	cronList []cron
}

type cron struct {
	Cron ICron
	Schedule time.Duration
	LastStart time.Time
}

func (c *Cron) Run() {
	for {
		for key, item := range c.cronList {
			if time.Now().Sub(item.LastStart) > item.Schedule {
				go item.Cron.Start()
				c.cronList[key].LastStart = time.Now()
			}
		}
	}
}

func (c *Cron) Add(iCron ICron, duration time.Duration) {
	c.cronList = append(c.cronList, cron{Cron: iCron, Schedule: duration})
}

