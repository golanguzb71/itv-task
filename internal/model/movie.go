package model

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Director    string         `gorm:"type:varchar(255);not null" json:"director"`
	Year        int            `gorm:"not null" json:"year"`
	Plot        string         `gorm:"type:text" json:"plot"`
	Rating      float32        `gorm:"type:decimal(3,1)" json:"rating"`
	Duration    int            `gorm:"not null" json:"duration"` // in minutes
	ReleaseDate string         `json:"release_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	Role      string         `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (m *Movie) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Movie) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
