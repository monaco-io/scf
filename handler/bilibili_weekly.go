package handler

import (
	"context"
	"fmt"
	"scf/config"
	"scf/pkg"
	"time"
)

func remind(name, mobile string) {
	msg := fmt.Sprintf("亲爱的%s, 今天是星期四, 要准备周会啦", name)
	pkg.WeComWebHookTextMsg(config.Config.BilibiliWeekly.RobotURL, msg, nil, []string{mobile})
}

func BilibiliWeeklyRemind(ctx context.Context, event interface{}) (resp interface{}, err error) {
	resp = event
	day20210101 := time.Date(2021, time.January, 0, 0, 0, 0, 0, time.UTC)
	now := time.Since(day20210101)
	aWeek := time.Hour * 24 * 7
	weekSinceDay20210101 := int(now.Hours() / aWeek.Hours())
	user := config.Config.BilibiliWeekly.Users[weekSinceDay20210101%len(config.Config.BilibiliWeekly.Users)]
	remind(user.Name, user.Mobile)
	return
}
