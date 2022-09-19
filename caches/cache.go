package caches

import "github.com/adminwjp/users-go/models"

type ICache interface {
	SaveRole(m  *models.RoleModel)
	SaveEmail(m  *models.EmailConfigModel)

	SaveSmses(ms  []*models.SmsConfigModel)
	SaveSms(m  *models.SmsConfigModel)

	SavePaies(ms  []*models.PaySecrtConfigModel)
	SavePay(m  *models.PaySecrtConfigModel)

}
