package sub_models

import (
	"commune_backend/internal/app/models"
	"gorm.io/gorm"
)

type OfficeWithDistance struct {
	models.Office
	RadiusDistance float64 `json:"radius_distance"`
}

type OfficesWithDistance []OfficeWithDistance

func (o *OfficesWithDistance) GetWithinRadius(db *gorm.DB, userLat float64, userLng float64, radius float64) error {
	sqlQuery := `		
		SELECT
			distances.*
		FROM (
			SELECT
				offices.*,
				(
					6371000 *
						acos(cos(radians(?)) *
							 cos(radians(offices.latitude)) *
							 cos(radians(offices.longitude) -
								 radians(?)) +
							 sin(radians(?)) *
							 sin(radians(offices.latitude)))
				) AS radius_distance
			FROM offices
		) as distances
		
		WHERE distances.radius_distance < ?
		ORDER BY distances.radius_distance
		LIMIT 20;
	`

	err := db.Raw(sqlQuery, userLat, userLng, userLat, radius).Scan(o).Error
	return err
}
