package controller

import (
	"github.com/dikyayodihamzah/simrs/handler"
	"github.com/gofiber/fiber/v2"
)

type AppController interface {
	Route(app *fiber.App)
}

type appController struct {
	DashboardHandler handler.DashboardHandler
	PatientHandler   handler.PatientHandler
}

func NewAppController(dashboardHandler handler.DashboardHandler, patientHandler handler.PatientHandler) AppController {
	return &appController{
		DashboardHandler: dashboardHandler,
		PatientHandler:   patientHandler,
	}
}

func (controller *appController) Route(app *fiber.App) {
	app.Static("/", "./static")

	// app.Get("/login", handler.Login)

	// dashboard routes
	app.Get("/", controller.DashboardHandler.Dashboard)

	// patient routes
	app.Get("/patients", controller.PatientHandler.GetAllPatients)
	app.Get("/patients/:id", controller.PatientHandler.GetAllPatients)
	app.Get("/input-patient", controller.PatientHandler.InputPatient)
	// app.Post("/patients/input", controller.AppHandler.GetAllPatients)
}
