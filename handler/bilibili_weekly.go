package handler

import (
	"context"
	"fmt"
	"log"
	"scf/config"
	"scf/pkg"
	"time"
)

func remindCurrentWeekOwner(name, mobile string) {
	msg := fmt.Sprintf("亲爱的 [周会主持人 - %s]:\n今天是星期三, 要准备周会啦", name)
	log.Println(msg)
	pkg.WeComWebHookTextMsg(config.Config.BilibiliWeekly.RobotURL, msg, nil, []string{mobile})
}

func remindNextWeekOwner(name, mobile string) {
	msg := fmt.Sprintf("亲爱的 [土豆记录小助手 - %s]:\n今天是星期三, 要准备周会啦", name)
	log.Println(msg)
	pkg.WeComWebHookTextMsg(config.Config.BilibiliWeekly.RobotURL, msg, nil, []string{mobile})
}

func BilibiliWeeklyRemind(ctx context.Context, event interface{}) (resp interface{}, err error) {
	const offset = 12

	day20210101 := time.Date(2021, time.January, 0, 0, 0, 0, 0, time.UTC)
	now := time.Since(day20210101)
	aWeek := time.Hour * 24 * 7
	weekSinceDay20210101 := int(now.Hours() / aWeek.Hours())

	// current week user
	index := (weekSinceDay20210101 + offset) % len(config.Config.BilibiliWeekly.Users)
	user := config.Config.BilibiliWeekly.Users[index]

	// next week user
	nextIndex := (index + 1) % len(config.Config.BilibiliWeekly.Users)
	nextUser := config.Config.BilibiliWeekly.Users[nextIndex]

	// send remind wecom msg
	remindCurrentWeekOwner(user.Name, user.Mobile)
	remindNextWeekOwner(nextUser.Name, nextUser.Mobile)

	resp = event
	return
}
