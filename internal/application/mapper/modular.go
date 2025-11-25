package mapper

type MapperModule struct {
	NotificationMapper NotificationMapper
}

func NewMapperModule() *MapperModule {
	notificationMapper := NewNotificationMapper()
	return &MapperModule{
		NotificationMapper: notificationMapper,
	}
}
