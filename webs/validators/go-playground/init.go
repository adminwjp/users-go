package go_playgrounds

import (
	"fmt"
	"github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	//"github.com/go-playground/validator/v10"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
	ut "github.com/go-playground/universal-translator"
	//"github.com/go-playground/validator/v10"

	//"github.com/go-playground/validator/v10"
	)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
	trans ut.Translator
)
func Init(locale string)  {

	en := en.New()
	zh := zh.New()
	zh_tw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zh_tw)
	Validate = validator.New()

	trans, _ := Uni.GetTranslator(locale)
	switch locale {
	case "zh":
		zh_translations.RegisterDefaultTranslations(Validate, trans)
		break
	case "en":
		en_translations.RegisterDefaultTranslations(Validate, trans)
		break
	case "zh_tw":
		zh_tw_translations.RegisterDefaultTranslations(Validate, trans)
		break
	default:
		zh_translations.RegisterDefaultTranslations(Validate, trans)
		break
	}

	//自定义错误内容
	Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}


func ValidateError(httpWeb webs.HttpWeb,obj interface{})bool  {
	err := Validate.Struct(obj)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		httpWeb.Response(200, dtos.ResponseDto{Status: false,
			Msg: fmt.Sprintf("%#v", sliceErrs),Code: 400})
		return false
	}
	return  true
}
