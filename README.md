# Email Checker
Email address validator for golang. Can make lexical validation and test whether domain exists.

It can both checks single or multiple emails.

# Usage
Multiple mails:
```example
emails := []string{
		"good@exmaple.com  ",
		"bad-example@ There Is No Domain Like That.com",
		"bad-example",
		"test@domain.com",
	}
        //Returns validmails,lexicalInvalidMails and domainInvalidMails as string slice
	validmails, lexicalInvalidMails, domainInvalidMails := GetValidEmails(emails)
```
Single Mail:
```example2
//It returns corrected string without empty space

//mail string, lexical boolean, domain boolean
mail, lexical, domain := CheckSingleEmail("abidik Gubidik@ whatThe Fuck.com")
```
You can inspect example.go for more instructions and usage.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
