package jmaseis

// ヘッダ部
type Head struct {
	// "震度速報" / "震源に関する情報" / "震源・震度情報" / "遠地地震に関する情報" / forecast-type *("・" forecast-type)
	// 	forecast-type = "大津波警報" / "津波警報" / "津波注意報" / "津波予報"
	Title string
	// 発表時刻 (分値まで有効)
	ReportDateTime DateTime
	/*
		基点時刻
			- 震度速報: 地震波の検知時刻
			- 地震情報（顕著な地震の震源要素更新のお知らせ）: 震源要素を切り替えた時刻
			- 津波の観測値を発表する津波情報、沖合の津波観測に関する情報: 津波の観測状況を確定した時刻
			- その他地震・津波: 発表時刻と同値
	*/
	TargetDateTime DateTime
	// 失効時刻 (津波予報(若干の海面変動)の発表またはそれのみが残る場合) (任意項目)
	ValidDateTime DateTime
	/*
		地震識別番号 (14 桁の数字)
			- 東京・大阪システム切り替えの際は別 ID となる
			- 複数地震発生で、不整合が生じているように見える場合がある
			- 津波に関連する情報で複数の地震が関係する場合、半角スペース区切りで複数列挙 (xs:list)
	*/
	EventID string
	// "発表" / "訂正" / "取消" (訂正は 情報名称・運用種別・地震識別番号 が一致する情報単位で最新の情報を訂正する)
	InfoType        string
	Serial          string   // 情報番号
	InfoKind        string   // スキーマの運用種別情報
	InfoKindVersion string   // スキーマの運用種別情報のバージョン番号
	Headline        Headline // 見出し要素
}

// 見出し要素
type Headline struct {
	Text string // 見出し文（自由形式）
}
