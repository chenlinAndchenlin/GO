package forms

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式规范，需要自定义validata
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
