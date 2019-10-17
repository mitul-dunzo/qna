package clients

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ISmsClient interface {
	SendOtpSms(otp string, number string) error
}

type SmsClient struct{}

func NewSmsClient() ISmsClient {
	return SmsClient{}
}

func (SmsClient) SendOtpSms(otp string, number string) error {
	message := fmt.Sprintf("This is a test message being sent using Exotel with a (%s) and (%d). If this is being abused, report to 08088919888", otp, 1)
	requestData := createSmsRequest(number, message)

	exotelApiKey := os.Getenv("ExotelAPIKey")
	exotelApiToken := os.Getenv("ExotelAPIToken")
	exotelSid := os.Getenv("ExotelSID")

	urlString := fmt.Sprintf("https://%s:%s@api.exotel.com/v1/Accounts/%s/Sms/send.json", exotelApiKey, exotelApiToken, exotelSid)

	req, err := http.NewRequest(http.MethodPost, urlString, strings.NewReader(requestData.Encode()))
	if err != nil {
		logrus.Error("Couldn't create exotel request: ", err.Error())
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Error("Exotel API failed: ", err.Error())
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Exotel API had no body ", err.Error())
		return err
	}
	logrus.Info("Exotel API response: ", string(body))

	if resp.StatusCode != 200 {
		return errors.New("failed to send SMS")
	}
	return nil
}

func createSmsRequest(to string, body string) url.Values {
	data := url.Values{}
	data.Set("From", os.Getenv("ExotelPhoneNumber"))
	data.Set("To", to)
	data.Set("Body", body)
	data.Set("Priority", "high")

	return data
}
