package data

func testGetSize(tb string) (int, error) {
	bizDB := GetDataDB("default")
	var hh []struct {
		Id int32
	}
	err := bizDB.Raw("select id from ?;", tb).Find(&hh).Error
	if err != nil {
		return 0, err
	}
	return len(hh), nil
}
