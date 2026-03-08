package cropcycle

import "time"

type CropCycle struct {
	id          string
	planID      string
	planVersion int

	facilityID string
	bedID      string

	status     Status
	startedAt  *time.Time
	finishedAt *time.Time
}

func (c *CropCycle) Id() string {
	return c.id
}

func (c *CropCycle) PlanVersion() int {
	return c.planVersion
}

func (c *CropCycle) FacilityID() string {
	return c.facilityID
}

func (c *CropCycle) BedID() string {
	return c.bedID
}

func (c *CropCycle) StartedAt() *time.Time {
	return c.startedAt
}

func (c *CropCycle) FinishedAt() *time.Time {
	return c.finishedAt
}

func New(
	id string,
	planID string,
	version int,
	facilityID string,
	bedID string,
) *CropCycle {

	return &CropCycle{
		id:          id,
		planID:      planID,
		planVersion: version,
		facilityID:  facilityID,
		bedID:       bedID,
		status:      Draft,
	}
}

func (c *CropCycle) Start() error {
	if c.status != Draft {
		return ErrInvalidState
	}

	now := time.Now()

	c.status = Active
	c.startedAt = &now

	return nil
}

func (c *CropCycle) ID() string {
	return c.id
}

func (c *CropCycle) PlanID() string {
	return c.planID
}

func (c *CropCycle) Status() Status {
	return c.status
}

func (c *CropCycle) Version() int {
	return 0
}

func (c *CropCycle) SetStatus(status Status) {
	c.status = status
}

func Rehydrate(
	id string,
	planID string,
	facilityID string,
	bedID string,
	status Status,
	version int,
) *CropCycle {

	return &CropCycle{
		id:          id,
		planID:      planID,
		planVersion: version,
		facilityID:  facilityID,
		bedID:       bedID,
		status:      status,
	}
}
