package handler

import (
	"github.com/dikyayodihamzah/simrs/exception"
	"github.com/dikyayodihamzah/simrs/model/domain"
	"github.com/dikyayodihamzah/simrs/model/web"
	"github.com/dikyayodihamzah/simrs/repository"
	"github.com/gofiber/fiber/v2"
)

type PatientHandler interface {
	GetAllPatients(c *fiber.Ctx) error
	InputPatient(c *fiber.Ctx) error
}
type patientHandler struct {
	PatientRepository repository.PatientRepository
	EMRRepository     repository.EMRRepository
	SymptomRepository repository.SymptomRepository
}

func NewPatientHandler(patientRepository repository.PatientRepository, emrRepository repository.EMRRepository, symptomRepository repository.SymptomRepository) PatientHandler {
	return &patientHandler{
		PatientRepository: patientRepository,
		EMRRepository:     emrRepository,
		SymptomRepository: symptomRepository,
	}
}

func (handler *patientHandler) GetAllPatients(c *fiber.Ctx) error {
	results, _ := handler.PatientRepository.FindAllPatients(c.Context())

	var emrs []domain.EMR
	for _, result := range results {
		emr, _ := handler.EMRRepository.FindOneEMR(c.Context(), "patients.nik", result.NIK)
		emrs = append(emrs, emr)
	}

	patients := web.NewAllPatientResponse(emrs)

	return c.Render("pages-data-pasien", fiber.Map{
		"Title":    "Data Pasien",
		"Patients": patients,
	})
}

func (handler *patientHandler) InputPatient(c *fiber.Ctx) error {
	symptoms, err := handler.SymptomRepository.FindAllSymtoms(c.Context())
	if err != nil {
		return exception.ErrInternalServer(err.Error())
	}
	return c.Render("pages-input-pasien", fiber.Map{
		"Title":    "Input Pasien",
		"Symptoms": symptoms,
	})
}
