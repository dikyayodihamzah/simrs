package handler

import (
	"fmt"
	"strconv"

	"github.com/dikyayodihamzah/simrs/repository"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AppHandler interface {
	Dashboard(c *fiber.Ctx) error
}
type appHandler struct {
	RoomRepository      repository.RoomRepository
	ImpatientRepository repository.ImpatientRepository
	Validate            *validator.Validate
}

func NewAppHandler(roomRepository repository.RoomRepository, impatientRepository repository.ImpatientRepository, validate *validator.Validate) AppHandler {
	return &appHandler{
		RoomRepository:      roomRepository,
		ImpatientRepository: impatientRepository,
		Validate:            validate,
	}
}

func (handler *appHandler) Dashboard(c *fiber.Ctx) error {
	var vipBed int
	var firstClassBed int
	var secondClassBed int

	beds, _ := handler.RoomRepository.CountBeds(c.Context())
	for _, bed := range beds {
		switch bed["_id"] {
		case "1":
			firstClassBed, _ = strconv.Atoi(fmt.Sprint(bed["bed_total"]))
		case "2":
			secondClassBed, _ = strconv.Atoi(fmt.Sprint(bed["bed_total"]))
		case "VIP":
			vipBed, _ = strconv.Atoi(fmt.Sprint(bed["bed_total"]))
		}
	}

	totalBed := vipBed + firstClassBed + secondClassBed
	totalPatient, _ := handler.ImpatientRepository.CountPatients(c.Context())

	var usedRoom int
	results, _ := handler.RoomRepository.CountNotEmptyRooms(c.Context())
	if results == nil {
		usedRoom = 0
	} else {
		usedRoom, _ = strconv.Atoi(fmt.Sprint(results[0]["room_total"]))
	}

	totalRoom, _ := handler.RoomRepository.CountRooms(c.Context())

	var vipPatient int
	var firstClassPatient int
	var secondClassPatient int

	patients, _ := handler.ImpatientRepository.CountPatientbyClass(c.Context())
	for _, patient := range patients {
		switch patient["_id"] {
		case "1":
			firstClassPatient, _ = strconv.Atoi(fmt.Sprint(patient["patient_total"]))
		case "2":
			secondClassPatient, _ = strconv.Atoi(fmt.Sprint(patient["patient_total"]))
		case "VIP":
			vipPatient, _ = strconv.Atoi(fmt.Sprint(patient["patient_total"]))
		}
	}

	return c.Render("pages-dashboard", fiber.Map{
		"Title":              "Dashboard",
		"TotalPatient":       totalPatient,
		"TotalBed":           totalBed,
		"UsedRoom":           usedRoom,
		"TotalRoom":          totalRoom,
		"VIPPatient":         vipPatient,
		"VIPBed":             vipBed,
		"FirstClassPatient":  firstClassPatient,
		"FirstClassBed":      firstClassBed,
		"SecondClassPatient": secondClassPatient,
		"SecondClassBed":     secondClassBed,
	})
}

func Login(c *fiber.Ctx) error {
	return c.Render("pages-login", fiber.Map{})
}

func InformasiKamar(c *fiber.Ctx) error {
	return c.Render("pages-data-pasien", fiber.Map{})
}

func DataPasien(c *fiber.Ctx) error {
	return c.Render("dashboard", fiber.Map{})
}

func About(c *fiber.Ctx) error {
	return c.Render("dashboard", fiber.Map{})
}

func Kontak(c *fiber.Ctx) error {
	return c.Render("pages-contact", fiber.Map{})
}
