package jobs

import (
	"mysql_api/databases"
	"mysql_api/settings"
	"mysql_api/structs/models"
	"time"
)

// 每天将status置为0以重查
func UpdateStatus() {
	if isBetweenUpdateTime() {
		databases.Db.Model(models.Product{}).
			Where("status = ? or status = ?", models.StatusSearching, models.StatusSearchSucceed).
			Update(map[string]interface{}{"status": 0})
	}
	time.Sleep(time.Minute)
}

func isBetweenUpdateTime() bool {
	now := time.Now()
	UpdateStartTime, _ := time.ParseInLocation("2006-01-02 15:04:05", now.Format("2006-01-02")+" "+settings.StartTime, now.Location())
	UpdateEndTime, _ := time.ParseInLocation("2006-01-02 15:04:05", now.Format("2006-01-02")+" "+settings.EndTime, now.Location())

	return now.Sub(UpdateStartTime) > 0 && now.Sub(UpdateEndTime) < 0
}
