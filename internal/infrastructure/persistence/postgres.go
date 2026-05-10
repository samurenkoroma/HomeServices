package persistence

import "strings"

// ParsePostgresArray преобразует строку массива PostgreSQL в []string
// Формат PostgreSQL массива: {value1,value2,value3} или {"value with space","value2"}
func ParsePostgresArray(arrStr string) []string {
	if arrStr == "" || arrStr == "{}" {
		return []string{}
	}

	// Убираем фигурные скобки
	content := arrStr[1 : len(arrStr)-1]
	if content == "" {
		return []string{}
	}

	// Разделяем по запятой
	parts := strings.Split(content, ",")
	result := make([]string, 0, len(parts))

	for _, p := range parts {
		// Убираем кавычки, если есть
		p = strings.TrimSpace(p)
		p = strings.Trim(p, "\"")
		result = append(result, p)
	}

	return result
}
