package model

import (
	_ "github.com/joho/godotenv"
	"gorm.io/gorm"
)

///////////////////////////DATABASE MIGRATION MODELS//////////////////////////////////////////////////////

type User struct {
	gorm.Model          //gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Email       *string `gorm:"column:email;type:varchar(256);unique;not null" json:"email"`
	FirstName   string  `gorm:"column:first_name;type:varchar(256);not null" json:"first_name"`
	LastName    string  `gorm:"column:last_name;type:varchar(256);not null" json:"last_name"`
	AccessToken string  `gorm:"column:access_token;type:varchar(256);not null" json:"access_token"`
}

type Track struct {
	gorm.Model //gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	//ID         int64          `gorm:"column:id;type:bigint(20) unsigned;primaryKey" json:"id"`
	UserId     string
	User       User   `gorm:"foreignKey:UserId"` //relationship
	TrackingId string `gorm:"column:tracking_id;type:varchar(256);not null" json:"tracking_id"`
	// CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	// UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	// DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////

type Location struct {
	gorm.Model
	Longitude string `gorm:"column:longitude;type:varchar(256);not null" json:"longitude" binding:"required"`
	Latitude  string `gorm:"column:latitude;type:varchar(256);not null" json:"latitude" binding:"required"`
	//Trackingid  string `gorm:"column:tracking_id;type:varchar(256);not null" json:"tracking_id" binding:"required"`
	//Accesstoken string `gorm:"column:access_token;type:varchar(256);not null" json:"access_token" binding:"required"`

	AccessToken string
	User        User `gorm:"foreignKey:AccessToken"` //relationship

	TrackingId string //relationship
	Track      Track  `gorm:"foreignKey:TrackingId"`
}

type DeviceInfo struct {
	gorm.Model
	Useragent                string `json:"user_agent" binding:"required"`
	Platform                 string `json:"platform" binding:"required"`
	Cpu                      string `json:"cpu" binding:"required"`
	Ram                      string `json:"ram" binding:"required"`
	Gpuingo                  string `json:"gpuinfo" binding:"required"`
	Screen_resolution_height uint32 `json:"screen_resolution_height" binding:"required"`
	Screen_resolution_width  uint32 `json:"screen_resolution_width" binding:"required"`
	Os                       string `json:"os" binding:"required"`
	TouchSupport             string `json:"touch_support" binding:"required"`
	CookieEnabled            bool   `json:"cookie_enabled" binding:"required"`
	DoNotTrack               bool   `json:"donot_track" binding:"required"`

	AccessToken string
	User        User `gorm:"foreignKey:AccessToken"` //relationship

	TrackingId string //relationship
	Track      Track  `gorm:"foreignKey:TrackingId"`
}
