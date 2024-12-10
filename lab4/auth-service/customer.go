package main

type Customer struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	BirthDate    string `json:"birth_date"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash"`
}
