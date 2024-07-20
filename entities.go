package main

type Meeting struct {
	ID     uint    `gorm:"primaryKey;index:,sort:desc"`
	Title  string  `gorm:"size:100;not null"`
	Images []Image `gorm:"foreignKey:MeetingID"`
	Videos []Video `gorm:"foreignKey:MeetingID"`
	Musics []Music `gorm:"foreignKey:MeetingID"`
}

type Image struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"size:255;not null"`
	MeetingID uint
}

type Video struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"size:255;not null"`
	MeetingID uint
}

type Music struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"size:255;not null"`
	Title     string `gorm:"size:100;null"`
	MeetingID uint
}
