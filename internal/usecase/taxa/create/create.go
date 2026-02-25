package create

import (
	"context"
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/domain/taxonomy"
	"samurenkoroma/services/internal/infrastructure/db_table"

	nestedset "github.com/longbridgeapp/nested-set"
	"gorm.io/gorm"
)

type UC struct {
	conn    *gorm.DB
	payload *Payload
}

type Payload struct {
	Rank   string
	Type   string
	Name   string
	Parent *db_table.TaxonomyNode
}

func NewUC(conn *gorm.DB) *UC {
	return &UC{
		conn:    conn,
		payload: &Payload{},
	}
}

func (uc *UC) Payload(dto *Payload) *UC {
	uc.payload = dto
	return uc
}
func (uc *UC) Execute(ctx context.Context) (*db_table.TaxonomyNode, error) {
	if uc.payload == nil {
		return nil, fmt.Errorf("payload is nil, fill payload")
	}
	node, err := gorm.G[db_table.TaxonomyNode](uc.conn).
		Where("rank = ? AND name = ?", uc.payload.Rank, uc.payload.Name).
		First(ctx)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("будет создана новая таксономия %v\n", uc.payload)
	}
	if node.ID == 0 {
		node.Name = uc.payload.Name
		node.Type = taxonomy.TypeFromString(uc.payload.Type)
		node.Rank = uc.payload.Rank

		if uc.payload.Parent != nil {
			node.ParentID = sql.NullInt64{Valid: true, Int64: uc.payload.Parent.ID}
			node.Type = uc.payload.Parent.Type
		}
		if err := nestedset.Create(uc.conn, &node, uc.payload.Parent); err != nil {
			return &db_table.TaxonomyNode{}, err
		}
	}
	return &node, nil
}
