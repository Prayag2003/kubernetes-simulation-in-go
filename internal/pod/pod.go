package pod

type Pod struct {
	ID 			string `json:"id"`
	Status  	string `json:"status"`
	StartTime 	time.Time `json:"start_time"`
	StopChan  	chan struct{}
}
