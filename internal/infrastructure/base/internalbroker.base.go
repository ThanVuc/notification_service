package base

// NotificationJob định nghĩa cấu trúc job gửi vào channel
type NotificationJob struct {
	NotificationID string
	RetryCount     int
	DirectEmails   []string
}

type NotificationJobOption func(*NotificationJob)

func NewNotificationJob(opts ...NotificationJobOption) NotificationJob {
	job := NotificationJob{
		RetryCount: 0,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&job)
		}
	}

	return job
}

func WithNotificationID(notificationID string) NotificationJobOption {
	return func(job *NotificationJob) {
		job.NotificationID = notificationID
	}
}

func WithRetryCount(retryCount int) NotificationJobOption {
	return func(job *NotificationJob) {
		job.RetryCount = retryCount
	}
}

func WithDirectEmails(directEmails []string) NotificationJobOption {
	return func(job *NotificationJob) {
		job.DirectEmails = append([]string(nil), directEmails...)
	}
}

// Dispatcher quản lý các channels
type Dispatcher struct {
	AppChan         chan NotificationJob
	EmailChan       chan NotificationJob
	DirectEmailChan chan NotificationJob
}

// NewDispatcher khởi tạo dispatcher mới
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		AppChan:         make(chan NotificationJob, 1000),
		EmailChan:       make(chan NotificationJob, 1000),
		DirectEmailChan: make(chan NotificationJob, 1000),
	}
}

func (d *Dispatcher) Close() {
	close(d.AppChan)
	close(d.EmailChan)
	close(d.DirectEmailChan)
}
