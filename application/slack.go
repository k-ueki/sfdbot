package application

import (
	"fmt"

	"github.com/k-ueki/sfdbot/config"
	"github.com/k-ueki/sfdbot/util"
)

func SendToSlack(body string) error {
	slackUrl := config.Config.SlackURL
	if err := util.HttpPost(slackUrl, convertStringToSlackJson("<!here>\n"+body)); err != nil {
		return err
	}
	return nil
}

func convertStringToSlackJson(str string) string {
	return fmt.Sprintf(`{"text":"%s"}`, str)
}
