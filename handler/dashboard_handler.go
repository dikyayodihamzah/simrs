package handler

import (
	"fmt"
	"strconv"

	"github.com/dikyayodihamzah/simrs/repository"
	"github.com/gofiber/fiber/v2"
)

type DashboardHandler interface {
	Dashboard(c *fiber.Ctx) error
}

type dashboardHandler struct {
	RoomRepository      repository.RoomRepository
	ImpatientRepository repository.ImpatientRepository
}

func NewDashboardHandler(roomRepository repository.RoomRepository, impatientRepository repository.ImpatientRepository) DashboardHandler {
	return &dashboardHandler{
		RoomRepository:      roomRepository,
		ImpatientRepository: impatientRepository,
	}
}

func (handler *dashboardHandler) Dashboard(c *fiber.Ctx) error {
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
