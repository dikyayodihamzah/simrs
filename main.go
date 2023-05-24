package main

import (
	"log"
	"time"

	"github.com/dikyayodihamzah/simrs/config"
	"github.com/dikyayodihamzah/simrs/controller"
	"github.com/dikyayodihamzah/simrs/handler"
	"github.com/dikyayodihamzah/simrs/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	time.Local = time.UTC

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/base-layout",
	})

	serverConfig := config.NewServerConfig()
	db := config.NewDB()

	roomRespoitory := repository.NewRoomRepository(db)
	impatientRespoitory := repository.NewImpatientRepository(db)
	emrRespoitory := repository.NewEMRRepository(db)
	patientRespoitory := repository.NewPatientRepository(db)
	symptomRespoitory := repository.NewSymptomRepository(db)

	dashboardHandler := handler.NewDashboardHandler(roomRespoitory, impatientRespoitory)
	patientHandler := handler.NewPatientHandler(patientRespoitory, emrRespoitory, symptomRespoitory)

	appController := controller.NewAppController(dashboardHandler, patientHandler)

	appController.Route(app)

	err := app.Listen(serverConfig.URI)
	log.Fatal(err)
}
