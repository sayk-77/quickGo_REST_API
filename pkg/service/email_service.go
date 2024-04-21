package service

import (
	"example.com/go/models"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strconv"
)

type EmailService struct {
	smtpHost string
	smtpPort string
	from     string
	password string
}

func NewEmailService(smtpHost, smtpPort, from, password string) *EmailService {
	return &EmailService{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		from:     from,
		password: password,
	}
}

func (es *EmailService) formatMessage(email *models.Email) []byte {
	htmlTemplate, err := ioutil.ReadFile("./templates/email.html")
	if err != nil {
		fmt.Printf("Ошибка чтения файла email.html: %s\n", err)
		return nil
	}

	htmlBody := fmt.Sprintf(string(htmlTemplate), email.Customer, email.Question, email.Solution, email.NameEmploy)

	body := []byte("From: " + email.NameEmploy + " <" + es.from + ">\r\n" +
		"To: " + email.Email + "\r\n" +
		"Subject: Ответ на ваш вопрос\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		htmlBody)

	return body
}

func (es *EmailService) formatRecoveryMail(code int, email string) []byte {
	htmlTemplate, err := ioutil.ReadFile("./templates/recovery.html")
	if err != nil {
		fmt.Printf("Ошибка чтения файла recovery.html: %s\n", err)
		return nil
	}

	htmlBody := fmt.Sprintf(string(htmlTemplate), strconv.Itoa(code))

	body := []byte("From: " + "QuickGo" + " <" + es.from + ">\r\n" +
		"To: " + email + "\r\n" +
		"Subject: Код подтверждения почты\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		htmlBody)

	return body
}

func (es *EmailService) SendMail(email *models.Email) error {
	message := es.formatMessage(email)

	auth := smtp.PlainAuth("", es.from, es.password, es.smtpHost)

	err := smtp.SendMail(es.smtpHost+":"+es.smtpPort, auth, es.from, []string{email.Email}, message)
	if err != nil {
		return err
	}

	return nil
}

func (es *EmailService) SendRecoveryMail(code int, email string) error {
	message := es.formatRecoveryMail(code, email)

	auth := smtp.PlainAuth("", es.from, es.password, es.smtpHost)

	err := smtp.SendMail(es.smtpHost+":"+es.smtpPort, auth, es.from, []string{email}, message)
	if err != nil {
		return err
	}

	return nil
}

func (es *EmailService) formatCreateOrderMessage(order *models.Order, emailAddress string) []byte {
	htmlTemplate, err := ioutil.ReadFile("./templates/order_create.html")
	if err != nil {
		fmt.Printf("Ошибка чтения файла order_create.html: %s\n", err)
		return nil
	}

	htmlBody := fmt.Sprintf(string(htmlTemplate), strconv.Itoa(int(order.ID)), order.CargoType.TypeName, order.DestinationAddress, order.Recipient, order.SendingAddress, strconv.Itoa(order.OrderPrice))

	body := []byte("From: QuickGo <" + es.from + ">\r\n" +
		"To: " + emailAddress + "\r\n" +
		"Subject: Оповещение о заказе\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		htmlBody)

	return body
}

func (es *EmailService) SendOrderCreateMail(order *models.Order, emailAddress string) error {
	message := es.formatCreateOrderMessage(order, emailAddress)

	auth := smtp.PlainAuth("", es.from, es.password, es.smtpHost)

	err := smtp.SendMail(es.smtpHost+":"+es.smtpPort, auth, es.from, []string{emailAddress}, message)
	if err != nil {
		return err
	}

	return nil
}

func (es *EmailService) formatConfirmOrderMessage(orderId int, dateSend string, dateDelivery string, driverName string, emailAddress string) []byte {
	htmlTemplate, err := ioutil.ReadFile("./templates/order_confirm.html")
	if err != nil {
		fmt.Printf("Ошибка чтения файла order_confirm.html: %s\n", err)
		return nil
	}

	htmlBody := fmt.Sprintf(string(htmlTemplate), strconv.Itoa(orderId), dateDelivery, dateSend, driverName)

	body := []byte("From: QuickGo <" + es.from + ">\r\n" +
		"To: " + emailAddress + "\r\n" +
		"Subject: Обновление статуса заказа\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		htmlBody)

	return body
}

func (es *EmailService) SendConfirmOrderMail(orderId int, dateSend string, dateDelivery string, driverName string, emailAddress string) error {
	message := es.formatConfirmOrderMessage(orderId, dateDelivery, dateSend, driverName, emailAddress)

	auth := smtp.PlainAuth("", es.from, es.password, es.smtpHost)

	err := smtp.SendMail(es.smtpHost+":"+es.smtpPort, auth, es.from, []string{emailAddress}, message)
	if err != nil {
		return err
	}

	return nil
}
