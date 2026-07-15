package model

import (
	"time"

	"github.com/google/uuid"
)

type Branch struct {
	ID   uuid.UUID
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Department struct {
	ID       uuid.UUID
	Name     string
	ParentID *uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Position struct {
	ID   uuid.UUID
	Name string
}

type Permission struct {
	ID     uuid.UUID
	Code   string
	Name   string
	Module string
}

type Role struct {
	ID          uuid.UUID
	Name        string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Staff struct {
	ID           uuid.UUID
	EmployeeCode string

	FirstName string
	LastName  string

	Username     string
	PasswordHash string

	BranchID     *uuid.UUID
	DepartmentID *uuid.UUID
	PositionID   *uuid.UUID

	Status string

	CreatedBy *uuid.UUID
	UpdatedBy *uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Session struct {
	ID uuid.UUID

	StaffID uuid.UUID

	RefreshToken string
	staffAgent   string
	IPAddress    string

	ExpiredAt time.Time
	CreatedAt time.Time
}

type ActivityLog struct {
	ID uuid.UUID

	StaffID *uuid.UUID

	Action   string
	Entity   string
	EntityID *uuid.UUID

	OldData []byte
	NewData []byte

	CreatedAt time.Time
}

type StaffList struct {
	ID             uuid.UUID
	EmployeeCode   string
	FirstName      string
	LastName       string
	Username       string
	BranchName     string
	DepartmentName string
	PositionName   string
	Status         string
}
