package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"`
	ResourceID string    `valid:"notnull" gorm:"type:varchar(255);notnull"`
	FilePath   string    `valid:"notnull" gorm:"type:varchar(255);notnull"`
	CreatedAt  time.Time `valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}

	return nil
}
