package rule

type Rule struct {
	Topic    string      `json:"topic"`
	Operator string      `json:"operator"`
	Field    string      `json:"field"`
	Value    interface{} `json:"value"`
}
