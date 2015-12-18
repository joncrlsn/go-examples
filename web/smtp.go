//
// Thank you to William Kennedy who shared his code here:
// http://www.goinggo.net/2013/06/send-email-in-go-with-smtpsendmail.html
//
// Example call:
//    SendEmail(
//        "smtp.gmail.com",
//        587,
//        "username@gmail.com",
//        "password",
//        []string{"me@domain.com"},
//        "testing subject",
//        "<html><body>Exception 1</body></html>Exception 1")
//    }
//
package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"runtime"
	"strings"
	"text/template"
)

var _emailScript = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"

{{.Message}}`

// sendMail sends an email with values from the config file
func sendMail(subject, message string) error {
	return _sendEmail(mailHost, mailPort, mailUsername, mailPassword, mailFrom, mailTo, subject, message)
}

// _sendEmail does the detailed-work for sending an email
func _sendEmail(host string, port int, userName string, password string, from string, to []string, subject string, message string) (err error) {
	defer _catchPanic(&err, "_sendEmail")

	if len(host) == 0 {
		if verbose {
			fmt.Println("No smtp host.  Skipping email.")
		}
		return nil
	}

	parameters := struct {
		From    string
		To      string
		Subject string
		Message string
	}{
		userName,
		strings.Join([]string(to), ","),
		subject,
		message,
	}

	buffer := new(bytes.Buffer)

	template := template.Must(template.New("emailTemplate").Parse(_emailScript))
	template.Execute(buffer, &parameters)

	auth := smtp.PlainAuth("", userName, password, host)
	if len(userName) == 0 {
		auth = nil
	}

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		from,
		to,
		buffer.Bytes())

	return err
}

func _catchPanic(err *error, functionName string) {
	if r := recover(); r != nil {
		fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	} else if err != nil && *err != nil {
		fmt.Printf("%s : ERROR : %v\n", functionName, *err)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))
	}
}
