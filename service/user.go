package service

import "github.com/makki0205/vuejs/model"

var User = userImpl{}

type userImpl struct {
}

func (self userImpl) Create(user model.User) model.User {
	err := db.Create(&user).Error
	if err != nil {
		panic(err)
	}
	return user
}

func (self userImpl) Find(id uint) *model.User {
	user := []model.User{}
	db.Find(&user, id)
	if len(user) == 0 {
		return nil
	}
	return &user[0]
}
func (self userImpl) Update(id uint, user model.User) *model.User {
	user.ID = id
	db.Model(&user).Update(&User)
	return &user
}

func (self userImpl) Delete(id uint) {
	user := model.User{}
	db.Delete(&user, id)
}

func (self userImpl) All() []model.User {
	var user []model.User
	db.Find(&user)
	return user
}

func (self userImpl) Exists(id uint) bool {
	return self.Find(id) != nil
}
func (self userImpl) FindByEmail(email string) *model.User {
	return self.FindByOne("email = ?", email)
}
func (self userImpl) FindBy(query interface{}, args ...interface{}) []model.User {
	user := []model.User{}
	db.Where(query, args...).Find(&user)
	if len(user) == 0 {
		return nil
	}
	return user
}

func (self userImpl) FindByOne(query interface{}, args ...interface{}) *model.User {
	user := []model.User{}
	db.Where(query, args...).Find(&user)
	if len(user) == 0 {
		return nil
	}
	return &user[0]
}
