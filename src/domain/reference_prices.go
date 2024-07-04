package domain

// TODO: 複数か単数か
// TODO: IDいるかどうか問題(一旦いらない)
// TODO: 外部キー制約についてfundID
type ReferencePrices struct {
	FundID             string
	ReferencePrice     int
	ReferencePriceDate string
}

func NewReferencePrice(fundID string, referencePrice int, referencePriceDate string) *ReferencePrices {
	return &ReferencePrices{
		FundID:             fundID,
		ReferencePrice:     referencePrice,
		ReferencePriceDate: referencePriceDate,
	}
}
