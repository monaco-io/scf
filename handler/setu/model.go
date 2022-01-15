package setu

var (
	SeTuURL = "https://api.lolicon.app/setu/v2"
	AllSize = []string{"original", "regular", "small", "thumb", "mini"}
)

type Request struct {
	// 0为非 R18，1为 R18，2为混合（在库中的分类，不等同于作品本身的 R18 标识）
	R18 int64 `json:"r18"`
	// 一次返回的结果数量，范围为1到100；在指定关键字或标签的情况下，结果数量可能会不足指定的数量
	Num int64 `json:"num"`
	// 返回指定uid作者的作品，最多20个
	UID int64 `json:"uid"`
	// 返回从标题、作者、标签中按指定关键字模糊匹配的结果，大小写不敏感，性能和准度较差且功能单一，建议使用tag代替
	Keyword string `json:"keyword"`
	// 返回匹配指定标签的作品
	Tag []string `json:"tag"`
	// 返回指定图片规格的地址 original/regular/small/thumb/mini
	Size []string `json:"size"`
	// 设置图片地址所使用的在线反代服务 i.pixiv.cat
	Proxy string `json:"proxy"`
	// 返回在这个时间及以后上传的作品；时间戳，单位为毫秒
	DataAfter int64 `json:"dataAfter"`
	// 返回在这个时间及以前上传的作品；时间戳，单位为毫秒
	DataBefore int64 `json:"dataBefore"`
	// 设置为任意真值以禁用对某些缩写keyword和tag的自动转换
	Dsc bool `json:"dsc"`
}

type Response struct {
	// 错误信息
	Error string `json:"error"`
	// 色图数组
	Data []Data `json:"data"`
}

type Data struct {
	// 作品 pid
	PID int64 `json:"pid"`
	// 作品所在页
	P int64 `json:"p"`
	// 作者 uid
	UID int64 `json:"uid"`
	// 作品标题
	Title string `json:"title"`
	// 作者名（入库时，并过滤掉 @ 及其后内容）
	Author string `json:"author"`
	// 是否 R18（在库中的分类，不等同于作品本身的 R18 标识）
	R18 bool `json:"r18"`
	// 原图宽度 px
	Width int64 `json:"width"`
	// 原图高度 px
	Height int64 `json:"height"`
	// 作品标签，包含标签的中文翻译（有的话）
	Tags []string `json:"tags"`
	// 图片扩展名
	Ext string `json:"ext"`
	// 作品上传日期；时间戳，单位为毫秒
	UploadDate int64 `json:"uploadDate"`
	// 包含了所有指定size的图片地址
	URLs URLs `json:"urls"`
}

type URLs struct {
	Original string `json:"original"`
	Regular  string `json:"regular"`
	Small    string `json:"small"`
	Thumb    string `json:"thumb"`
	Mini     string `json:"mini"`
}

type Picture struct {
	Title string
	Tags  []string
	URL   string
	Ext   string
}
