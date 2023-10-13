package sub_models

import (
	"commune_backend/internal/app/models"
	"gorm.io/gorm"
)

type AtmWithDistance struct {
	models.Atm
	RadiusDistance float64 `json:"radius_distance"`
}

type AtmsWithDistance []AtmWithDistance

func (o *AtmsWithDistance) GetWithinRadius(db *gorm.DB, userLat float64, userLng float64, radius float64) error {
	sqlQuery := `		
		SELECT
			distances.*
		FROM (
			SELECT
				atms.*,
				(
					6371000 *
						acos(cos(radians(?)) *
							 cos(radians(atms.latitude)) *
							 cos(radians(atms.longitude) -
								 radians(?)) +
							 sin(radians(?)) *
							 sin(radians(atms.latitude)))
				) AS radius_distance
			FROM atms
		) as distances
		
		WHERE distances.radius_distance < ?
		ORDER BY distances.radius_distance
		LIMIT 40;
	`

	err := db.Raw(sqlQuery, userLat, userLng, userLat, radius).Scan(o).Error
	return err
}
