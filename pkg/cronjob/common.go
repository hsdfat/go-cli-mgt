package cronjob

import (
	"github.com/hsdfat/go-cli-mgt/pkg/svc"

	"github.com/robfig/cron/v3"
)

const (
	EveryOneMinute      = "EveryOneMinute"
	EveryFiveMinutes    = "EveryFiveMinute"
	EveryTenMinutes     = "EveryTenMinutes"
	EveryFifteenMinutes = "EveryFifteenMinutes"
	EveryThirtyMinutes  = "EveryThirtyMinutes"
	EveryHour           = "EveryHour"
	EveryDay            = "EveryDay" // 00h 05p
)

var mapStringToTimeScheduler = map[string]string{
	EveryOneMinute:      "*/1 * * * *",
	EveryFiveMinutes:    "*/5 * * * *",
	EveryTenMinutes:     "*/10 * * * *",
	EveryFifteenMinutes: "*/15 * * * *",
	EveryThirtyMinutes:  "*/30 * * * *",
	EveryHour:           "* */1 * * *",
	EveryDay:            "5 0 * * *",
}

var mapTimeSchedulerToString = map[string]string{
	"*/1 * * * *":  EveryOneMinute,
	"*/5 * * * *":  EveryFiveMinutes,
	"*/10 * * * *": EveryTenMinutes,
	"*/15 * * * *": EveryFifteenMinutes,
	"*/30 * * * *": EveryThirtyMinutes,
	"* */1 * * *":  EveryHour,
	"5 0 * * *":    EveryDay,
}

type CronManager struct {
	c               *cron.Cron
	mapIdToCronFunc map[string]*CronInfo
}

type CronInfo struct {
	Identifier    string
	JobId         cron.EntryID
	Name          string
	TimeScheduler string
	IsActive      bool
	HandlerFunc   func()
}

var cronManager *CronManager

var (
	savingLogHistory = &CronInfo{
		Identifier:    "SLH",
		Name:          "Cronjob Saving Log history everyday",
		TimeScheduler: mapStringToTimeScheduler[EveryDay],
		HandlerFunc:   svc.SavingLogHistory,
	}

	testCronjob = &CronInfo{
		Identifier:    "TCJ",
		Name:          "Cronjob for testing",
		TimeScheduler: mapStringToTimeScheduler[EveryOneMinute],
		HandlerFunc:   TestCronjob,
	}
)
