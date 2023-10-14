package support_models

import (
	"commune_backend/internal/app/filters"
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models"
	"gorm.io/gorm"
)

type OfficeWithRadius struct {
	models.Office
	RadiusDistance float64 `json:"radius_distance"`
}

type OfficesWithRadius []OfficeWithRadius

func (o *OfficesWithRadius) GetWithinRadius(
	db *gorm.DB,
	ga *params.GeoArea,
	f *filters.Filters,
) error {
	sqlSubQuery := db.Select(`		
		offices.id,
		(
			6371000 *
				acos(cos(radians(?)) *
					 cos(radians(offices.latitude)) *
					 cos(radians(offices.longitude) -
						 radians(?)) +
					 sin(radians(?)) *
					 sin(radians(offices.latitude)))
		) AS radius_distance
	`, ga.UserLat, ga.UserLng, ga.UserLat).Table("offices")

	err := db.Select("offices.*, sub.radius_distance").
		Table("offices").
		Joins("INNER JOIN (?) AS sub ON sub.id = offices.id", sqlSubQuery).
		Scopes(f.Scopes()...).
		Where("sub.radius_distance <= ?", ga.Radius).
		Order("sub.radius_distance").
		Limit(20).
		Find(o).
		Error
	return err
}
