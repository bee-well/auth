package config

func MqConnectionUrl() string {
	if IsProduction() {
		return ""
	}
	return "amqp://guest:guest@localhost:5672"
}

func SqlConnectionUrl() string {
	if IsProduction() {
		return "postgres://qadxdkigrkvaui:49a2719a2121ada6b6e99d9a78b537660bbfc04f523d07979820381d2e07bbca@ec2-54-90-211-192.compute-1.amazonaws.com:5432/d1ee3rm12bro17"
	}
	return "postgres://qadxdkigrkvaui:49a2719a2121ada6b6e99d9a78b537660bbfc04f523d07979820381d2e07bbca@ec2-54-90-211-192.compute-1.amazonaws.com:5432/d1ee3rm12bro17"
}

func IsProduction() bool {
	return false
}
