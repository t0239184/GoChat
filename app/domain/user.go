package domain

type User struct {
	Id       int64  `gorm:"column_name: id type: bigint(20) not null auto_increment; primary_key" json:"id"`
	Account  string `gorm:"column_name: account type: varchar(50) not null; index:idx_account,unique" json:"account"`
	Password string `gorm:"column_name: password type: varchar(50) not null" json:"-"`
	SaltId   int64  `gorm:"column_name: salt_id type: bigint(20) not null;" json:"salt_id"`
	Status   string `gorm:"column_name: statustype: varchar(1) not null default '0';" json:"status"`
}

func (User) TableName() string {
	return "user"
}

type Salt struct {
	Id          int64  `gorm:"column_name: id type: bigint(20) not null auto_increment; primary_key" json:"id"`
	Salt        string `gorm:"column_name: salt type: varchar(100) not null;" json:"salt"`
	Iteration   int16  `gorm:"column_name: iteration type: smallint(4) not null;" json:"iteration"`
}

func (Salt) TableName() string {
	return "salt"
}


type IUserRepository interface {
	CreateUser(user *User, salt *Salt) (id *int64, error error)
}

type IUserUsecase interface {
	CreateUser(user *User) (id *int64, error error)
}
