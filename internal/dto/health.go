package dto

type HealthCheck struct {
	Status string  `json:"status"`
}

func (h *HealthCheck) Up(status bool) {
	if status {
		h.Status = "UP"
	} else {
		h.Status = "DOWN"
	}
}

func (h HealthCheck) Healthy() bool {
	return h.Status == "UP"
}