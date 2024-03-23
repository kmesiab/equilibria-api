package models

import "time"

const (
	AccountStatusPendingActivation = 1
	AccountStatusActive            = 2
	AccountStatusSuspended         = 3
	AccountStatusExpired           = 4
)

func StringToAccountStatus(status string) AccountStatus {

	switch status {
	case "Pending Activation":

		return AccountStatus{
			ID:   AccountStatusPendingActivation,
			Name: "Pending Activation",
		}

	case "Active":

		return AccountStatus{
			ID:   AccountStatusActive,
			Name: "Active",
		}

	case "Suspended":
		return AccountStatus{
			ID:   AccountStatusSuspended,
			Name: "Suspended",
		}

	case "Expired":
		return AccountStatus{
			ID:   AccountStatusExpired,
			Name: "Expired",
		}
	}

	return AccountStatus{
		ID:   AccountStatusPendingActivation,
		Name: "Pending Activation",
	}
}

// AccountStatus represents the account_statuses table in the database.
type AccountStatus struct {
	ID        int64     `json:"status_id" gorm:"primaryKey; autoIncrement"`
	Name      string    `gorm:"size:100; unique; not null" json:"status_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func (a *AccountStatus) TableName() string {
	return "account_statuses"
}
