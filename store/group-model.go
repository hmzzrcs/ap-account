package store

import "time"

type UserGroup struct {
	Id         int64     `gorm:"type:BIGINT(20);size:20;not null;auto_increment;primary_key;column:id;comment:主键ID;"`
	UUID       string    `gorm:"size:36;not null;column:uuid;uniqueIndex:gid;comment:分组id;uniqueIndex:gid_uid;"`
	Name       string    `gorm:"size:36;not null;column:name;uniqueIndex:name;comment:分组名称;uniqueIndex:gid_uid;"`
	Creator    string    `gorm:"size:36;not null;column:creator;uniqueIndex:creator;comment:创建人;uniqueIndex:gid_uid;"`
	CreateTime time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP;column:create_at;comment:创建时间"`
}

func (u *UserGroup) TableName() string {
	return "user_group"
}
func (u *UserGroup) IdValue() int64 {
	return u.Id
}
