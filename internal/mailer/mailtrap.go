package mailer

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"text/template"

	gomail "gopkg.in/mail.v2"
)

type MailtrapClient struct {
	fromEmail string
	username  string
	password  string
}

func NewMailtrapClient(username, password, fromEmail string) (*MailtrapClient, error) {
	if username == "" || password == "" {
		return nil, errors.New("SMTP credentials are required")
	}

	return &MailtrapClient{
		fromEmail: fromEmail,
		username:  username,
		password:  password,
	}, nil
}

func (m *MailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	// Шаг 1: Парсинг шаблонов
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(body, "body", data); err != nil {
		return -1, err
	}

	// Шаг 2: Создание email-сообщения
	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject.String())
	message.SetBody("text/html", body.String())

	// Шаг 3: Создание SMTP-драйвера
	dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, m.username, m.password)

	// Шаг 4: Повторные попытки отправки (retry + backoff)
	var retryErr error
	for i := 0; i < maxRetries; i++ {
		retryErr = dialer.DialAndSend(message)
		if retryErr != nil {
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}

		return -1, err
	}

	return -1, fmt.Errorf("failed to send email after %d attempt(s), error: %v", maxRetries, retryErr)
}
