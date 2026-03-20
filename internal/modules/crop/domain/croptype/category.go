package croptype

import (
	"fmt"
	"strings"
)

// CropCategory — категория сельскохозяйственной культуры
type CropCategory string

// Определение категорий культур
const (
	// Зерновые культуры
	CategoryCereal CropCategory = "cereal" // Зерновые (пшеница, рис, кукуруза)
	CategoryGrain  CropCategory = "grain"  // Зернобобовые (горох, фасоль)

	// Технические культуры
	CategoryOilseed    CropCategory = "oilseed"    // Масличные (подсолнечник, рапс)
	CategoryIndustrial CropCategory = "industrial" // Технические (лен, хлопок)
	CategorySugar      CropCategory = "sugar"      // Сахароносные (сахарная свекла)

	// Овощные культуры
	CategoryVegetable CropCategory = "vegetable" // Овощные (томаты, огурцы)
	CategoryMelon     CropCategory = "melon"     // Бахчевые (арбуз, дыня)
	CategoryRoot      CropCategory = "root"      // Корнеплоды (морковь, картофель)

	// Плодовые культуры
	CategoryFruit CropCategory = "fruit" // Плодовые (яблоки, груши)
	CategoryBerry CropCategory = "berry" // Ягодные (клубника, малина)
	CategoryNut   CropCategory = "nut"   // Орехоплодные (грецкий орех)

	// Кормовые культуры
	CategoryForage CropCategory = "forage" // Кормовые (люцерна, клевер)
	CategorySilage CropCategory = "silage" // Силосные (кукуруза на силос)

	// Специальные
	CategoryGreenManure CropCategory = "green_manure" // Сидераты
	CategoryCover       CropCategory = "cover"        // Покровные культуры
	CategoryMedicinal   CropCategory = "medicinal"    // Лекарственные
	CategorySpice       CropCategory = "spice"        // Пряные

	// Декоративные
	CategoryOrnamental CropCategory = "ornamental" // Декоративные
	CategoryFloral     CropCategory = "floral"     // Цветочные
)

// CategoryInfo — информация о категории
type CategoryInfo struct {
	Code          CropCategory   `json:"code"`
	Name          string         `json:"name"`
	NameEn        string         `json:"name_en"`
	Description   string         `json:"description"`
	Parent        *CropCategory  `json:"parent,omitempty"`
	Subcategories []CropCategory `json:"subcategories,omitempty"`
}

// Создаем переменные для указателей, так как нельзя брать адрес константы
var (
	ptrCereal     = CategoryCereal
	ptrVegetable  = CategoryVegetable
	ptrFruit      = CategoryFruit
	ptrForage     = CategoryForage
	ptrOrnamental = CategoryOrnamental
)

// newCategoryInfo — конструктор для CategoryInfo
func newCategoryInfo(
	code CropCategory,
	name, nameEn, description string,
	parent *CropCategory,
	subcategories []CropCategory,
) CategoryInfo {
	return CategoryInfo{
		Code:          code,
		Name:          name,
		NameEn:        nameEn,
		Description:   description,
		Parent:        parent,
		Subcategories: subcategories,
	}
}

// categoryRegistry — реестр всех категорий
var categoryRegistry = map[CropCategory]CategoryInfo{
	// Зерновые
	CategoryCereal: newCategoryInfo(
		CategoryCereal,
		"Зерновые",
		"Cereals",
		"Зерновые культуры для производства зерна",
		nil,
		[]CropCategory{CategoryGrain},
	),
	CategoryGrain: newCategoryInfo(
		CategoryGrain,
		"Зернобобовые",
		"Grain Legumes",
		"Бобовые культуры, выращиваемые на зерно",
		&ptrCereal, // Используем переменную, а не константу
		nil,
	),

	// Технические
	CategoryOilseed: newCategoryInfo(
		CategoryOilseed,
		"Масличные",
		"Oilseeds",
		"Культуры для получения растительных масел",
		nil,
		nil,
	),
	CategoryIndustrial: newCategoryInfo(
		CategoryIndustrial,
		"Технические",
		"Industrial",
		"Культуры для промышленной переработки",
		nil,
		nil,
	),
	CategorySugar: newCategoryInfo(
		CategorySugar,
		"Сахароносные",
		"Sugar Crops",
		"Культуры для производства сахара",
		nil,
		nil,
	),

	// Овощные
	CategoryVegetable: newCategoryInfo(
		CategoryVegetable,
		"Овощные",
		"Vegetables",
		"Овощные культуры открытого и закрытого грунта",
		nil,
		[]CropCategory{CategoryMelon, CategoryRoot},
	),
	CategoryMelon: newCategoryInfo(
		CategoryMelon,
		"Бахчевые",
		"Melons",
		"Бахчевые культуры (арбузы, дыни, тыквы)",
		&ptrVegetable,
		nil,
	),
	CategoryRoot: newCategoryInfo(
		CategoryRoot,
		"Корнеплоды",
		"Root Vegetables",
		"Корнеплодные овощные культуры",
		&ptrVegetable,
		nil,
	),

	// Плодовые
	CategoryFruit: newCategoryInfo(
		CategoryFruit,
		"Плодовые",
		"Fruits",
		"Плодовые деревья и кустарники",
		nil,
		[]CropCategory{CategoryBerry, CategoryNut},
	),
	CategoryBerry: newCategoryInfo(
		CategoryBerry,
		"Ягодные",
		"Berries",
		"Ягодные культуры",
		&ptrFruit,
		nil,
	),
	CategoryNut: newCategoryInfo(
		CategoryNut,
		"Орехоплодные",
		"Nuts",
		"Орехоплодные культуры",
		&ptrFruit,
		nil,
	),

	// Кормовые
	CategoryForage: newCategoryInfo(
		CategoryForage,
		"Кормовые",
		"Forage Crops",
		"Культуры для кормопроизводства",
		nil,
		[]CropCategory{CategorySilage, CategoryGreenManure},
	),
	CategorySilage: newCategoryInfo(
		CategorySilage,
		"Силосные",
		"Silage Crops",
		"Культуры для производства силоса",
		&ptrForage,
		nil,
	),
	CategoryGreenManure: newCategoryInfo(
		CategoryGreenManure,
		"Сидераты",
		"Green Manure",
		"Культуры для улучшения почвы",
		&ptrForage,
		nil,
	),
	CategoryCover: newCategoryInfo(
		CategoryCover,
		"Покровные",
		"Cover Crops",
		"Культуры для защиты почвы",
		nil,
		nil,
	),

	// Специальные
	CategoryMedicinal: newCategoryInfo(
		CategoryMedicinal,
		"Лекарственные",
		"Medicinal Plants",
		"Лекарственные растения",
		nil,
		nil,
	),
	CategorySpice: newCategoryInfo(
		CategorySpice,
		"Пряные",
		"Spice Crops",
		"Пряные и ароматические культуры",
		nil,
		nil,
	),

	// Декоративные
	CategoryOrnamental: newCategoryInfo(
		CategoryOrnamental,
		"Декоративные",
		"Ornamental Plants",
		"Декоративные растения для озеленения",
		nil,
		[]CropCategory{CategoryFloral},
	),
	CategoryFloral: newCategoryInfo(
		CategoryFloral,
		"Цветочные",
		"Flowers",
		"Цветочные культуры",
		&ptrOrnamental,
		nil,
	),
}

