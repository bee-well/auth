package mq

type MqMock struct {
	PublishMock        func(string, []byte) error
	PublishCalls       int
	AttachHandlerMock  func(string, MqHandlerFunc) error
	AttachHandlerCalls int
}

func MockMq(m *MqMock) {
	mock = m
}

func (mock *MqMock) Publish(q string, b []byte) error {
	mock.PublishCalls++
	return mock.PublishMock(q, b)
}

func (mock *MqMock) AttachHandler(q string, h MqHandlerFunc) error {
	mock.AttachHandlerCalls++
	return mock.AttachHandlerMock(q, h)
}
