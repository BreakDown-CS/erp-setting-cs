package repository

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/model"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/setup/posts"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ports.SetupRepository {
	return &repository{db: db}
}

func (r *repository) FindAllBranches() ([]model.Branches, error) {
	// รูปแบบที่ 1: ใช้ GORM Query Builder
	var result []model.Branches
	err := r.db.Table("erp.branches p").
		Select("p.id, p.name").
		Where("p.mysql_shop_id = ?", 1). // สมมติค่า parameter
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	// รูปแบบที่ 2: ใช้ GORM ในการยิง Raw SQL (เหมือน pgx)
	// var user model.Branches
	// err = r.db.Raw(`
	// 	SELECT s.id, s.name
	// 	FROM erp.staffs s
	// 	WHERE s.username = ?`, username).Scan(&user).Error

	return result, nil
}
