package roomcmds

import (
	"fmt"
	"github.com/Diarkis/diarkis/log"
	"github.com/Diarkis/diarkis/matching"
	"github.com/Diarkis/diarkis/user"
	"github.com/Diarkis/diarkis/room"
)

var logger = log.New("matching")

func Expose(rootpath string) {
	// Set up matching package and load configuration file
	matching.Setup(fmt.Sprintf("%s/configs/shared/matching.json", rootpath))

	// Set up matching ticket
	matching.SetOnIssueTicket(func(userData *user.User) *matching.TicketParams {
		return &matching.TicketParams{
			ProfileIDs:     []string{"RankMatch", "RankMatch2"},
			MaxMembers:     2,
			SearchInterval: 100, // 100ms
			TicketDuration: 60,  // 1m
			HowMany:        20,
			// Change here as you see fit according to your application needs
			Tag: "",
			// Change here as you see fit according to your application needs
			Properties: map[string]int{"rank": 1},
		}
	})
	matching.SetOnTicketMatch(func(userData *user.User, roomID string, memberIDs []string) bool {
		// add custom logic to decide matchmaking completion here
		return false
	})
	matching.SetOnTicketComplete(func(roomID string, userData *user.User) []byte {
		memberIDs := room.GetMemberIDs(roomID)
		return []byte(fmt.Sprintf("Ticket matchmaking complete => %v", memberIDs))
	})

	// Expose built-in commands to the client
	matching.ExposeCommands()
}