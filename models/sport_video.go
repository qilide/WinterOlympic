package models

type SportVideo struct {
	ID        int    `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	Describe1 string `json:"describe1" gorm:"column:describe1"`
	Video1    string `json:"video1" gorm:"column:video1"`
	Img1      string `json:"img1" gorm:"column:img1"`
	Describe2 string `json:"describe2" gorm:"column:describe2"`
	Video2    string `json:"video2" gorm:"column:video2"`
	Img2      string `json:"img2" gorm:"column:img2"`
	Describe3 string `json:"describe3" gorm:"column:describe3"`
	Video3    string `json:"video3" gorm:"column:video3"`
	Img3      string `json:"img3" gorm:"column:img3"`
	Year      int    `json:"year" gorm:"column:year"`
}
