package finance

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"samurenkoroma/services/internal/app"
)

type Handler struct {
	router fiber.Router
}

func NewFinanceHandler(app *app.Polevod) {
	h := Handler{
		router: app.App,
	}
	h.router.Post("/invoice", h.createInvoice)
}

// UserData represents the JSON data structure
type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h Handler) createInvoice(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Error parsing multipart form: %v", err))
	}

	// Get the file
	fileHeader := form.File["file"]
	if len(fileHeader) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("No file uploaded")
	}
	file := fileHeader[0]

	// Get the JSON data
	jsonData := form.Value["json_data"]
	if len(jsonData) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("No JSON data provided")
	}

	// Unmarshal JSON data
	var userData UserData
	if err := c.BodyParser(&userData); err != nil { // Note: Fiber's BodyParser can handle form-data values if tagged correctly
		// Alternatively, you can unmarshal from jsonData[0] directly
		// if err := json.Unmarshal([]byte(jsonData[0]), &userData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Error unmarshaling JSON: %v", err))
	}

	// Save the file
	filePath := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error saving file: %v", err))
	}

	log.Printf("File '%s' saved to '%s'", file.Filename, filePath)
	log.Printf("Received JSON data: Name=%s, Email=%s", userData.Name, userData.Email)

	return c.SendString("File and JSON received successfully!")
}
