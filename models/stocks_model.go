package models

import (
	"time"
)

type Stock struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	TIMESTAMP        time.Time `gorm:"column:TIMESTAMP" json:"TIMESTAMP"`
	ScCode           float64   `gorm:"column:ScCode" json:"ScCode"`
	ScName           string    `gorm:"column:ScName" json:"ScName"`
	SYMBOL           string    `gorm:"column:SYMBOL" json:"SYMBOL"`
	OPEN             float64   `gorm:"column:OPEN" json:"OPEN"`
	HIGH             float64   `gorm:"column:HIGH" json:"HIGH"`
	LOW              float64   `gorm:"column:LOW" json:"LOW"`
	CLOSE            float64   `gorm:"column:CLOSE" json:"CLOSE"`
	Last             float64   `gorm:"column:Last" json:"Last"`
	PREVCLOSE        float64   `gorm:"column:PREVCLOSE" json:"PREVCLOSE"`
	SERIES           string    `gorm:"column:SERIES" json:"SERIES"`
	BSENSETotalTrades float64 `gorm:"column:BSE+NSE_TotalTrades" json:"BSE+NSE_TotalTrades"`
	BSENSEVol        float64   `gorm:"column:BSE+NSE_Vol" json:"BSE+NSE_Vol"`
	BSENSEDeliVol    float64   `gorm:"column:BSE+NSE_DeliVol" json:"BSE+NSE_DeliVol"`
	BSENSEDeli       float64   `gorm:"column:BSE+NSE_Deli%" json:"BSE+NSE_Deli%"`
	BSENSETurnover   float64   `gorm:"column:BSE+NSE_TurnoverRsCr" json:"BSE+NSE_TurnoverRsCr"`
	BSENSEDeliVal    float64   `gorm:"column:BSE+NSE_DeliValRsCr" json:"BSE+NSE_DeliValRsCr"`
	BSENSEPerOrder   float64   `gorm:"column:BSE+NSE_PerOrderWorhRs" json:"BSE+NSE_PerOrderWorhRs"`
	BSENSEPerOrderQty float64 `gorm:"column:BSE+NSE_PerOrderQty" json:"BSE+NSE_PerOrderQty"`
	BSENSEAvgPrice    float64 `gorm:"column:BSE+NSE_AvgPrice" json:"BSE+NSE_AvgPrice"`
	ISIN              string    `gorm:"column:ISIN" json:"ISIN"`
}

func (Stock) TableName() string {
	return "tbl_BSE_NSE_Historic_Data"
}
