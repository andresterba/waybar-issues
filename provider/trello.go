package provider

import (
	"fmt"
	"time"

	"github.com/adlio/trello"
)

type trelloStats struct {
	appKey        string
	authToken     string
	todayDueCards int
	displayName   string
	boardID       string
}

func NewTrelloStats(appKey, authToken, boardID, displayName string) *trelloStats {

	// client := trello.NewClient(config.AppKey, config.Token)

	return &trelloStats{
		appKey:      appKey,
		authToken:   authToken,
		boardID:     boardID,
		displayName: displayName,
	}
}

func (ts *trelloStats) getDueTodayCards(trelloClient *trello.Client) (int, error) {
	board, err := trelloClient.GetBoard(ts.boardID, trello.Defaults())
	if err != nil {
		return 0, fmt.Errorf("could not find board with ID %s", ts.boardID)
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return 0, err
	}

	var cardsDueToday []*trello.Card

	for _, card := range cards {
		if isDueSetOnCard(card) && isCardDueToday(*card.Due) && !isCardDueCompleted(card) {
			cardsDueToday = append(cardsDueToday, card)
		}
	}

	return len(cardsDueToday), nil
}

func (ts *trelloStats) Process() error {
	client := trello.NewClient(ts.appKey, ts.authToken)

	dueCards, err := ts.getDueTodayCards(client)
	if err != nil {
		return err
	}

	ts.todayDueCards = dueCards

	return nil
}

func (ts *trelloStats) GetFormatedOutput() string {
	return fmt.Sprintf("%s: %d ", ts.displayName, ts.todayDueCards)
}

func isDueSetOnCard(card *trello.Card) bool {
	return card.Due != nil
}

func isCardDueCompleted(card *trello.Card) bool {
	return card.DueComplete
}

func isCardDueToday(dueTime time.Time) bool {
	year, month, day := dueTime.Date()
	yearNow, monthNow, dayNow := time.Now().Date()

	if (year == yearNow) && (month == monthNow) && (day == dayNow) {
		return true
	}

	return false
}
