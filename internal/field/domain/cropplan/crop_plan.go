package cropplan

import "time"

type Status string

const (
	Planned   Status = "planned"
	Growing   Status = "growing"
	Harvested Status = "harvested"
)

type CropPlan struct {
	id        CropPlanID
	bedID     BedID
	cropName  string
	status    Status
	createdAt time.Time
	harvestKg float64
}

func (c *CropPlan) ID() CropPlanID {
	return c.id
}

func New(id CropPlanID, bedID BedID, crop string) *CropPlan {
	return &CropPlan{
		id:        id,
		bedID:     bedID,
		cropName:  crop,
		status:    Planned,
		createdAt: time.Now(),
	}
}

func (c *CropPlan) StartGrowing() {
	c.status = Growing
}

func (c *CropPlan) Harvest(kg float64) error {
	if c.status == Harvested {
		return ErrAlreadyHarvested
	}
	c.status = Harvested
	c.harvestKg = kg
	return nil
}

func (c *CropPlan) BedID() BedID {
	return c.bedID
}

func (c *CropPlan) CropName() string {
	return c.cropName
}

func (c *CropPlan) Status() Status {
	return c.status
}

func (c *CropPlan) HarvestKg() float64 {
	return c.harvestKg
}
