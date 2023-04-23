package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/leaderboard", getLeaderboardByPage)

	app.Listen(":8080")
}

func getLeaderboardByPage(c *fiber.Ctx) error {

	value := c.Query("page")
	intPage, _ := strconv.Atoi(value)

	for _, a := range leaderboard {
		if a.Page == intPage {
			return c.Status(200).JSON(a)
		}
	}

	return c.Status(404).JSON("Leaderboard Page Not Found")
}

type leaderboardPlayerData struct {
	Rank     int    `json:"rank"`
	Nickname string `json:"nickname"`
	Score    int    `json:"score"`
}

type leaderboardPageData struct {
	Page   int                     `json:"page"`
	IsLast bool                    `json:"is_last"`
	Data   []leaderboardPlayerData `json:"data"`
}

var leaderboard = []leaderboardPageData{
	{
		Page: 0, IsLast: false,
		Data: []leaderboardPlayerData{
			{Rank: 1, Nickname: "dberwick0", Score: 1},
			{Rank: 2, Nickname: "fblakebrough1", Score: 2},
			{Rank: 3, Nickname: "ewisbey2", Score: 3},
			{Rank: 4, Nickname: "mflay3", Score: 4},
			{Rank: 5, Nickname: "bpray4", Score: 5},
			{Rank: 6, Nickname: "ncrack5", Score: 6},
			{Rank: 7, Nickname: "tgiacomuzzi6", Score: 7},
			{Rank: 8, Nickname: "ocaesman7", Score: 8},
			{Rank: 9, Nickname: "psteggals8", Score: 9},
			{Rank: 10, Nickname: "mbracken9", Score: 10},
		},
	},
	{
		Page: 1, IsLast: true,
		Data: []leaderboardPlayerData{
			{Rank: 11, Nickname: "dwheatleya", Score: 11},
			{Rank: 12, Nickname: "lspohrmannb", Score: 12},
			{Rank: 13, Nickname: "lpinsentc", Score: 13},
			{Rank: 14, Nickname: "molifauntd", Score: 14},
			{Rank: 15, Nickname: "kplumptree", Score: 15},
			{Rank: 16, Nickname: "gnouryf", Score: 16},
			{Rank: 17, Nickname: "pkirmang", Score: 17},
			{Rank: 18, Nickname: "jbaddileyh", Score: 18},
			{Rank: 19, Nickname: "kchristofori", Score: 19},
			{Rank: 20, Nickname: "oapplebeej", Score: 20},
		},
	},
}
