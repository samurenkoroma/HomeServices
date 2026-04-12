package bbch

import "errors"

type Phase string

const (
	PhaseGermination      Phase = "germination"       // 00-09
	PhaseLeafDevelopment  Phase = "leaf_development"  // 10-19
	PhaseStemElongation   Phase = "stem_elongation"   // 30-39
	PhaseBudDevelopment   Phase = "bud_development"   // 50-59
	PhaseFlowering        Phase = "flowering"         // 60-69
	PhaseFruitDevelopment Phase = "fruit_development" // 70-79
	PhaseRipening         Phase = "ripening"          // 80-89
	PhaseSenescence       Phase = "senescence"        // 90-99
)

type BBCHRange struct {
	Start int // код начала фазы
	End   int // код конца фазы
	Phase Phase
}

var BBCHRanges = map[Phase]BBCHRange{
	PhaseGermination:      {Start: 0, End: 9, Phase: PhaseGermination},
	PhaseLeafDevelopment:  {Start: 10, End: 19, Phase: PhaseLeafDevelopment},
	PhaseStemElongation:   {Start: 30, End: 39, Phase: PhaseStemElongation},
	PhaseBudDevelopment:   {Start: 50, End: 59, Phase: PhaseBudDevelopment},
	PhaseFlowering:        {Start: 60, End: 69, Phase: PhaseFlowering},
	PhaseFruitDevelopment: {Start: 70, End: 79, Phase: PhaseFruitDevelopment},
	PhaseRipening:         {Start: 80, End: 89, Phase: PhaseRipening},
	PhaseSenescence:       {Start: 90, End: 99, Phase: PhaseSenescence},
}

func (r BBCHRange) Contains(code int) bool {
	return code >= r.Start && code <= r.End
}

// Определить фазу по BBCH коду
func GetPhase(code int) (Phase, error) {
	for _, r := range BBCHRanges {
		if r.Contains(code) {
			return r.Phase, nil
		}
	}
	return "", errors.New("invalid BBCH code")
}
