package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"sort"
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

func doShuffle(groupNumber int, inputtedMembers []string) Members {
	var members Members
	shuffleMembers := inputtedMembers

	rand.Seed(time.Now().UnixNano()) //乱数の初期化
	rand.Shuffle(len(shuffleMembers), func(i, j int) {
		shuffleMembers[i], shuffleMembers[j] = shuffleMembers[j], shuffleMembers[i]
	})
	for i := range shuffleMembers {
		number := i % groupNumber
		members = append(members, Member{
			Name:  shuffleMembers[i],
			Group: number,
		})
	}

	return members
}
