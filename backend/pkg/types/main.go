package types

type MetricsData struct {
	Epoch        int64  `json:"epoch"`
	Batch        int64  `json:"batch"`
	LossName     string `json:"loss_name"`
	LossValue    string `json:"loss_value"`
	MatricsName  string `json:"metrics_name"`
	MatricsValue string `json:"metrics_value"`
}

type UserData struct {
	UserName  string `json:"user_id"`
	Email     string `json:"email"`
	Passwoard string `json:"passwoard"`
}

type ProjectData struct {
	ProjectName string `json:"project_name"`
}

type ModelsData struct {
	ModelName     string `json:"model_name"`
	ConnectorName string `json:"connector"`
	Architecture  string `json:"architecture"`
	Weights       string `json:"weights"`
}

type SessionData struct {
	User    UserData    `json:"userData"`
	Models  ModelsData  `json:"modelsData"`
	Project ProjectData `json:"projectData"`
}

type SessionMetrics struct {
	Metrics MetricsData `json:"metricsData"`
}
