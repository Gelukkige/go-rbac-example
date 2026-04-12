package model

type Data struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255)"`
	Tag  string `gorm:"type:varchar(255)"`
	Desc string `gorm:"type:text"`
}

type DataResp struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
	Desc string `json:"desc,omitempty"`
}
