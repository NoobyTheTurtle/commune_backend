package db

import (
	"commune_backend/internal/app/models"
	"math/rand"
	"time"
)

func mockOffices(offices []*models.Office) {
	rand.Seed(time.Now().UnixNano())
	for i := range offices {
		offices[i].Withdrawal = rand.Float64() < 0.8
		offices[i].Replenishment = rand.Float64() < 0.8

		// Random services
		numServices := rand.Intn(len(models.ServicesList)) + 1
		perm := rand.Perm(len(models.ServicesList))
		offices[i].Services = make(models.ServiceList, numServices)
		for j := 0; j < numServices; j++ {
			offices[i].Services[j] = models.ServicesList[perm[j]]
		}
	}
}

func mockAtms(atms []*models.Atm) {
	rand.Seed(time.Now().UnixNano())
	for i := range atms {
		atms[i].Withdrawal = rand.Float64() < 0.8
		atms[i].Replenishment = rand.Float64() < 0.8
	}
}