// GetAllCategories возвращает все доступные категории
func GetAllCategories() []CategoryInfo {
	categories := make([]CategoryInfo, 0, len(categoryRegistry))
	for _, info := range categoryRegistry {
		categories = append(categories, info)
	}
	return categories
}

// GetCategoryInfo возвращает информацию о категории
func GetCategoryInfo(category CropCategory) (CategoryInfo, bool) {
	info, ok := categoryRegistry[category]
	return info, ok
}

// IsValidCategory проверяет, существует ли категория
func IsValidCategory(category CropCategory) bool {
	_, ok := categoryRegistry[category]
	return ok
}

// GetCategoriesByParent возвращает все дочерние категории
func GetCategoriesByParent(parent CropCategory) []CategoryInfo {
	var result []CategoryInfo
	for _, info := range categoryRegistry {
		if info.Parent != nil && *info.Parent == parent {
			result = append(result, info)
		}
	}
	return result
}

// GetRootCategories возвращает все корневые категории
func GetRootCategories() []CategoryInfo {
	var result []CategoryInfo
	for _, info := range categoryRegistry {
		if info.Parent == nil {
			result = append(result, info)
		}
	}
	return result
}

// CategoryHierarchy — иерархия категорий
type CategoryHierarchy struct {
	Category CategoryInfo        `json:"category"`
	Children []CategoryHierarchy `json:"children"`
}

// BuildHierarchy строит иерархию категорий
func BuildHierarchy() []CategoryHierarchy {
	rootCategories := GetRootCategories()
	hierarchy := make([]CategoryHierarchy, 0, len(rootCategories))

	for _, root := range rootCategories {
		hierarchy = append(hierarchy, CategoryHierarchy{
			Category: root,
			Children: buildChildren(root.Code),
		})
	}

	return hierarchy
}

func buildChildren(parent CropCategory) []CategoryHierarchy {
	children := GetCategoriesByParent(parent)
	result := make([]CategoryHierarchy, 0, len(children))

	for _, child := range children {
		result = append(result, CategoryHierarchy{
			Category: child,
			Children: buildChildren(child.Code),
		})
	}

	return result
}

// CategoryString возвращает строковое представление категории
func (c CropCategory) String() string {
	return string(c)
}

// DisplayName возвращает отображаемое имя категории
func (c CropCategory) DisplayName() string {
	if info, ok := categoryRegistry[c]; ok {
		return info.Name
	}
	return string(c)
}

// DisplayNameEn возвращает английское имя категории
func (c CropCategory) DisplayNameEn() string {
	if info, ok := categoryRegistry[c]; ok {
		return info.NameEn
	}
	return string(c)
}

// Description возвращает описание категории
func (c CropCategory) Description() string {
	if info, ok := categoryRegistry[c]; ok {
		return info.Description
	}
	return ""
}

// IsValid проверяет валидность категории
func (c CropCategory) IsValid() bool {
	return IsValidCategory(c)
}

// GetParent возвращает родительскую категорию
func (c CropCategory) GetParent() *CropCategory {
	if info, ok := categoryRegistry[c]; ok {
		return info.Parent
	}
	return nil
}

// GetChildren возвращает дочерние категории
func (c CropCategory) GetChildren() []CropCategory {
	if info, ok := categoryRegistry[c]; ok {
		return info.Subcategories
	}
	return nil
}

// MarshalJSON реализует JSON marshaling для CropCategory
func (c CropCategory) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(c) + `"`), nil
}

// UnmarshalJSON реализует JSON unmarshaling для CropCategory
func (c *CropCategory) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	*c = CropCategory(str)

	if !c.IsValid() {
		return fmt.Errorf("invalid crop category: %s", str)
	}

	return nil
}

// CategoryList — список категорий для использования в API
type CategoryList struct {
	Categories []CategoryInfo      `json:"categories"`
	Hierarchy  []CategoryHierarchy `json:"hierarchy"`
}

// GetCategoryList возвращает полный список категорий
func GetCategoryList() CategoryList {
	return CategoryList{
		Categories: GetAllCategories(),
		Hierarchy:  BuildHierarchy(),
	}
}
