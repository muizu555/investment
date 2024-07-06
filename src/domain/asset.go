package domain

type Asset struct {
	Data         string `json:"data"`
	CurrentValue int    `json:"current_value"`
	CurrentPL    int    `json:"current_pl"`
}

type Assets []Asset

type AssetSetting struct {
	AppraisedAsset   int
	PurchasePriceSum int
	ProfitLoss       int
}

type AssetSettings []AssetSetting

type AssetYear struct {
	Year         int `json:"year"`
	CurrentValue int `json:"current_value"`
	CurrentPL    int `json:"current_pl"`
}

type AssetYears []AssetYear

type AssetResponse struct {
	Date   string     `json:"date"`
	Assets AssetYears `json:"assets"`
}

// 一年ごとのAssetSetting(Yearを最初に追加)
type AssetYearSetting struct {
	TradeYear        string
	AppraisedAsset   int
	PurchasePriceSum int
	ProfitLoss       int
}

type AssetYearSettings []AssetYearSetting
