package tools

import (
	"encoding/json"
	"errors"
	. "go_cms_reptile/models"
	"io"
	"io/ioutil"
	"net/http"
)

// SearchName  搜索 searchName http://api.kunyu77.com/api.php/provide/searchVideo?searchName=斗罗
func SearchName(searchName string) (*KanYuModel, error) {
	//爬取地址
	searchUrl := "http://api.kunyu77.com/api.php/provide/searchVideo"
	if searchName == "" {
		return nil, errors.New("searchName is nil")
	}

	req, _ := http.NewRequest(http.MethodGet, searchUrl, nil)

	//设置header
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; VTR-AL00 Build/V417IR)")
	//设置查询参数
	query := req.URL.Query()
	query.Add("searchName", searchName)

	req.URL.RawQuery = query.Encode()
	//发起 请求 响应
	resp, _ := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	//json 转 结构体
	var searchKanYu *KanYuModel
	err := json.Unmarshal(body, &searchKanYu)
	return searchKanYu, err
}

// VideoDetail 影片详情 获取指定id 获取影片信息  http://api.kunyu77.com/api.php/provide/videoDetail?ids=121904
func VideoDetail(id string) (*KanYuDetailModel, error) {
	if id == "" {
		return nil, errors.New("id is  nil")
	}

	detailUrl := "http://api.kunyu77.com/api.php/provide/videoDetail?ids=" + id

	req, err := http.NewRequest(http.MethodGet, detailUrl, nil)
	//设置header
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; VTR-AL00 Build/V417IR)")
	if err != nil {
		return nil, err
	}
	//设置header
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; VTR-AL00 Build/V417IR)")
	//发起 请求 响应
	resp, _ := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	//json 转 结构体
	var detail *KanYuDetailModel
	err = json.Unmarshal(body, &detail)
	return detail, err
}

// VideoPlaylist 获取指定影片播放链接 http://api.kunyu77.com/api.php/provide/videoPlaylist?ids=121904
func VideoPlaylist(ids string) (*VideoPlayModel, error) {

	if ids == "" {
		return nil, errors.New("ids is nil")
	}
	playUrl := "http://api.kunyu77.com/api.php/provide/videoPlaylist?ids=" + ids
	req, err := http.NewRequest(http.MethodGet, playUrl, nil)
	if err != nil {
		return nil, errors.New("ids is nil")
	}
	//设置header
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; VTR-AL00 Build/V417IR)")
	resp, err := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var videoPlay *VideoPlayModel
	err = json.Unmarshal(body, &videoPlay)
	return videoPlay, err

}

//
//func main() {
//
//	searchName := "雪中悍刀行"
//	searchKanYu, err := SearchName(searchName)
//	if err != nil {
//		fmt.Println("SearchName failed, err:", err)
//		return
//	}
//	data, err := json.Marshal(searchKanYu)
//	if err != nil {
//		fmt.Println("json.marshal failed, err:", err)
//		return
//	}
//	fmt.Printf("%s\n", string(data))
//
//	id := "129905"
//	detail, err := VideoDetail(id)
//	if err != nil {
//		fmt.Println("VideoDetail failed, err:", err)
//		return
//	}
//	fmt.Printf("%s\n\n", detail)
//
//	videoPlaylist, err := VideoPlaylist(id)
//	if err != nil {
//		fmt.Println("videoPlaylist failed, err:", err)
//		return
//	}
//	fmt.Printf("%s\n\n", videoPlaylist)
//
//}
