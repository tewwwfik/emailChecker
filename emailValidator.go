package main

import (
	"net"
	"net/mail"
	"strings"
)

func CheckSingleEmail(mail string) (email string, isLexicalValid, isMxValid bool) {
	email = fixEmail(mail)
	return email, lexicalEmailValidator(email), checkMx(email)
}

//Returns validMails, lexicalInvalidMails, mxInvalidMails
func GetValidEmails(emails []string) (validMails, lexicalInvalidMails, mxInvalidMails []string) {
	for _, email := range emails {
		email = fixEmail(email)
		if lexicalEmailValidator(email) {
			if checkMx(email) {
				validMails = append(validMails, email)
			} else {
				mxInvalidMails = append(mxInvalidMails, email)
			}
		} else {
			lexicalInvalidMails = append(lexicalInvalidMails, email)
		}
	}
	return validMails, lexicalInvalidMails, mxInvalidMails
}

func fixEmail(email string) string {
	return strings.Replace(strings.ToLower(email), " ", "", -1)
}

func lexicalEmailValidator(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func checkMx(email string) bool {
	domain := email[strings.Index(email, "@")+1 : len(email)]
	mxrecords, _ := net.LookupMX(domain)

	//for _, mx := range mxrecords {
	//	fmt.Println(mx.Host, mx.Pref)
	//}

	if len(mxrecords) > 0 {
		return true
	}

	return false
}
