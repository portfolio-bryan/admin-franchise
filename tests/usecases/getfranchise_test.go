package usecases_test

import (
	"context"
	"fmt"

	"github.com/bperezgo/admin_franchise/graph"
	"github.com/bperezgo/admin_franchise/graph/model"
	"github.com/bperezgo/admin_franchise/internal/platform/repositories/postgres"
	utilstests "github.com/bperezgo/admin_franchise/tests/utils"
	"github.com/google/uuid"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetFranchise", func() {
	utilstests.LoadEnv()
	docker := utilstests.UpDocker()
	conn := utilstests.Connect(docker.Host, docker.Port).PostgresDB
	utilstests.UpMigrations(docker.Host, docker.Port)

	var locModel postgres.LocationModel
	var addModel postgres.AddressLocationModel
	var comModel postgres.CompanyModel
	var franModel postgres.FranchiseModel

	BeforeEach(func() {
		if err := conn.Exec("DELETE FROM franchise; DELETE FROM company; DELETE FROM address_location; DELETE FROM locations").
			Error; err != nil {
			panic(fmt.Errorf("failed to delete data from tables. %w", err))
		}

		locID := uuid.NewString()
		addID := uuid.NewString()
		comID := uuid.NewString()
		franID := uuid.NewString()

		locModel = postgres.LocationModel{
			ID:      locID,
			City:    "city",
			Country: "country",
			State:   "state",
		}
		addModel = postgres.AddressLocationModel{
			ID:         addID,
			LocationID: locID,
			Address:    "address",
			ZipCode:    "zipcode",
		}
		comModel = postgres.CompanyModel{
			ID:                comID,
			Name:              "company name",
			CompanyOwnerID:    uuid.NewString(),
			TaxNumber:         "tax number",
			LocationID:        locID,
			AddressLocationID: addID,
		}
		franModel = postgres.FranchiseModel{
			ID:                   franID,
			CompanyID:            comID,
			Title:                "franchise title",
			SiteName:             "site name",
			Description:          "franchise description",
			Image:                "marriott.com/content/dam/marriott-homepage/homepageassets/destination-theme-beach-6336-ver-clsc.jpg",
			URL:                  "https://www.marriott.com",
			Protocol:             "https",
			DomainJumps:          1,
			ServerNames:          []string{"marriott.com"},
			DomainCreationDate:   "2021-01-01",
			DomainExpirationDate: "2021-01-01",
			RegistrantName:       "registrant name",
			RegistrantEmail:      "registrant@email.com",
			LocationID:           locID,
			AddressLocationID:    addID,
		}

		conn.Create(&locModel)
		conn.Create(&addModel)
		conn.Create(&comModel)
		conn.Create(&franModel)
	})

	AfterEach(func() {
		if err := conn.Exec("DELETE FROM franchise; DELETE FROM company; DELETE FROM address_location; DELETE FROM locations").
			Error; err != nil {
			panic(fmt.Errorf("failed to delete data from tables. %w", err))
		}
	})

	Describe("GetByName", func() {
		resolver := graph.NewResolver()

		ctx := context.Background()
		criteria := &model.FranchiseCriteria{
			Name: &franModel.SiteName,
		}

		Context("should get a franchise", func() {
			It("franchiseID should be equal to franModel.ID", func() {
				franResp, err := resolver.Query().GetFranchise(ctx, criteria)

				Expect(err).To(BeNil())
				Expect(franResp.ID).To(Equal(franModel.ID))
			})
		})
	})
})
