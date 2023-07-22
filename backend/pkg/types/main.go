package types

type MetricsData struct {
	ModelId      int64   `json:"model_id"`
	Epoch        int64   `json:"epoch"`
	Batch        int64   `json:"batch"`
	LossName     string  `json:"loss_name"`
	LossValue    float64 `json:"loss_value"`
	MatricsName  string  `json:"metrics_name"`
	MatricsValue float64 `json:"metrics_value"`
}

type UserData struct {
	UserName  string `json:"user_id"`
	Email     string `json:"email"`
	Password string `json:"password"`
}

type ProjectData struct {
	ProjectName string `json:"project_name"`
}

type ModelsData struct {
	ModelName     string `json:"model_name"`
	ConnectorName string `json:"connector"`
	Architecture  string `json:"architecture"`
}

type SessionData struct {
	User    UserData    `json:"userData"`
	Models  ModelsData  `json:"modelsData"`
	Project ProjectData `json:"projectData"`
}

type SessionMetrics struct {
	Metrics MetricsData `json:"metricsData"`
}
