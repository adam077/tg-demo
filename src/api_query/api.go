package api_query

var MultiApi = struct {
	WeatherApi string
}{}

func init() {
	MultiApi.WeatherApi = "http://t.weather.sojson.com/api/weather/city"
}
