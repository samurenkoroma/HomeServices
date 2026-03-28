package queries

import (
	"context"
	"encoding/json"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/projections"
)

type GetCategoriesQuery struct {
	ID string `json:"id"`
}

type getCategoriesHandler struct {
	projector *projections.CropProjectionsProvider
}

func (h getCategoriesHandler) Handle(ctx context.Context, payload any) (any, error) {

	var response []croptype.CategoryDTO

	json.Unmarshal([]byte(data), &response)

	return response, nil
}

func ptr(s string) *string {
	return &s
}
func (h getCategoriesHandler) Name() string {
	return "GetCategories"
}

func NewGetCategoriesHandler(p *projections.CropProjectionsProvider) query.Handler {
	return &getCategoriesHandler{
		projector: p,
	}
}

var data = `[
    {
        "code": "cereal",
        "name": "Зерновые",
        "nameEn": "Cereals",
        "description": "Зерновые культуры для производства зерна",
        "parent": null,
        "subcategories": ["grain"]
    },
    {
        "code": "grain",
        "name": "Зернобобовые",
        "nameEn": "Grain Legumes",
        "description": "Бобовые культуры, выращиваемые на зерно",
        "parent": "cereal",
        "subcategories": []
    },
    {
        "code": "oilseed",
        "name": "Масличные",
        "nameEn": "Oilseeds",
        "description": "Культуры для получения растительных масел",
        "parent": null,
        "subcategories": []
    },
    {
        "code": "industrial",
        "name": "Технические",
        "nameEn": "Industrial",
        "description": "Культуры для промышленной переработки",
        "parent": null,
        "subcategories": []
    },
    {
        "code": "sugar",
        "name": "Сахароносные",
        "nameEn": "Sugar Crops",
        "description": "Культуры для производства сахара",
        "parent": null,
        "subcategories": []
    },
    {
        "code": "vegetable",
        "name": "Овощные",
        "nameEn": "Vegetables",
        "description": "Овощные культуры открытого и закрытого грунта",
        "parent": null,
        "subcategories": ["melon", "root"]
    },
    {
        "code": "melon",
        "name": "Бахчевые",
        "nameEn": "Melons",
        "description": "Бахчевые культуры (арбузы, дыни, тыквы)",
        "parent": "vegetable",
        "subcategories": []
    },
    {
        "code": "root",
        "name": "Корнеплоды",
        "nameEn": "Root Vegetables",
        "description": "Корнеплодные овощные культуры",
        "parent": "vegetable",
        "subcategories": []
    },
    {
        "code": "fruit",
        "name": "Плодовые",
        "nameEn": "Fruits",
        "description": "Плодовые деревья и кустарники",
        "parent": null,
        "subcategories": ["berry", "nut"]
    },
    {
        "code": "berry",
        "name": "Ягодные",
        "nameEn": "Berries",
        "description": "Ягодные культуры",
        "parent": "fruit",
        "subcategories": []
    },
    {
        "code": "nut",
        "name": "Орехоплодные",
        "nameEn": "Nuts",
        "description": "Орехоплодные культуры",
        "parent": "fruit",
        "subcategories": []
    },
    {
        "code": "forage",
        "name": "Кормовые",
        "nameEn": "Forage Crops",
        "description": "Культуры для кормопроизводства",
        "parent": null,
        "subcategories": ["silage", "green_manure"]
    },
    {
        "code": "silage",
        "name": "Силосные",
        "nameEn": "Silage Crops",
        "description": "Культуры для производства силоса",
        "parent": "forage",
        "subcategories": []
    },
    {
        "code": "green_manure",
        "name": "Сидераты",
        "nameEn": "Green Manure",
        "description": "Культуры для улучшения почвы",
        "parent": "forage",
        "subcategories": []
    },
    {
        "code": "medicinal",
        "name": "Лекарственные",
        "nameEn": "Medicinal Plants",
        "description": "Лекарственные растения",
        "parent": null,
        "subcategories": []
    },
    {
        "code": "spice",
        "name": "Пряные",
        "nameEn": "Spice Crops",
        "description": "Пряные и ароматические культуры",
        "parent": null,
        "subcategories": []
    },
    {
        "code": "ornamental",
        "name": "Декоративные",
        "nameEn": "Ornamental Plants",
        "description": "Декоративные растения для озеленения",
        "parent": null,
        "subcategories": ["floral"]
    },
    {
        "code": "floral",
        "name": "Цветочные",
        "nameEn": "Flowers",
        "description": "Цветочные культуры",
        "parent": "ornamental",
        "subcategories": []
    }
]`
