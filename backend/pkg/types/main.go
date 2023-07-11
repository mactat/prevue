package types

type MetricsData struct {
	Accuracy float64 `json:"accuracy"`
	MSE      float64 `json:"mse"`
}

type ConnectorData struct {
	Uid           string      `json:"uid"`
	ProjectName   string      `json:"projectName"`
	ConnectorName string      `json:"connectorName"`
	Metrics       MetricsData `json:"metricsData"`
}
