package croptemplate

type CropTemplate struct {
	planID  string
	version int
	active  bool
}

func New(planID string, version int) *CropTemplate {
	return &CropTemplate{
		planID:  planID,
		version: version,
		active:  true,
	}
}

func (t *CropTemplate) Disable() {
	t.active = false
}

func (t *CropTemplate) Active() bool {
	return t.active
}

func (t *CropTemplate) Version() int {
	return t.version
}

func (t *CropTemplate) PlanID() string {
	return t.planID
}
