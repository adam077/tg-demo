package data

import (
	"testing"
)

func TestMigrateTable(t *testing.T) {
	//t1()
	t3()
}

func t1() {
	db := GetDataDB("default")

	db.AutoMigrate(&WeiBoHotSpotMinuteReport{})
	db.Exec("CREATE UNIQUE INDEX weibo_hot_spot_minute_report_unique_idx ON weibo_hot_spot_minute_report " +
		"USING btree (log_date, log_hour, log_min, content);")
	db.AutoMigrate(&BaiDuHotSpotMinuteReport{})
	db.Exec("CREATE UNIQUE INDEX baidu_hot_spot_minute_report_unique_idx ON baidu_hot_spot_minute_report " +
		"USING btree (log_date, log_hour, log_min, content);")
	db.AutoMigrate(&ZhiHuHotSpotMinuteReport{})
	db.Exec("CREATE UNIQUE INDEX zhihu_hot_spot_minute_report_unique_idx ON zhihu_hot_spot_minute_report " +
		"USING btree (log_date, log_hour, log_min, content);")
}

func t2() {
	db := GetDataDB("default")
	db.AutoMigrate(&EatWhatTable{})
	db.AutoMigrate(&Ding{})
}

func t3() {
	db := GetDataDB("config")
	db.AutoMigrate(&ConfigTable{})
}
