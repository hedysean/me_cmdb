package model

import "time"

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 10:12
 */

type Base struct {
	ID         uint `gorm:"primary_key"`
	CreateTime time.Time
	UpdateTime time.Time
	Status     bool `gorm:"default:1"`
}
