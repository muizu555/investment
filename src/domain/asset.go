package domain

type Asset struct {
	Data         string
	CurrentValue int
	CurrentPL    int
}

type Assets []Asset

type AssetSetting struct {
	AppraisedAsset   int
	PurchasePriceSum int
	ProfitLoss       int
}

type AssetSettings []AssetSetting
