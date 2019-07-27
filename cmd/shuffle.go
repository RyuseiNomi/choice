package cmd

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/spf13/cobra"
)

type Member struct {
	Name  string
	Group int
}

type Members []Member

func shuffle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shuffle",
		Short: "shuffle member",
		Run: func(cmd *cobra.Command, args []string) {
			groupNumber, _ := strconv.Atoi(args[0])
			members := args[1:]
			doShuffle(groupNumber, members)
		},
	}

	return cmd
}

func doShuffle(groupNumber int, inputedMembers []string) {
	var members Members
	groupMaxCount := len(inputedMembers) / groupNumber

	//membersの長さ分forを回して全員に番号をランダムで振る
	for i := range inputedMembers {
		name := inputedMembers[i]

		var member = Member{
			Name:  name,
			Group: rand.Intn(groupMaxCount),
		}

		members = append(members, member)
	}

	fmt.Println(members)
}
