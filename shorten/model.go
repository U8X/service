package shorten

// Record 用于保存的短链接记录
type Record struct {
	LongUrl  string `json:"long_url,omitempty" bson:"long_url,omitempty"`
	ShortUrl string `json:"short_url,omitempty" bson:"short_url,omitempty"`
}
