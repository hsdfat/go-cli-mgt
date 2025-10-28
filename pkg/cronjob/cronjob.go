package cronjob

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/robfig/cron/v3"
)

func initCronManager() {
	cronManager = &CronManager{
		c:               cron.New(),
		mapIdToCronFunc: make(map[string]*CronInfo),
	}
}

func addCron(cronInfo *CronInfo) {
	cronManager.mapIdToCronFunc[cronInfo.Identifier] = cronInfo
}

func runCronList() {
	for _, cronInfo := range cronManager.mapIdToCronFunc {
		_, err := cron.ParseStandard(cronInfo.TimeScheduler)
		if err != nil {
			logger.Logger.Error("Invalid cron expression for identifier:", cronInfo.Identifier, "err:", err)
			continue
		}

		jobEntryId, err := cronManager.c.AddJob(cronInfo.TimeScheduler, cron.FuncJob(cronInfo.HandlerFunc))
		if err != nil {
			logger.Logger.Errorf("Cannot run cronjob %s, err: %s", cronInfo.Name, err.Error())
			continue
		}
		cronInfo.IsActive = true
		cronInfo.JobId = jobEntryId
	}
	logger.Logger.Info("Init Cronjob Complete")
	cronManager.c.Start()
}

func InitCronjobList() {
	initCronManager()
	addCron(savingLogHistory)
	addCron(testCronjob)
	runCronList()
}

func TestCronjob() {
	logger.Logger.Info("Have been a while...")
}
