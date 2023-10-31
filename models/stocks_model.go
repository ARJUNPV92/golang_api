package models

import (
    "time"
)

type Stock struct {
    ID               uint      `gorm:"primary_key" json:"id"`
    TIMESTAMP        time.Time `json:"TIMESTAMP"`
    ScCode           float64   `json:"ScCode"`
    ScName           string    `json:"ScName"`
    SYMBOL           string    `json:"SYMBOL"`
    OPEN             float64   `json:"OPEN"`
    HIGH             float64   `json:"HIGH"`
    LOW              float64   `json:"LOW"`
    CLOSE            float64   `json:"CLOSE"`
    Last             float64   `json:"Last"`
    PREVCLOSE        float64   `json:"PREVCLOSE"`
    SERIES           string    `json:"SERIES"`
    BSENSETotalTrades float64   `json:"BSE+NSE_TotalTrades"`
    BSENSEVol        float64   `json:"BSE+NSE_Vol"`
    BSENSEDeliVol    float64   `json:"BSE+NSE_DeliVol"`
    BSENSEDeli       float64   `json:"BSE+NSE_Deli%"`
    BSENSETurnover   float64   `json:"BSE+NSE_TurnoverRsCr"`
    BSENSEDeliVal    float64   `json:"BSE+NSE_DeliValRsCr"`
    BSENSEPerOrder   float64   `json:"BSE+NSE_PerOrderWorhRs"`
    BSENSEPerOrderQty float64   `json:"BSE+NSE_PerOrderQty"`
    BSENSEAvgPrice    float64   `json:"BSE+NSE_AvgPrice"`
    ISIN              string    `json:"ISIN"`
}

func (Stock) TableName() string {
    return "tbl_BSE_NSE_Historic_Data"
}
