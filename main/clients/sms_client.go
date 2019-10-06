package clients

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type SmsClient struct{}

func NewSmsClient() SmsClient {
	return SmsClient{}
}

func (SmsClient) SendOtpSms(otp string, number string) error {
	message := fmt.Sprintf("This is a test message being sent using Exotel with a (%s) and (%d). If this is being abused, report to 08088919888", otp, 1)
	requestData := createSmsRequest(number, message)

	exotelApiKey := "cbc479b06e98e31fd741413f4a78b521abfb4c39c5030859"
	exotelApiToken := "d6cb2e61f39046bbe32327665c2e7b3afe46a0a45137002a"
	exotelSid := "mitultest1"

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
	data.Set("From", "080-471-12716")
	data.Set("To", to)
	data.Set("Body", body)
	data.Set("Priority", "high")

	return data
}
