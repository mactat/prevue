package types

type MetricsData struct {
	Epoch 			int64 	 	`json:"epoch"`
	Batch 			int64 		`json:"batch"`
	LossName    	string  	`json:"loss_name"`
	LossValue   	string  	`json:"loss_value"`
	MatricsName  	string  	`json:"metrics_name"`
	MatricsValue  	string  	`json:"metrics_value"`
}

type UserData struct {
	UserName 	string 	`json:"user_id"`
	Email 		string	`json:"email"`
	Passwoard 	string	`json:"passwoard"`
}

type ProjectData struct {
	ProjectName string 	`json:"project_name"`
}

type ModelsData struct {
	ModelName 		string 	`json:"model_name"`
	ConnectorName 	string 	`json:"connector"`
	Architecture 	string 	`json:"architecture"`
	Weights 		string 	`json:"weights"`
}


type ConnectorData struct {
	Metrics       		MetricsData  	`json:"metricsData"`
	User       			UserData  	 	`json:"userData"`
	Project       		ProjectData  	`json:"projectData"`
	Models       		ModelsData  	`json:"modelsData"`
}
