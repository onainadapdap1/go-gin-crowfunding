package user

import "gorm.io/gorm"

// set interface
type Repository interface{
	Save(user User) (User, error)
}

type repository struct{
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{DB: db}
}

func (r *repository) Save(user User) (User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return user, nil
	}

	return user, nil
}
