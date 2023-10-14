package support_models

import (
	"commune_backend/internal/app/filters"
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models"
	"gorm.io/gorm"
)

type AtmWithRadius struct {
	models.Atm
	RadiusDistance float64 `json:"radius_distance"`
}

type AtmsWithRadius []AtmWithRadius

func (o *AtmsWithRadius) GetWithinRadius(
	db *gorm.DB,
	ga *params.GeoArea,
	f *filters.Filters,
) error {
	sqlSubQuery := db.Select(`		
		atms.id,
		(
			6371000 *
				acos(cos(radians(?)) *
					 cos(radians(atms.latitude)) *
					 cos(radians(atms.longitude) -
						 radians(?)) +
					 sin(radians(?)) *
					 sin(radians(atms.latitude)))
		) AS radius_distance
	`, ga.UserLat, ga.UserLng, ga.UserLat).Table("atms")

	err := db.Select("atms.*, sub.radius_distance").
		Table("atms").
		Joins("INNER JOIN (?) AS sub ON sub.id = atms.id", sqlSubQuery).
		Where("sub.radius_distance <= ?", ga.Radius).
		Scopes(f.Scopes()...).
		Order("sub.radius_distance").
		Limit(20).
		Find(o).
		Error
	return err
}
