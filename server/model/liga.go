package model

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	_ "cicada/server/util"

	"github.com/PuerkitoBio/goquery"
)

// Schedule 赛程
type Schedule struct {
	Name    string   // 名字，如：liga，表示西甲
	Kind    string   // 类型，如：联赛(T)/杯赛(C)
	TeamURL string   // 球队信息链接，如：http://www.laliga.es/en/laliga-santander
	Result  []Round  // 所有比赛
	Teams   []string // 所有队伍，用来缓存
}

// Round 回合
type Round struct {
	Version int             // 回合数
	Match   []Confrontation // 比赛列表
}

// Confrontation 对阵
type Confrontation struct {
	Home  string // 主队
	Guest string // 客队
}

// NewSchedule 创建赛程
func NewSchedule(name, kind, url string, teams []string) *Schedule {
	s := new(Schedule)
	s.Name = name
	s.Kind = kind
	s.TeamURL = url
	s.Teams = teams

	// 制定赛程
	err := s.makeSchedule()
	if err != nil {
		panic(err)
	}

	return s
}

// 寻找所有队伍
func (s *Schedule) findAllTeam() (teams []string, err error) {
	// 请求
	res, err := http.Get(s.TeamURL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	// 加载 html 文档
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	// 寻找队伍
	doc.Find(".container #equipos .columna").Each(func(i int, s *goquery.Selection) {
		s.Find(".nombre").Each(func(i int, s *goquery.Selection) {
			team := s.Text()
			teams = append(teams, team)
		})
	})

	return
}

func (s *Schedule) makeSchedule2() (err error) {
	if len(s.Teams) == 0 {
		// 寻找队伍
		var teams []string
		teams, err = s.findAllTeam()
		if err != nil {
			return
		}
		s.Teams = teams
	}
	tl := len(s.Teams)
	if tl < 2 {
		err = fmt.Errorf("不能少于两队：%d", tl)
		return
	}

	// 制定赛程
	var homeTeam = make([]string, tl) // 主队
	copy(homeTeam, s.Teams)
	var guestTeam = make([]string, tl) // 客队
	copy(guestTeam, s.Teams)

	var homeMap = make(map[string]int)  // 球队在主场的次数
	var guestMap = make(map[string]int) // 球队在客场的次数
	var confMap = make(map[string]int)  // 两只球队相遇的次数，最大为1，如：{FC Barcelona-R. Madrid: 1, R. Madrid-FC Barcelona: 1}

	roundNum := (tl - 1) * 2         // 回合数
	for i := 1; i <= roundNum; i++ { // 逐个回合
		var leftHomeTeam = make([]string, len(homeTeam)) // 主队
		copy(leftHomeTeam, homeTeam)
		var leftGuestTeam = make([]string, len(guestTeam)) // 客队
		copy(leftGuestTeam, guestTeam)

		round := Round{Version: i}
		for j := 1; j <= tl/2; j++ {
			// log.Printf("%d, %d, %v\n%v\n%v\n%v\n", i, j, leftHomeTeam, homeTeam, leftGuestTeam, guestTeam)
			conf := Confrontation{}

			var foundConf bool
			for !foundConf {
				// 随机选一支队伍，作为主队
				if len(leftHomeTeam) == 0 {
					break
				}
				homeIndex := rand.Intn(len(leftHomeTeam)) // 选出主队
				conf.Home = leftHomeTeam[homeIndex]       // 记录主队
				// 检查该队的主场数
				if v, ok := homeMap[conf.Home]; ok {
					if v >= tl-1 { // 如果数量已达到，另外找一支
						continue
					}
				}

				var found bool
				var guestIndex int
				for !found {
					guestIndex = rand.Intn(len(leftGuestTeam)) // 选出客队
					conf.Guest = leftGuestTeam[guestIndex]     // 记录客队
					if conf.Guest == conf.Home {
						continue
					}
					if v, ok := guestMap[conf.Guest]; ok {
						if v >= tl-1 { // 如果数量已达到，再找一支
							continue
						}
					}
					found = true
				}

				confKey := conf.Home + "-" + conf.Guest
				if v, ok := confMap[confKey]; ok {
					if v >= 1 {
						continue
					}
				}
				// log.Printf("%d, %d, %v\n", i, j, confMap)
				// log.Printf("%d, %d, %v\n", i, j, confKey)
				foundConf = true
				round.Match = append(round.Match, conf)
				leftHomeTeam = append(leftHomeTeam[:homeIndex], leftHomeTeam[homeIndex+1:]...)
				leftGuestTeam = append(leftGuestTeam[:guestIndex], leftGuestTeam[guestIndex+1:]...)
				homeMap[conf.Home]++
				if homeMap[conf.Home] == tl-1 {
					delHomeIndex := -1
					for i, team := range homeTeam {
						if team == conf.Home {
							delHomeIndex = i
							break
						}
					}
					if delHomeIndex != -1 {
						homeTeam = append(homeTeam[:delHomeIndex], homeTeam[delHomeIndex+1:]...) // 从主队列表中删除
					}
				}
				guestMap[conf.Guest]++
				if guestMap[conf.Guest] == tl-1 {
					delGuestIndex := -1
					for i, team := range guestTeam {
						if team == conf.Guest {
							delGuestIndex = i
							break
						}
					}
					if delGuestIndex != -1 {
						guestTeam = append(guestTeam[:delGuestIndex], guestTeam[delGuestIndex+1:]...) // 从客队列表中删除
					}
				}
				confMap[confKey]++

				// fmt.Printf("%v, %v, %v\n%v\n%v\n%v\n", i, j, homeMap, guestMap, homeTeam, guestTeam)
			}
		}

		log.Printf("%d, %v\n", i, round)
		s.Result = append(s.Result, round)
	}

	return
}

func (s *Schedule) makeSchedule3() (err error) {
	if len(s.Teams) == 0 {
		// 寻找队伍
		var teams []string
		teams, err = s.findAllTeam()
		if err != nil {
			return
		}
		s.Teams = teams
	}
	tl := len(s.Teams)
	if tl < 2 {
		err = fmt.Errorf("不能少于两队：%d", tl)
		return
	}

	// 制定赛程
	var confMap = make(map[string]int) // 两只球队相遇的次数，最大为1，如：{FC Barcelona-R. Madrid: 1, R. Madrid-FC Barcelona: 1}

	roundNum := (tl - 1) * 2         // 回合数
	for i := 1; i <= roundNum; i++ { // 逐个回合
		var leftTeam = make([]string, tl)
		copy(leftTeam, s.Teams)

		round := Round{Version: i}
		for j := 1; j <= tl/2; j++ {
			conf := Confrontation{}

			var found bool
			for !found {
				log.Printf("%v\n", leftTeam)
				if len(leftTeam) == 0 {
					break
				}
				// 主队
				homeIndex := rand.Intn(len(leftTeam))
				conf.Home = leftTeam[homeIndex]
				log.Printf("%v\n", conf.Home)
				// 客队
				var foundGuest bool
				var guestIndex int
				for !foundGuest {
					if len(leftTeam) == 0 {
						break
					}
					guestIndex = rand.Intn(len(leftTeam))
					conf.Guest = leftTeam[guestIndex]
					log.Printf("%v\n", conf.Guest)
					if conf.Guest != conf.Home {
						foundGuest = true
					}
				}

				confKey := conf.Home + "-" + conf.Guest
				if v, ok := confMap[confKey]; ok {
					if v >= 1 { // 已经存在，则继续
						log.Printf("%d, %d, %v\n", i, j, confKey)
						continue
					}
				}

				confMap[confKey]++ // 记录对阵
				if homeIndex > guestIndex {
					leftTeam = append(leftTeam[:homeIndex], leftTeam[homeIndex+1:]...)   // 去掉主队
					leftTeam = append(leftTeam[:guestIndex], leftTeam[guestIndex+1:]...) // 去掉客队
				} else {
					leftTeam = append(leftTeam[:guestIndex], leftTeam[guestIndex+1:]...) // 去掉客队
					leftTeam = append(leftTeam[:homeIndex], leftTeam[homeIndex+1:]...)   // 去掉主队
				}
				found = true // 成功找到
				log.Printf("%v, %v\n", leftTeam, confMap)
			}

			round.Match = append(round.Match, conf)
		}

		s.Result = append(s.Result, round)
	}

	return
}

func (s *Schedule) makeSchedule() (err error) {
	if len(s.Teams) == 0 {
		// 寻找队伍
		var teams []string
		teams, err = s.findAllTeam()
		if err != nil {
			return
		}
		s.Teams = teams
	}
	tl := len(s.Teams)
	if tl < 2 {
		err = fmt.Errorf("不能少于两队：%d", tl)
		return
	}

	// 制定赛程
	confMap := make(map[string][]string) // 主队还未抽到的对手

	for i, team := range s.Teams {
		prefixTeam := make([]string, len(s.Teams[:i]))
		suffixTeam := make([]string, len(s.Teams[i+1:]))
		copy(prefixTeam, s.Teams[:i])
		copy(suffixTeam, s.Teams[i+1:])
		opponentTeams := append(prefixTeam, suffixTeam...)
		confMap[team] = opponentTeams
	}

	roundNum := (tl - 1) * 2
	for i := 1; i <= roundNum; i++ {
		leftTeam := make([]string, tl)
		copy(leftTeam, s.Teams)

		round := Round{Version: i}
		for j := 1; j <= tl/2; j++ { // 第i轮，第j场
			conf := Confrontation{}

			for {
				if len(leftTeam) == 0 {
					break
				}
				homeIndex := rand.Intn(len(leftTeam))
				conf.Home = leftTeam[homeIndex]
				if v, ok := confMap[conf.Home]; ok {
					if len(v) == 0 {
						continue
					}
					guestIndex := rand.Intn(len(v)) // 在主队还能抽的客队列表的下标
					conf.Guest = v[guestIndex]
					// 抽出的guest也要在leftTeam里才行
					var found bool
					for _, team := range leftTeam {
						if team == conf.Guest {
							found = true
							break
						}
					}
					if !found {
						continue
					}

					confMap[conf.Home] = append(v[:guestIndex], v[guestIndex+1:]...) // 删除主队还能抽的客队
					log.Printf("%d,%d,%v\n", i, j, conf)
					log.Printf("%v\n", confMap)
					leftTeam = append(leftTeam[:homeIndex], leftTeam[homeIndex+1:]...) // 删除主队
					for k, team := range leftTeam {
						if team == conf.Guest {
							guestIndex = k
							break
						}
					}
					// log.Printf("%v\n", conf)
					leftTeam = append(leftTeam[:guestIndex], leftTeam[guestIndex+1:]...) // 删除客队
					// log.Printf("%v\n", leftTeam)
					break
				}
			}

			round.Match = append(round.Match, conf)
			// log.Printf("%d,%d,%v\n", i, j, conf)
		}

		s.Result = append(s.Result, round)
	}

	return
}

// Refresh 重制赛程
func (s *Schedule) Refresh() (err error) {
	s.Result = []Round{}    // 清空
	return s.makeSchedule() // 重制
}
