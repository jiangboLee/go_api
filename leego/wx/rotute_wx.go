package wx

import (
	"net/http"
	"leego/utils"
	"fmt"
	"encoding/json"
)

const (
	appId = "wx1b6316b214ec7e55"
	appSecret = "74613cb2678dcfbeb19d87a5dae7af53"
)
type Access_token struct {
	Access_token string `json:"access_token"`
}

type Template struct {
	Touser string `json:"touser"`
	Template_id string `json:"template_id"`
	Page string `json:"page"`
	Form_id string `json:"form_id"`
	Data Keyword `json:"data"`
	Emphasis_keyword string `json:"emphasis_keyword"`
}

type Keyword struct {
	Keyword1 Content `json:"keyword1"`
	Keyword2 Content `json:"keyword2"`
	Keyword3 Content `json:"keyword3"`
	Keyword4 Content `json:"keyword4"`
	Keyword5 Content `json:"keyword5"`
}
type Content struct {
	Value string `json:"value"`
	Color string `json:"color"`
}
// killall -9 main 终止程序
func GetAppid(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	code :=  r.FormValue("code")
	if len(code) > 0 {
		fmt.Fprintf(w, wxgetAppid(code))
	} else {
		fmt.Fprintf(w, "code为空")
	}
}

func SendTemple() {
	str := utils.HttpGet(utils.GetAccess_token + "appid=" + appId + "&secret=" + appSecret)
	accessToken := Access_token{}
	err := json.Unmarshal([]byte(str), &accessToken)
	if err != nil {
		utils.Warning(err, "json解析失败")
	}
	utils.P(accessToken)
	template := Template{
		Touser: "of5WV5KKLMFN8IYPFW00QUbfGLk8",
		Template_id: "dKqVzsWdlOgnQkwCX0EDjHNEcbg5JYZNlFI-AbNK_Fo",
		Page: "pages/index/index",
		Form_id: "5d4c89163ecb87047da900f47fc020f3",
		Data: Keyword{
			Keyword1: Content{
				Value: "加油啊",
				Color: "#000000",
			},
			Keyword2: Content{
				Value:"1991-7-9",
				Color: "#000000",
			},
			Keyword3: Content{
				Value:"1111",
				Color: "#000000",
			},
			Keyword4: Content{
				Value:"1111",
				Color: "#000000",
			},
			Keyword5: Content{
				Value:"1111",
				Color: "#000000",
			},
		},
		Emphasis_keyword: "keyword1.DATA",
	}
	output, err := json.MarshalIndent(&template, "", "\t\t")
	if err != nil {
		utils.Danger(err, "对象转json错误")
	}
	utils.P(output)
	str = utils.HttpPost(utils.SendTemplate+accessToken.Access_token, output)
	utils.P(str)
}

func wxgetAppid(code string) string {
	return utils.HttpGet("https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecret + "&js_code=" + code+ "&grant_type=authorization_code")
}

