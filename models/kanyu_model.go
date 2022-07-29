package models

// KanYu 搜索 searchName http://api.kunyu77.com/api.php/provide/searchVideo?searchName=斗罗
type KanYuModel struct {
	Code         int            `json:"code"`
	Msg          string         `json:"msg"`
	Data         []Data         `json:"data"`
	Ext          Ext            `json:"ext"`
	Pages        int            `json:"pages"`
	ClassifyList []ClassifyList `json:"classifyList"`
}
type Data struct {
	ID              int      `json:"id"`
	VideoName       string   `json:"videoName"`
	StarName        string   `json:"starName"`
	VideoCover      string   `json:"videoCover"`
	Imgh            string   `json:"imgh"`
	TypeID          int      `json:"type_id"`
	Subtitle        string   `json:"subtitle"`
	Msg             string   `json:"msg"`
	Year            string   `json:"year"`
	Tag             string   `json:"tag"`
	Finish          string   `json:"finish"`
	PlayNum         string   `json:"playNum"`
	Score           string   `json:"score"`
	UpdateTime      string   `json:"update_time"`
	Brief           string   `json:"brief"`
	VodClass        string   `json:"vod_class"`
	BriefContext    string   `json:"briefContext"`
	SubCategory     string   `json:"subCategory"`
	VodHitsWeek     int      `json:"vod_hits_week"`
	VideoCommentNum int      `json:"videoCommentNum"`
	VideoUrls       []string `json:"videoUrls"`
	IsLike          string   `json:"isLike"`
	Cent            int      `json:"cent"`
	StarID          int      `json:"starId"`
	IsCare          string   `json:"isCare"`
	Tags            string   `json:"tags"`
}
type Page struct {
	CurPage int `json:"cur_page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}
type ClassifyList struct {
	ID         int    `json:"id"`
	VideoName  string `json:"videoName"`
	VideoCover string `json:"videoCover"`
	TypeID     int    `json:"type_id"`
	Score      string `json:"score"`
	Msg        string `json:"msg"`
	Year       string `json:"year"`
	VideoURL   string `json:"videoUrl"`
}
type Ext struct {
	Page         Page           `json:"page"`
	ClassifyList []ClassifyList `json:"classifyList"`
}

// KanYuDetailModel 详细 获取指定id 获取影片信息  http://api.kunyu77.com/api.php/provide/videoDetail?ids=121904
type KanYuDetailModel struct {
	Code            int             `json:"code"`
	Msg             string          `json:"msg"`
	UseView         string          `json:"useView"`
	KanYuDetailData KanYuDetailData `json:"data"`
}
type KanYuDetailLikeVideoList struct {
	VideoCover string      `json:"videoCover"`
	ID         int         `json:"id"`
	TagID      int         `json:"tagId"`
	VideoName  string      `json:"videoName"`
	Score      string      `json:"score"`
	Msg        string      `json:"msg"`
	Year       string      `json:"year"`
	VideoURL   string      `json:"videoUrl"`
	TagsName   interface{} `json:"tagsName"`
	Brief      string      `json:"brief"`
}
type KanYuDetailData struct {
	VideoCover      string                     `json:"videoCover"`
	Imgh            string                     `json:"imgh"`
	ID              int                        `json:"id"`
	TypeID          int                        `json:"type_id"`
	VideoName       string                     `json:"videoName"`
	Score           string                     `json:"score"`
	Msg             string                     `json:"msg"`
	Year            string                     `json:"year"`
	Area            string                     `json:"area"`
	Director        string                     `json:"director"`
	Actor           string                     `json:"actor"`
	SubCategory     string                     `json:"subCategory"`
	Episode         int                        `json:"episode"`
	Finish          int                        `json:"finish"`
	Brief           string                     `json:"brief"`
	ClassifyName    string                     `json:"classifyName"`
	PlayNum         int                        `json:"playNum"`
	StarName        string                     `json:"starName"`
	IsLike          string                     `json:"isLike"`
	VideoCommentNum int                        `json:"videoCommentNum"`
	Cent            int                        `json:"cent"`
	StarID          int                        `json:"starId"`
	IsCare          string                     `json:"isCare"`
	LikeVideoList   []KanYuDetailLikeVideoList `json:"likeVideoList"`
	CategoryEn      string                     `json:"categoryEn"`
	Download        int                        `json:"download"`
}

// VideoPlayModel 3、获取指定影片播放链接 http://api.kunyu77.com/api.php/provide/videoPlaylist?ids=121904
type VideoPlayModel struct {
	Code    int           `json:"code"`
	Msg     string        `json:"msg"`
	UseView string        `json:"useView"`
	Data    VideoPlayData `json:"data"`
}
type VideoPlayPlayurls struct {
	Title    string `json:"title"`
	Playfrom string `json:"playfrom"`
	Playurl  string `json:"playurl"`
}
type VideoPlayEpisodes struct {
	ID           string              `json:"id"`
	CollectionID int                 `json:"collection_id"`
	Imgv         string              `json:"imgv"`
	Episode      int                 `json:"episode"`
	Title        string              `json:"title"`
	Playurl      string              `json:"playurl"`
	Playurls     []VideoPlayPlayurls `json:"playurls"`
	Source       string              `json:"source"`
	Download     int                 `json:"download"`
	Needjump     int                 `json:"needjump"`
	JumpURL      string              `json:"jump_url"`
	AlbumTitle   string              `json:"albumTitle"`
}
type VideoPlayData struct {
	Episodes []VideoPlayEpisodes `json:"episodes"`
}
