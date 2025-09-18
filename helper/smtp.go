package helper

import "fmt"

type SMTP_SUBJECT string
type SMTP_TEMPLATE string

const (
	SMTP_SubjectRegisterSuccess SMTP_SUBJECT = "Register Success"

	SMTP_TemplateRegisterSuccess SMTP_TEMPLATE = "Congratulations %s!\nYou have successfully registered with JIMA!"
	SMTP_TemplateForgotPassword  SMTP_TEMPLATE = "Please access the link below to reset your password.\n %s"
)

func GenerateSMTPTemplate(template SMTP_TEMPLATE, params ...any) string {
	return fmt.Sprintf(string(template), params...)
}
