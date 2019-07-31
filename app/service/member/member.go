package s_member

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/pkg/errors"
	"fmt"
)

const (
	USER_SESSION_MARK = "user_info"
)

var (
	// 表对象
	table = g.DB().Table("member").Safe()
)

// 用户注册
func SignUp(data g.MapStrStr) error {
	// 数据校验
	rules := map[string]string {
		"email"  : "required| email#邮箱不能为空|邮箱格式不正确",
		"password"  : "required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间",
		"password2" : "required|length:6,16#密码不能为空|密码长度应当在:min到:max之间",
	}

	if e := gvalid.CheckMap(data, rules); e != nil {
		return errors.New(e.String())
	}

	// 唯一性数据检查
	if !CheckEmail(data["email"]) {
		return errors.New(fmt.Sprintf("邮箱 %s 已经存在", data["email"]))
	}

	// 记录账号创建/注册时间
	if _, ok := data["create_time"]; !ok {
		data["create_time"] = gtime.Now().String()
	}
	if _, err := table.Filter().Data(data).Save(); err != nil {
		return err
	}
	return nil
}

//检查电话号码 是否存在 ,存在则返回 false ,否则 返回 true
func CheckPhoneNumber(phoneNumber string) bool  {
	if i,err := table.Where("phoneNumber",phoneNumber).Count();err!=nil {
		return false
	} else {
		return i == 0
	}
}

//检查用户邮箱 是否存在 ,存在则返回 false ,否则 返回 true
func CheckEmail(email string) bool  {
	if i,err := table.Where("email",email).Count();err!=nil {
		return false
	} else {
		return i == 0
	}
}

func Activated( valida string) bool {

}

