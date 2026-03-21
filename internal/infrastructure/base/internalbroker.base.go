package base

// NotificationJob định nghĩa cấu trúc job gửi vào channel
type NotificationJob struct {
	NotificationID string
	RetryCount     int
}

// Dispatcher quản lý các channels
type Dispatcher struct {
	AppChan   chan NotificationJob
	EmailChan chan NotificationJob
}

// NewDispatcher khởi tạo dispatcher mới
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		AppChan:   make(chan NotificationJob, 1000),
		EmailChan: make(chan NotificationJob, 1000),
	}
}

func (d *Dispatcher) Close() {
	close(d.AppChan)
	close(d.EmailChan)
}
