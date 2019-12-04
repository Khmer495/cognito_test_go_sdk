package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	Env_load()

	userName := "test_SignUp"
	password := "test_SignUp"
	clientId := os.Getenv("CLIENT_ID")
	email := os.Getenv("EMAIL")
	//phone_number := os.Getenv("PHONE_NUMBER")

	mySession := session.Must(session.NewSession())
	svc := cognitoidentityprovider.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1"))

	params := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(clientId),
		Password: aws.String(password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			//{
			//Name:  aws.String("phone_number"),
			//Value: aws.String(phone_number),
			//},
		},
		Username: aws.String(userName),
	}

	resp, err := svc.SignUp(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
	fmt.Println("Completed")
}
