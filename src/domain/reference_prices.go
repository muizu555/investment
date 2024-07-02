package domain

// TODO: 複数か単数か
// TODO: IDいるかどうか問題(一旦いらない)
// TODO: 外部キー制約についてfundID
type ReferencePrice struct {
	FundID             string
	ReferencePriceDate string
	ReferencePrice     int
}

func NewReferencePrice(fundID, referencePriceDate string, referencePrice int) *ReferencePrice {
	return &ReferencePrice{
		FundID:             fundID,
		ReferencePriceDate: referencePriceDate,
		ReferencePrice:     referencePrice,
	}
}
