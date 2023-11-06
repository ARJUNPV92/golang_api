package models

import (
    "time"
)

type Stock struct {
    ID               uint      `gorm:"primary_key" json:"id"`
    Timestamp        time.Time `json:"TIMESTAMP"`
    ScCode           float64   `json:"ScCode"`
    ScName           string    `json:"ScName"`
    Symbol           string    `json:"SYMBOL"`
    Open             float64   `json:"OPEN"`
    High             float64   `json:"HIGH"`
    Low              float64   `json:"LOW"`
    Close            float64   `json:"CLOSE"`
    Last             float64   `json:"Last"`
    PrevClose        float64   `json:"PREVCLOSE"`
    Series           string    `json:"SERIES"`
    BSE_NSE_TotalTrades float64   `json:"BSE+NSE_TotalTrades"`
    BSE_NSE_Vol        float64   `json:"BSE+NSE_Vol"`
    BSE_NSE_DeliVol    float64   `json:"BSE+NSE_DeliVol"`
    BSE_NSE_Deli       float64   `json:"BSE+NSE_Deli%"`
    BSE_NSE_Turnover   float64   `json:"BSE+NSE_TurnoverRsCr"`
    BSE_NSE_DeliVal    float64   `json:"BSE+NSE_DeliValRsCr"`
    BSE_NSE_PerOrder   float64   `json:"BSE+NSE_PerOrderWorhRs"`
    BSE_NSE_PerOrderQty float64   `json:"BSE+NSE_PerOrderQty"`
    BSE_NSE_AvgPrice    float64   `json:"BSE+NSE_AvgPrice"`
    ISIN              string    `json:"ISIN"`
}

func (Stock) TableName() string {
    return "tbl_BSE_NSE_Historic_Data"
}
