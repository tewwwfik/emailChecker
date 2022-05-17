package main

import "fmt"

func main() {
	emails := []string{
		"good@exmaple.com  ",
		"bad-example@ There Is No Domain Like That.com",
		"bad-example",
		"test@domain.com",
	}
	validmails, lexicalInvalidMails, domainInvalidMails := GetValidEmails(emails)
	fmt.Println("--------Multiple Mails Check--------")
	fmt.Println("Valid Mails->")
	for _, validmail := range validmails {
		fmt.Println(validmail)
	}
	fmt.Println("Lexical Invalid Mails->")
	for _, mail := range lexicalInvalidMails {
		fmt.Println(mail)
	}
	fmt.Println("Domain name not exist->")
	for _, mail := range domainInvalidMails {
		fmt.Println(mail)
	}
	fmt.Println("--------Single Mail Check--------")
	mail, lexical, domain := CheckSingleEmail("abidik Gubidik@ whatThe Fuck.com")
	fmt.Printf("mail: %s, lexical: %t, domain: %t", mail, lexical, domain)

}
