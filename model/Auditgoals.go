package model

import "time"

type Auditgoals struct {
	goal_id    uint `gorm:"column:goal_id;type:int;not null;primary_key;auto_increment"`
	goal_text  string
	user_id    uint
	start_time time.Time
	end_time   time.Time
}
