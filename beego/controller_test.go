package beego

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"testing"
)

// Test_GetPostJson
func Test_GetPostJson(t *testing.T) {
	i := context.NewInput()
	i.RequestBody = []byte(`{"age":40}`)

	ctx := &context.Context{Input: i}
	contr := Controller{beego.Controller{Ctx: ctx}}

	user := struct {
		Age int64 `json:"age"`
	}{}
	_, e := contr.GetPostJson(&user)

	if e != nil {
		t.Error(e)
	}

	val := user.Age
	if val != 40 {
		t.Errorf("TestGeetInt64 expect 40,get %T,%v", val, val)
	}
}

// Test_ResJson TODO Test_ResJson
func Test_ResJson(t *testing.T) {

}
