package handler

import (
	"context"
	"fmt"
	"log"
	"scf/config"
	"scf/pkg"
	"time"
)

var (
	hour8   = 60 * 60 * 8
	localCN = time.FixedZone("UTC+8", hour8)
	aDay    = time.Hour * 24
	aWeek   = aDay * 7
)

func nowCN() time.Time {
	return time.Now().UTC().In(localCN)
}

func remindCurrentWeekOwner(name, mobile, name2, mobile2 string) {
	var (
		title  = "本周周会信息:\n"
		holder = fmt.Sprintf("主持人 - %s\n请整理好周会所需文档, 更新群公告\n", name)
		todo   = fmt.Sprintf("\n土豆记录 - %s\n请预定下周周会会议室\n", name2)
		mob    = []string{mobile}
	)
	msg := fmt.Sprintf("%s\n%s", title, holder)
	if nowCN().Weekday() == time.Thursday {
		msg += todo
		mob = append(mob, mobile2)
	}
	msg += "\n"
	log.Println(msg)
	return
	pkg.WeComWebHookTextMsg(config.Config.BilibiliWeekly.RobotURL, msg, nil, mob)
}

func BilibiliWeeklyRemind(ctx context.Context, event interface{}) (resp interface{}, err error) {
	const offset = 3

	day20210101 := time.Date(2021, time.January, 4, 0, 0, 0, 0, localCN)
	now := time.Since(day20210101)
	weekSinceDay20210101 := int(now.Hours() / aWeek.Hours())

	// current week user
	index := (weekSinceDay20210101 + offset) % len(config.Config.BilibiliWeekly.Users)
	user := config.Config.BilibiliWeekly.Users[index]

	// next week user
	nextIndex := (index + 1) % len(config.Config.BilibiliWeekly.Users)
	nextUser := config.Config.BilibiliWeekly.Users[nextIndex]

	// send remind wecom msg
	remindCurrentWeekOwner(user.Name, user.Mobile, nextUser.Name, nextUser.Mobile)

	resp = event
	return
}
