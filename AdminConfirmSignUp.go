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

	userName := "test_AdminConfirmSignUp"
	password := "test_AdminConfirmSignUp"
	clientId := os.Getenv("CLIENT_ID")
	userPoolId := os.Getenv("USER_POOL_ID")
	email := os.Getenv("EMAIL")
	//phone_number := os.Getenv("PHONE_NUMBER")

	mySession := session.Must(session.NewSessionWithOptions(session.Options{Profile: "default"}))
	svc := cognitoidentityprovider.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1"))

	// SignUp
	paramsSignUp := &cognitoidentityprovider.SignUpInput{
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

	respSignUp, errSignUp := svc.SignUp(paramsSignUp)
	if errSignUp != nil {
		fmt.Println(errSignUp.Error())
		return
	}
	fmt.Println(respSignUp)
	fmt.Println("Completed")

	// Confirm
	paramsConfirm := &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(userName),
	}

	respConfirm, errConfirm := svc.AdminConfirmSignUp(paramsConfirm)
	if errConfirm != nil {
		fmt.Println(errConfirm.Error())
		return
	}
	fmt.Println(respConfirm)
	fmt.Println("Completed")
}
