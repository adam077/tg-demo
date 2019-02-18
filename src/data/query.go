package data

func GetSize() (int, error) {
	bizDB := GetDataDB("default")
	var hh []struct {
		Id int32
	}
	err := bizDB.Raw("select id from zs_campaign_day_report limit ?;", 15).Find(&hh).Error
	if err != nil {
		return 0, err
	}
	return len(hh), nil
}
