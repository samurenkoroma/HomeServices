package cropplan

type Status string
type PlanID string
type GrowingAreaID string

// StageType для удобного использования в других пакетах
var StageTypes = struct {
	SoilPreparation StageType
	Sowing          StageType
	Fertilization   StageType
	Protection      StageType
	Irrigation      StageType
	Harvest         StageType
}{
	SoilPreparation: StageSoilPreparation,
	Sowing:          StageSowing,
	Fertilization:   StageFertilization,
	Protection:      StageProtection,
	Irrigation:      StageIrrigation,
	Harvest:         StageHarvest,
}

// StageStatuses для удобного использования
var StageStatuses = struct {
	Pending    StageStatus
	InProgress StageStatus
	Completed  StageStatus
	Skipped    StageStatus
}{
	Pending:    StageStatusPending,
	InProgress: StageStatusInProgress,
	Completed:  StageStatusCompleted,
	Skipped:    StageStatusSkipped,
}
