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

	userName := "test_AdminCreateUser"
	userPoolId := os.Getenv("USER_POOL_ID")
	email := os.Getenv("EMAIL")
	//phone_number := os.Getenv("PHONE_NUMBER")

	mySession := session.Must(session.NewSessionWithOptions(session.Options{Profile: "default"}))
	svc := cognitoidentityprovider.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1"))

	params := &cognitoidentityprovider.AdminCreateUserInput{
		DesiredDeliveryMediums: []*string{
			aws.String("EMAIL"),
			//aws.String("SMS"),
		},
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
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(userName),
	}

	resp, err := svc.AdminCreateUser(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
	fmt.Println("Completed")
}
