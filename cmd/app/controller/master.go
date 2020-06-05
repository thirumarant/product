package controller

import "github.com/jinzhu/gorm"

type Master struct {
	DebugOn bool
	DB      *gorm.DB
}

func (m *Master) Debug(debugMode bool) {
}

func (m *Master) Add(controller *Master) (err error) {
	return err
}

func (m *Master) Delete(controller *Master) (err error) {
	return err
}

func (m *Master) Update(controller *Master) (err error) {
	return err
}
