package categorymodel

import data "scrap/Model/Data_Model"

type Category struct {
	ID            int         `json:"id"`
	Category_Name string      `gorm:"unique:category_name"`
	Datas         []data.Data `gorm:"FooreignKey:Category_ID"`
}
