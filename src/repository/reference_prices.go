package repository

func ExistReferencePriceByDate(date string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM ReferencePrices WHERE ReferencePriceDate = ?", date).Scan(&count)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
