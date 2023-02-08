package model

import "time"

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 10:12
 */

type Base struct {
	ID         uint      `gorm:"column:id;primary_key"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	Status     bool      `gorm:"column:status;default:1"`
}
