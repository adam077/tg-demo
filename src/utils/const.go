package utils

var ConstData = new(StructConstData)

type StructConstData struct {
	ZSReport struct {
		Status struct {
			CheckFail int32 `default:"-1"` // 检查失败
			NoReport  int32 `default:"0"`  // 钻展上无报表
			GotReport int32 `default:"1"`  // 钻展上有报表
		}
	}
	TZReport struct {
		Status struct {
			CheckFail int32 `default:"-1"` // 检查失败
			NoReport  int32 `default:"0"`  // 天钻上无报表
			GotReport int32 `default:"1"`  // 天钻上有报表
		}
	}
	Redis struct {
		KeyFix struct {
			ReportTrack string `default:"Aibi2MonitorReport"`
		}
	}
}

func init() {
	FillWithDefault(ConstData)
}
