package common

type ApplicationContext struct {
	InstanceId string
	AppBaseUrl string
}

func NewApplicationContext(instanceId string, appBaseUrl string) *ApplicationContext {
	return &ApplicationContext{
		InstanceId: instanceId,
		AppBaseUrl: appBaseUrl,
	}
}
