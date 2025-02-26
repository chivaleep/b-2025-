package main

//this code is to acquire video info
//this code takes the last aid as the last page then go to next page
//for more crawling
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Rights struct {
	Bp            int `json:"bp"`
	Elec          int `json:"elec"`
	Download      int `json:"download"`
	Movie         int `json:"movie"`
	Pay           int `json:"pay"`
	Hd5           int `json:"hd5"`
	NoReprint     int `json:"no_reprint"`
	Autoplay      int `json:"autoplay"`
	UgcPay        int `json:"ugc_pay"`
	IsCooperation int `json:"is_cooperation"`
	UgcPayPreview int `json:"ugc_pay_preview"`
	NoBackground  int `json:"no_background"`
	ArcPay        int `json:"arc_pay"`
	PayFreeWatch  int `json:"pay_free_watch"`
}

type Owner struct {
	Mid  int    `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}

type Stat struct {
	Aid      int `json:"aid"`
	View     int `json:"view"`
	Danmaku  int `json:"danmaku"`
	Reply    int `json:"reply"`
	Favorite int `json:"favorite"`
	Coin     int `json:"coin"`
	Share    int `json:"share"`
	Like     int `json:"like"`
	Dislike  int `json:"dislike"`
	Vt       int `json:"vt"`
	Vv       int `json:"vv"`
}

type VideoItem struct {
	Aid            int    `json:"aid"`
	Videos         int    `json:"videos"`
	Tid            int    `json:"tid"`
	Tname          string `json:"tname"`
	Copyright      int    `json:"copyright"`
	Pic            string `json:"pic"`
	Title          string `json:"title"`
	PubDate        int    `json:"pubdate"`
	Ctime          int    `json:"ctime"`
	Desc           string `json:"desc"`
	State          int    `json:"state"`
	Duration       int    `json:"duration"`
	Owner          Owner  `json:"owner"`
	Stat           Stat   `json:"stat"`
	Bvid           string `json:"bvid"`
	ShortLink      string `json:"short_link"`
	FirstFrame     string `json:"first_frame"`
	PubLocation    string `json:"pub_location"`
	Cover43        string `json:"cover_43"`
	Tidv2          int    `json:"tid_v2"`
	Tnamev2        string `json:"tname_v2"`
	SeasonType     int    `json:"season_type"`
	IsOgv          bool   `json:"is_ogv"`
	EnableVt       int    `json:"enable_vt"`
	AiRcmd         int    `json:"ai_rcmd"`
	RcmdContent    string `json:"rcmd_content"`
	RcmdCornerMark string `json:"rcmd_corner_mark"`
}

type Response struct {
	Data struct {
		List []VideoItem `json:"list"`
	} `json:"data"`
}

func fetchBilibiliData(lastAid int) ([]VideoItem, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.bilibili.com/x/web-interface/popular?ps=20&web_location=333.934")
	if lastAid != 0 {
		url += fmt.Sprintf("&last_aid=%d", lastAid)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", randomUserAgent())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		log.Println("请求过于频繁，暂停 60 秒...")
		time.Sleep(60 * time.Second)
		return nil, fmt.Errorf("请求过于频繁")
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.List, nil
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/popular?ps=20&pn=1&web_location=333.934&w_rid=feb8a1b47a6978a2867e5beb1b8997a3&wts=1740310618", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh,zh-TW;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("cookie", "buvid3=200818E0-92AF-BEFA-BEEB-F15FEE1306E864195infoc; b_nut=1735567264; _uuid=1BE91876-C6E9-235D-4618-1F8A7B542A8260340infoc; buvid_fp=9a99db0341ad48243613d64ea34359eb; buvid4=85688011-7D42-6856-087D-5CFAABBA2CB064796-024123014-Q%2BrYyl7l7bw%2BbssEgBD97Q%3D%3D; enable_web_push=DISABLE; rpdid=|(YuJJl)|JR0J'u~JllY~k)R; DedeUserID=289070327; DedeUserID__ckMd5=52ceabd7a9dcfce6; header_theme_version=CLOSE; enable_feed_channel=DISABLE; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAzNzk3MjcsImlhdCI6MTc0MDEyMDQ2NywicGx0IjotMX0.HFHiU2H2BiX0IRnAMcKBcOgjx1wT0cE0PdTCWhFVyAw; bili_ticket_expires=1740379667; SESSDATA=3b26acd0%2C1755672528%2Cbe4eb%2A21CjDcIpeUenl_YIsWKXbncULl3PVXilxGWW9bSzzXoh_Lnxm_jn6ABgr319pAkOS9iYwSVkotTVhvekdnamhnUWUycUxGWGZmY1J5M0JwQ2NZUjlVd1gySlpQNkdFZ2RGZ2dtcGhwNjZKRkdnTUZKRnN6N1VJaXR4SmZOR09MVERCTlNWMWwwS1RBIIEC; bili_jct=a45c3cc89c4664dfa278deac037fb12b; sid=8r962k9i; bsource=search_google; b_lsid=98FDDAD2_19532832502; home_feed_column=5; browser_resolution=1467-689; bp_t_offset_289070327=1037109687155163136; CURRENT_FNVAL=2000")
	req.Header.Set("origin", "https://www.bilibili.com")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.bilibili.com/v/popular/all/?spm_id_from=333.1007.0.0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the body of the response
	file, err := os.Create("bilibili_popular.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Aid", "Videos", "Tid", "Tname", "Copyright", "Pic", "Title", "PubDate", "Ctime", "Desc",
		"State", "Duration", "OwnerMid", "OwnerName", "OwnerFace", "View", "Danmaku", "Reply", "Favorite",
		"Coin", "Share", "Like", "Dislike", "Vv", "ShortLink", "FirstFrame", "PubLocation", "Cover43",
		"Tidv2", "Tnamev2", "Bvid", "SeasonType", "IsOgv", "EnableVt", "AiRcmd", "RcmdContent", "RcmdCornerMark"})

	var lastAid int
	maxRequests := 10

	for i := 0; i < maxRequests; i++ {
		videos, err := fetchBilibiliData(lastAid)
		if err != nil {
			log.Println("爬取失败:", err)
			continue
		}

		if len(videos) == 0 {
			log.Println("没有更多数据，停止爬取。")
			break
		}

		for _, item := range videos {
			writer.Write([]string{
				fmt.Sprintf("%d", item.Aid),
				fmt.Sprintf("%d", item.Videos),
				fmt.Sprintf("%d", item.Tid),
				item.Tname,
				fmt.Sprintf("%d", item.Copyright),
				item.Pic,
				item.Title,
				fmt.Sprintf("%d", item.PubDate),
				fmt.Sprintf("%d", item.Ctime),
				item.Desc,
				fmt.Sprintf("%d", item.State),
				fmt.Sprintf("%d", item.Duration),
				fmt.Sprintf("%d", item.Owner.Mid),
				item.Owner.Name,
				item.Owner.Face,
				fmt.Sprintf("%d", item.Stat.View),
				fmt.Sprintf("%d", item.Stat.Danmaku),
				fmt.Sprintf("%d", item.Stat.Reply),
				fmt.Sprintf("%d", item.Stat.Favorite),
				fmt.Sprintf("%d", item.Stat.Coin),
				fmt.Sprintf("%d", item.Stat.Share),
				fmt.Sprintf("%d", item.Stat.Like),
				fmt.Sprintf("%d", item.Stat.Dislike),
				fmt.Sprintf("%d", item.Stat.Vv),
				item.ShortLink,
				item.FirstFrame,
				item.PubLocation,
				item.Cover43,
				fmt.Sprintf("%d", item.Tidv2),
				item.Tnamev2,
				item.Bvid,
				fmt.Sprintf("%d", item.SeasonType),
				fmt.Sprintf("%t", item.IsOgv),
				fmt.Sprintf("%t", item.EnableVt),
				fmt.Sprintf("%d", item.AiRcmd),
				item.RcmdContent,
				item.RcmdCornerMark,
			})
			lastAid = item.Aid
		}

		writer.Flush()

		delay := time.Duration(rand.Intn(5)+3) * time.Second
		log.Printf("第 %d 轮爬取成功，休息 %v 秒\n", i+1, delay.Seconds())
		time.Sleep(delay)
	}

	fmt.Println("数据爬取完成，CSV 已生成！")
}

func randomUserAgent() string {
	agents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Mobile/15E148 Safari/604.1",
	}
	return agents[rand.Intn(len(agents))]
}
