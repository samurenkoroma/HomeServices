package catalog

// StageType тип агрономической работы
type StageType string

const (
	StageSoilPreparation StageType = "soil_preparation"
	StageSowing          StageType = "sowing"
	StageFertilization   StageType = "fertilization"
	StageProtection      StageType = "protection"
	StageIrrigation      StageType = "irrigation"
	StagePruning         StageType = "pruning"
	StageHarvest         StageType = "harvest"
)

// StageTemplate шаблон этапа для культуры
type StageTemplate struct {
	Type        StageType `json:"type"`
	Name        string    `json:"name"`
	BBCHStart   int       `json:"bbch_start"`
	BBCHEnd     int       `json:"bbch_end"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	IsRequired  bool      `json:"is_required"`
}

// StageTemplatesBySpecies шаблоны этапов по видам культур
var StageTemplatesBySpecies = map[string][]StageTemplate{
	"tomato": {
		{Type: StageSoilPreparation, Name: "Подготовка почвы", BBCHStart: 0, BBCHEnd: 9,
			Description: "Вспашка, внесение базовых удобрений", Priority: "high", IsRequired: true},
		{Type: StageSowing, Name: "Посев семян", BBCHStart: 0, BBCHEnd: 9,
			Description: "Посев семян в грунт или на рассаду", Priority: "high", IsRequired: true},
		{Type: StagePruning, Name: "Пикировка", BBCHStart: 10, BBCHEnd: 19,
			Description: "Рассаживание растений по отдельным емкостям", Priority: "medium", IsRequired: false},
		{Type: StageFertilization, Name: "Подкормка азотом", BBCHStart: 19, BBCHEnd: 39,
			Description: "Внесение азотных удобрений для роста", Priority: "medium", IsRequired: true},
		{Type: StagePruning, Name: "Формирование куста", BBCHStart: 30, BBCHEnd: 39,
			Description: "Удаление пасынков, формирование стебля", Priority: "medium", IsRequired: true},
		{Type: StageProtection, Name: "Обработка от вредителей", BBCHStart: 50, BBCHEnd: 69,
			Description: "Опрыскивание от тли, клещей и других вредителей", Priority: "high", IsRequired: true},
		{Type: StageFertilization, Name: "Калийная подкормка", BBCHStart: 70, BBCHEnd: 79,
			Description: "Внесение калийных удобрений для плодоношения", Priority: "high", IsRequired: true},
		{Type: StageHarvest, Name: "Сбор урожая", BBCHStart: 80, BBCHEnd: 89,
			Description: "Сбор спелых плодов", Priority: "high", IsRequired: true},
	},
	"eggplant": {
		{Type: StageSoilPreparation, Name: "Подготовка почвы", BBCHStart: 0, BBCHEnd: 9,
			Description: "Вспашка, внесение удобрений", Priority: "high", IsRequired: true},
		{Type: StageSowing, Name: "Посев", BBCHStart: 0, BBCHEnd: 9,
			Description: "Посев семян", Priority: "high", IsRequired: true},
		{Type: StageFertilization, Name: "Подкормка", BBCHStart: 19, BBCHEnd: 39,
			Description: "Комплексная подкормка", Priority: "medium", IsRequired: true},
		{Type: StageProtection, Name: "Обработка от вредителей", BBCHStart: 50, BBCHEnd: 69,
			Description: "Защита от колорадского жука и тли", Priority: "high", IsRequired: true},
		{Type: StageHarvest, Name: "Сбор урожая", BBCHStart: 80, BBCHEnd: 89,
			Description: "Сбор плодов", Priority: "high", IsRequired: true},
	},
	"cucumber": {
		{Type: StageSoilPreparation, Name: "Подготовка почвы", BBCHStart: 0, BBCHEnd: 9,
			Description: "Вспашка, внесение удобрений", Priority: "high", IsRequired: true},
		{Type: StageSowing, Name: "Посев", BBCHStart: 0, BBCHEnd: 9,
			Description: "Посев семян", Priority: "high", IsRequired: true},
		{Type: StageFertilization, Name: "Подкормка", BBCHStart: 19, BBCHEnd: 39,
			Description: "Комплексная подкормка", Priority: "medium", IsRequired: true},
		{Type: StageProtection, Name: "Обработка от вредителей", BBCHStart: 50, BBCHEnd: 69,
			Description: "Защита от вредителей", Priority: "high", IsRequired: true},
		{Type: StageHarvest, Name: "Сбор урожая", BBCHStart: 80, BBCHEnd: 89,
			Description: "Сбор плодов", Priority: "high", IsRequired: true},
	},
}

// GetStageTemplatesForSpecies возвращает шаблоны этапов для вида
func GetStageTemplatesForSpecies(speciesKey string) []StageTemplate {
	templates, ok := StageTemplatesBySpecies[speciesKey]
	if !ok {
		return []StageTemplate{}
	}
	return templates
}
