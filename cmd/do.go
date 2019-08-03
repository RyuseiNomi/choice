package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"strconv"
	"time"
)

type Member struct {
	Name  string
	Group int
}

type Members []Member

func do() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "do",
		Short: "divide members to group",
		Run: func(cmd *cobra.Command, args []string) {
			groupNumber, _ := strconv.Atoi(args[0])
			members := args[1:]
			doShuffle(groupNumber, members)
		},
	}

	return cmd
}

func doShuffle(groupNumber int, inputtedMembers []string) {
	var members Members
	//groupMaxCount := len(inputtedMembers) / groupNumber

	rand.Seed(time.Now().UnixNano()) //乱数の初期化
	//membersの長さ分forを回して全員に番号をランダムで振る
	for i := range inputtedMembers {
		name := inputtedMembers[i]

		var member = Member{
			Name:  name,
			Group: rand.Intn(groupNumber),
		}

		members = append(members, member)
	}

	fmt.Println(members)
}
