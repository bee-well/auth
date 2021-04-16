package config

func MqConnectionUrl() string {
	if IsProduction() {
		return ""
	}
	return "amqp://guest:guest@localhost:5672"
}

func IsProduction() bool {
	return false
}
