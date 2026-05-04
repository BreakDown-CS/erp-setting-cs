package service

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/staffs/posts"
)

type service struct {
	repo ports.StaffRepository
}

var _ ports.StaffService = (*service)(nil)

func NewUsecase(r ports.StaffRepository) ports.StaffService {
	return &service{repo: r}
}

func (u *service) InsertStaff(req dto.CreateStaffRequest) error {
	staff := model.Staff{
		EmployeeCode: req.Name,
	}
	return u.repo.InsertStaff(staff)
}

// func (u *service) GetAllstaffs(page int, limit int) ([]dto.StaffListResponse, int, error) {

// 	offset := (page - 1) * limit

// 	staffs, total, err := u.repo.GetAllStaff(limit, offset)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	res := make([]dto.StaffListResponse, 0)

// 	for _, v := range staffs {
// 		res = append(res, mapper.ToStaffList(v))
// 	}

// 	return res, total, nil
// }

// func (u *service) GetstaffById(id int) (dto.StaffDetailResponse, error) {
// 	staff, err := u.repo.GetstaffById(id)
// 	if err != nil {
// 		return dto.StaffDetailResponse{}, err
// 	}

// 	return mapper.ToStaffDetail(staff), nil
// }
