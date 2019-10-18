package beego

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type Res struct {
	Meta Meta            `json:"meta"`
	Data json.RawMessage `json:"data"`
}

type Meta struct {
	Status      int `json:"status"`
	Code        int `json:"code"`
	TotalCount  int `json:"total_count"`
	PageCount   int `json:"page_count"`
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
}

type Controller struct {
	beego.Controller
}

/*
* 成功跳转
 */
func (c *Controller) ResJson(data interface{}) {
	c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()
}

// 获取 post JSON 数据 转换成入参类型
func (c *Controller) GetPostJson(v interface{}) ([]byte, error) {
	b := c.Ctx.Input.RequestBody
	err := json.Unmarshal(b, &v)
	return b, err
}
