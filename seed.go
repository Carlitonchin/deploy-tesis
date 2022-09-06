package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/service"
	"gorm.io/gorm"
)

func seed(db *gorm.DB) {
	var role *model.Role

	err := db.First(&role).Error

	if err != nil {
		seedRoles(db)
		seedAreas(db)
		seedUsers(db)
		seedStatus(db)
		seedQuestions(db)
	}
}

func seedRoles(db *gorm.DB) {

	db.AutoMigrate(&model.Role{})

	erra := db.Create(&model.Role{Name: os.Getenv("ROLE_ADMIN")}).Error
	errc := db.Create(&model.Role{Name: os.Getenv("ROLE_DEFAULT_WORKER")}).Error
	erre := db.Create(&model.Role{Name: os.Getenv("ROLE_DEFAULT_STUDENT")}).Error
	errs1 := db.Create(&model.Role{Name: os.Getenv("ROLE_SPECIALIST_LEVEL1")}).Error
	errs2 := db.Create(&model.Role{Name: os.Getenv("ROLE_SPECIALIST_LEVEL2")}).Error

	if atLeastOneError(erra, errc, erre, errs1, errs2) {
		log.Fatal("Error at seed data for roles")
	}
}

func seedUsers(db *gorm.DB) {

	db.AutoMigrate(&model.User{})
	var role *model.Role

	errf := db.First(&role, "name = ?", os.Getenv("ROLE_ADMIN")).Error
	hashed_pass, errp := service.HashPass("administrador")

	erru := db.Create(&model.User{
		Email:    "admin@admin.com",
		Name:     "admin",
		Password: hashed_pass,
		RoleID:   role.ID,
	}).Error

	if atLeastOneError(errf, errp, erru) {
		log.Fatal("Error seeding users at db")
	}

}

func seedStatus(db *gorm.DB) {
	db.AutoMigrate(&model.Status{})
	var status []model.Status

	statusSend := model.Status{Description: os.Getenv("STATUS_SEND")}
	statusSendId, err1 := strconv.ParseUint(os.Getenv("STATUS_SEND_CODE"), 10, 64)
	statusSend.ID = uint(statusSendId)

	statusCla1 := model.Status{Description: os.Getenv("STATUS_CLASIFIED1")}
	statusCla1Id, err2 := strconv.ParseUint(os.Getenv("STATUS_CLASIFIED1_CODE"), 10, 64)
	statusCla1.ID = uint(statusCla1Id)

	statusCla2 := model.Status{Description: os.Getenv("STATUS_CLASIFIED2")}
	statusCla2Id, err3 := strconv.ParseUint(os.Getenv("STATUS_CLASIFIED2_CODE"), 10, 64)
	statusCla2.ID = uint(statusCla2Id)

	statusAdmin := model.Status{Description: os.Getenv("STATUS_ADMIN")}
	statusAdminId, err4 := strconv.ParseUint(os.Getenv("STATUS_ADMIN_CODE"), 10, 64)
	statusAdmin.ID = uint(statusAdminId)

	statusFinish := model.Status{Description: os.Getenv("STATUS_FINISH")}
	statusFinishId, err5 := strconv.ParseUint(os.Getenv("STATUS_FINISH_CODE"), 10, 64)
	statusFinish.ID = uint(statusFinishId)

	status = append(status, statusSend)
	status = append(status, statusCla1)
	status = append(status, statusCla2)
	status = append(status, statusAdmin)
	status = append(status, statusFinish)

	err6 := db.Create(&status).Error

	if atLeastOneError(err1, err2, err3, err4, err5, err6) {
		log.Fatal("Error seeding status")
	}
}

func seedQuestions(db *gorm.DB) {
	db.AutoMigrate(&model.Question{})
}

func seedAreas(db *gorm.DB) {
	db.AutoMigrate(&model.Area{})
}

func atLeastOneError(errors ...error) bool {
	for _, err := range errors {
		if err != nil {
			return true
		}
	}

	return false
}
