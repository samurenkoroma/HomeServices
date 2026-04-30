package seasons

import (
	"context"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"time"
)

type ListSeasonsQuery struct {
}

func (h *QueryHandler) List(ctx context.Context, query any) (any, error) {
	_, ok := query.(*ListSeasonsQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}
	response := []SeasonDTO{}
	orgId, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	fmt.Print(orgId)

	data, err := h.seasons.FindAll(ctx, season.Filter{
		OwnerId: orgId,
	})
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		response = append(response, SeasonDTO{
			Id:        string(d.Id),
			Name:      d.GetName(),
			StartDate: d.GetStartDate(),
			EndDate:   d.GetEndDate(),
			Status:    string(d.GetStatus()),
			Statistics: Statistics{
				Crops: []Crops{},
			},
		})
	}
	return response, nil
}

type SeasonDTO struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Year          int       `json:"year"`
	Type          string    `json:"type"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	Status        string    `json:"status"`
	Statistics    `json:"statistics"`
	Weather       `json:"weather"`
	Notes         string `json:"notes"`
	Financial     `json:"financial"`
	PlantingAreas []PlantingArea `json:"plantingArea"`
}

type Financial struct {
	Revenue      float64 `json:"revenue"`
	Costs        float64 `json:"costs"`
	Profit       float64 `json:"profit"`
	ProfitMargin float64 `json:"profitMargin"`
}

type Weather struct {
	AvgTemp            float64 `json:"avgTemp"`
	TotalPrecipitation float64 `json:"totalPrecipitation"`
	SunnyDays          int     `json:"sunnyDays"`
}
type Statistics struct {
	TotalPlans     int     `json:"totalPlans"`
	CompletedPlans int     `json:"completedPlans"`
	ActivePlans    int     `json:"activePlans"`
	TotalArea      float64 `json:"totalArea"`
	TotalHarvest   float64 `json:"totalHarvest"`
	AvgYield       float64 `json:"avgYield"`
	Crops          []Crops `json:"crops"`
}

type Crops struct {
	Name       string  `json:"name"`
	Area       float64 `json:"area"`
	Yield      float64 `json:"yield"`
	Icon       string  `json:"icon"`
	YieldPerHa float64 `json:"yieldPerHa"`
}

type PlantingArea struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Area            float64 `json:"area"`
	VarietyId       string  `json:"varietyId"`
	VarietyName     string  `json:"varietyName"`
	CropName        string  `json:"cropName"`
	PlantedDate     string  `json:"plantedDate"`
	HarvestDate     string  `json:"harvestDate"`
	ActualYield     float64 `json:"actualYield"`
	ExpectedYield   float64 `json:"expectedYield"`
	YieldEfficiency float64 `json:"yieldEfficiency"`
	Resources       []struct {
		WaterUsed      float64 `json:"waterUsed"`
		FertilizerUsed []struct {
			Name   string  `json:"name"`
			Amount float64 `json:"amount"`
			Unit   string  `json:"unit"`
		} `json:"fertilizerUsed"`
		FuelUsed   float64 `json:"fuelUsed"`
		LaborHours float64 `json:"laborHours"`
		SeedsUsed  float64 `json:"seedsUsed"`
	} `json:"resources"`
	Deviations []struct {
		Type        string `json:"type"`
		Description string `json:"description"`
		Impact      string `json:"impact"`
	} `json:"deviations"`
}
