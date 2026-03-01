package query

import "database/sql"

type GetLandUnitResult struct {
	ID       string
	Name     string
	Type     string
	Length   float64
	Width    float64
	Sections []SectionDTO
	Beds     []BedDTO
}

type SectionDTO struct {
	ID     string
	Name   string
	Length float64
	Width  float64
	Beds   []BedDTO
}

type BedDTO struct {
	ID     string
	Name   string
	Length float64
	Width  float64
}

type GetLandUnitHandler struct {
	db *sql.DB
}

func NewGetLandUnitHandler(db *sql.DB) *GetLandUnitHandler {
	return &GetLandUnitHandler{db: db}
}

func (h *GetLandUnitHandler) Handle(id string) (*GetLandUnitResult, error) {

	rows, err := h.db.Query(`
		SELECT id, parent_id, unit_type, name, length, width
		FROM land_structure
		WHERE root_id = $1
	`, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var root *GetLandUnitResult
	sections := make(map[string]*SectionDTO)

	for rows.Next() {

		var rowID string
		var parentID *string
		var unitType string
		var name string
		var length float64
		var width float64

		if err := rows.Scan(
			&rowID,
			&parentID,
			&unitType,
			&name,
			&length,
			&width,
		); err != nil {
			return nil, err
		}

		if parentID == nil {
			root = &GetLandUnitResult{
				ID:     rowID,
				Name:   name,
				Type:   unitType,
				Length: length,
				Width:  width,
			}
			continue
		}

		if unitType == "section" {
			sec := &SectionDTO{
				ID:     rowID,
				Name:   name,
				Length: length,
				Width:  width,
			}
			sections[rowID] = sec
			root.Sections = append(root.Sections, *sec)
			continue
		}

		if unitType == "bed" {
			bed := BedDTO{
				ID:     rowID,
				Name:   name,
				Length: length,
				Width:  width,
			}

			if sec, ok := sections[*parentID]; ok {
				sec.Beds = append(sec.Beds, bed)
			} else {
				root.Beds = append(root.Beds, bed)
			}
		}
	}

	return root, nil
}
