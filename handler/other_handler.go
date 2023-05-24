package handler

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/dikyayodihamzah/simrs/model/domain"
// 	"github.com/dikyayodihamzah/simrs/model/web"
// 	"github.com/dikyayodihamzah/simrs/repository"
// 	"github.com/go-playground/validator"
// 	"github.com/gofiber/fiber/v2"
// )

// type AppHandler interface {
// 	Dashboard(c *fiber.Ctx) error
// 	GetAllPatients(c *fiber.Ctx) error
// 	InputPatient(c *fiber.Ctx) error
// }
// type appHandler struct {
// 	RoomRepository      repository.RoomRepository
// 	ImpatientRepository repository.ImpatientRepository
// 	PatientRepository   repository.PatientRepository
// 	EMRRepository       repository.EMRRepository
// 	Validate            *validator.Validate
// }

// func NewAppHandler(roomRepository repository.RoomRepository, impatientRepository repository.ImpatientRepository, emrRepository repository.EMRRepository, patientRepository repository.PatientRepository, validate *validator.Validate) AppHandler {
// 	return &appHandler{
// 		RoomRepository:      roomRepository,
// 		ImpatientRepository: impatientRepository,
// 		PatientRepository:   patientRepository,
// 		EMRRepository:       emrRepository,
// 		Validate:            validate,
// 	}
// }

// func Login(c *fiber.Ctx) error {
// 	return c.Render("pages-login", fiber.Map{})
// }
