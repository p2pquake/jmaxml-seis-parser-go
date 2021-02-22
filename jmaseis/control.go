package jmaseis

// 管理部
type Control struct {
	Title            string   // 情報名称
	DateTime         DateTime // 発表時刻 (秒値まで有効)
	Status           string   // 通常/訓練/試験
	EditorialOffice  string   // 編集官署名
	PublishingOffice []string // 発表官署名 (地震火山関連情報は要素数 1)
}
