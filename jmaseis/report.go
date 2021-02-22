package jmaseis

// 気象庁防災情報XML
type Report struct {
	Control Control // 管理部
	Head    Head    // ヘッダ部
	Body    Body    // 内容部
}
