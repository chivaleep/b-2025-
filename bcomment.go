package main

//this code is to acquire comments of one video
//need to have a page flipping function
//need to add a text to go through different videos

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type BilibiliResponse struct {
	Data struct {
		Replies []struct {
			Member struct {
				Mid       string `json:"mid"`   // 用户ID
				Uname     string `json:"uname"` // 用户名
				Sex       string `json:"sex"`   // 用户性别
				LevelInfo struct {
					CurrentLevel int `json:"current_level"` // 用户等级
				} `json:"level_info"`
				VIP struct {
					VipStatus int `json:"vipStatus"` // 0 = 非VIP, 1 = VIP
				} `json:"vip"`
			} `json:"member"`
			ReplyControl struct {
				Location string `json:"location"` // IP属地
			} `json:"reply_control"`
			Content struct {
				Message string `json:"message"` // 评论内容
			} `json:"content"`
			Like int `json:"like"` // 点赞数
		} `json:"replies"`
	} `json:"data"`
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/v2/reply/wbi/main?oid=610513693&type=1&mode=3&pagination_str=%7B%22offset%22:%22%22%7D&plat=1&seek_rpid=&web_location=1315875&w_rid=ffed64f4e76af5c9d84dcd0abb1effbb&wts=1740560793", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh,zh-TW;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("cookie", "buvid3=200818E0-92AF-BEFA-BEEB-F15FEE1306E864195infoc; b_nut=1735567264; _uuid=1BE91876-C6E9-235D-4618-1F8A7B542A8260340infoc; buvid_fp=9a99db0341ad48243613d64ea34359eb; buvid4=85688011-7D42-6856-087D-5CFAABBA2CB064796-024123014-Q%2BrYyl7l7bw%2BbssEgBD97Q%3D%3D; enable_web_push=DISABLE; rpdid=|(YuJJl)|JR0J'u~JllY~k)R; DedeUserID=289070327; DedeUserID__ckMd5=52ceabd7a9dcfce6; header_theme_version=CLOSE; enable_feed_channel=DISABLE; bsource=search_google; hit-dyn-v2=1; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA1Nzg3MjIsImlhdCI6MTc0MDMxOTQ2MiwicGx0IjotMX0.CQio2oz9qfj_6mWQnferihCQ4MK4PuvKE-rb2KT060c; bili_ticket_expires=1740578662; msource=pc_web; deviceFingerprint=d983ba35a266005a1fe4316297a3dde5; SESSDATA=42e58015%2C1755942020%2Ca328d%2A21CjBLEZHuRk5U4VmUgsxcSthQGPl6O27jhH1h-R71-SaD-A-mDmqVsE8WPnHJ-lfBW9ISVjBhaXlVdktONks4UFlPZHFFSXlmUC1WSjM2RGdZS1V1eUdPbVA1ekIxTEE5c3lPNUIxejVuVlpEb2d1azk3UnJQay1GUEo0Z0JDTkFOWDhHNlg3ZnNnIIEC; bili_jct=443749b4834ddb7eb6c52b6d63c55fd5; sid=6tmo6vl9; LIVE_BUVID=AUTO2317403901475805; PVID=2; is-2022-channel=1; home_feed_column=5; browser_resolution=1467-689; b_lsid=98F2B5E1_195413222F9; bp_t_offset_289070327=1038173340100984832; CURRENT_FNVAL=4048")
	req.Header.Set("origin", "https://www.bilibili.com")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.bilibili.com/video/BV1X84y1K7bv/?spm_id_from=333.337.search-card.all.click&vd_source=bd5f4acd5efa8f96db84cbe2bfa94918")
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
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	var data BilibiliResponse
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		log.Fatal("JSON 解析失败:", err)
	}

	// 4. 创建 CSV 文件
	file, err := os.Create("bilibili_comments.csv")
	if err != nil {
		log.Fatal("无法创建 CSV 文件:", err)
	}
	defer file.Close()

	// 5. 写入 CSV 头部
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"用户ID", "用户名", "性别", "用户等级", "IP属地", "是否VIP", "评论内容", "点赞数"})

	// 6. 遍历数据并写入 CSV
	for _, reply := range data.Data.Replies {
		vipStatus := "否"
		if reply.Member.VIP.VipStatus == 1 {
			vipStatus = "是"
		}
		record := []string{
			reply.Member.Mid,   // 用户ID
			reply.Member.Uname, // 用户名
			reply.Member.Sex,   // 性别
			fmt.Sprintf("%d", reply.Member.LevelInfo.CurrentLevel), // 用户等级
			reply.ReplyControl.Location,                            // IP 属地
			vipStatus,                                              // 是否 VIP
			reply.Content.Message,                                  // 评论内容
			fmt.Sprintf("%d", reply.Like),                          // 点赞数
		}
		writer.Write(record)
	}

	fmt.Println("数据已成功存入 bilibili_comments.csv")
}
