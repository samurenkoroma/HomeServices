package stat

import (
	"samurenkoroma/services/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	database *db.Db
}

func NewStatRepo(database *db.Db) *StatRepository {
	return &StatRepository{
		database: database,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.database.Find(&stat, "link_id = ? and date = ? ", linkId, currentDate)

	if stat.ID == 0 {
		repo.database.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.database.Save(stat)
	}
}

func (repo *StatRepository) GetStats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse

	var selectQuery string

	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, SUM(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, SUM(clicks)"
	}

	repo.database.
		Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? and ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)
		
	return stats
}
