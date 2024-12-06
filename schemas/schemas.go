package schemas

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP"`
}

type Profile struct {
	ID              uint `gorm:"primary_key"`
	UserID          uint `gorm:"not null"`
	GamePreference  string
	PlayStyle       string
	Region          string
	ExperienceLevel string
	Bio             string
	UpdatedAt       string `gorm:"default:CURRENT_TIMESTAMP"`
}

type Dupla struct {
	ID        uint   `gorm:"primary_key"`
	Player1ID uint   `gorm:"not null"`
	Player2ID uint   `gorm:"not null"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP"`
}

type Match struct {
	ID           uint `gorm:"primary_key"`
	DuplaID      uint `gorm:"not null"`
	GameMode     string
	MatchDate    string
	Result       string
	Player1Score int
	Player2Score int
	UpdatedAt    string `gorm:"default:CURRENT_TIMESTAMP"`
}
