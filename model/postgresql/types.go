package postgresql

import (
	"gorm.io/gorm"
	"time"
)

// User is the PostgreSQL model for the user
type User struct {
	gorm.Model
	UserIDHex string `json:"user_id_hex" gorm:"uniqueIndex"`
}

// OrderPayments is the PostgreSQL model for the order payments
type OrderPayments struct {
	gorm.Model
	OrderID uint64 `json:"order_id" gorm:"uniqueIndex"`
}

// OrderPayment is the PostgreSQL model for the order payment
type OrderPayment struct {
	gorm.Model
	OrderPaymentsID uint64        `json:"order_payments_id"`
	OrderPayments   OrderPayments `gorm:"foreignKey:OrderPaymentsID"`
	PaymentID       uint64        `json:"payment_id" gorm:"uniqueIndex"`
	Payment         Payment       `gorm:"foreignKey:PaymentID"`
}

// BranchRentPayments is the PostgreSQL model for the branch rent payments
type BranchRentPayments struct {
	gorm.Model
	BranchRentIDHex string `json:"branch_rent_id_hex" gorm:"uniqueIndex"`
}

// BranchRentPayment is the PostgreSQL model for the branch rent payment
type BranchRentPayment struct {
	gorm.Model
	BranchRentPaymentsID uint64             `json:"branch_rent_payments_id"`
	BranchRentPayments   BranchRentPayments `gorm:"foreignKey:BranchRentPaymentsID"`
	PaymentID            uint64             `json:"payment_id" gorm:"uniqueIndex"`
	Payment              Payment            `gorm:"foreignKey:PaymentID"`
}

// Payment is the PostgreSQL model for the payment
type Payment struct {
	gorm.Model
	PaymentAccountID  uint64         `json:"payment_account_id"`
	PaymentAccount    PaymentAccount `gorm:"foreignKey:PaymentAccountID"`
	PaymentIdentifier string         `json:"payment_identifier" gorm:"uniqueIndex"`
	Amount            float64        `json:"amount"`
	PaymentDate       time.Time      `json:"payment_date"`
	RefundedAt        *time.Time     `json:"refunded_at,omitempty"`
	VerifiedByUserID  *uint64        `json:"verified_by_user_id"`
	VerifiedByUser    User           `gorm:"foreignKey:VerifiedByUserID"`
}

// PaymentAccount is the PostgreSQL model for the payment account
type PaymentAccount struct {
	gorm.Model
	BranchIDHex       *string    `json:"branch_id_hex,omitempty"`
	Platform          string     `json:"platform"`
	AccountIdentifier string     `json:"account_identifier" gorm:"uniqueIndex"`
	Owner             *string    `json:"owner,omitempty"`
	Email             *string    `json:"email,omitempty"`
	PhoneNumber       *string    `json:"phone_number,omitempty"`
	SuspendedAt       *time.Time `json:"suspended_at,omitempty"`
}
