package constant

type RiskType int

const (
	Low RiskType = iota
	Medium
	High
)

func (riskType RiskType) String() string {
	switch riskType {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	}
	return ""
}
