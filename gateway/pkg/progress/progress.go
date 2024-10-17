package progress

import "time"

var Map map[string][]*Step

type Step struct {
	Name      ProgressStep  `json:"name"`
	State     ProgressState `json:"state"`
	StartTime time.Time     `json:"start_time"`
	Logs      []string      `json:"logs"`
}

type ProgressState int
type ProgressStep int

const (
	PushDockerfile ProgressStep = iota
	DeployKaniko
	PushManifest
	DeployResource
)

const (
	Success ProgressState = iota
	Fail
	Continue
)

// func New() *Progress {
// 	return &Progress{
// 		Map: make(map[string][]*Step),
// 	}
// }

// func (p *Progress) IsExist(uuid string) bool {
// 	if _, ok := p.Map[uuid]; ok {
// 		return true
// 	}
// 	return false
// }

// func (p *Progress) LastStep(uuid string) *Step {
// 	l := len(p.Map[uuid])
// 	return p.Map[uuid][l-1]
// }

func (s *Step) Add(imgTag string) {
	Map[imgTag] = append(Map[imgTag], s)
}

func (s *Step) UpdateLog(log string) {
	s.Logs = append(s.Logs, log)
}

func (s *Step) UpdateState(state ProgressState, log string) {
	s.State = state
	s.Logs = append(s.Logs, log)
}

func NewStep(name ProgressStep, state ProgressState, logs []string) *Step {
	return &Step{
		Name:      name,
		State:     state,
		StartTime: time.Now(),
		Logs:      []string{},
	}
}
