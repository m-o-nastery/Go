package main

import (
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User string
	Status string
	Date time.Time
}

var statuses map[string]int = map[string]int{
	"Готово": 0,
	"В работе": 1,
	"Не будет сделано": 2,
}

func GetTasks(text string, user *string, status *string) []Ticket {
	tickets := []Ticket{}
	for _, line := range strings.Split(text, "\n") {
		words := strings.Split(line, "_")
		if len(words) != 4 {
			continue
		}
		tick, name, stat, date := words[0], words[1], words[2], words[3]
		day, err := time.Parse("2006-01-02", date)
		_, ok := statuses[stat]
		if strings.Contains(tick, "TICKET-") == false || err != nil ||  ok == false{
			continue
		}
		
		ticket := Ticket{
			Ticket: tick,
			User: name,
			Status: stat,
			Date: day,
		}

		switch {
		case user != nil && status != nil:
			if ticket.User == *user && ticket.Status == *status {
				tickets = append(tickets, ticket)
			} else {
				continue
			}
		case user != nil:
			if ticket.User == *user {
				tickets = append(tickets, ticket)
			} else {
				continue
			}
		case status != nil:
			if ticket.Status == *status {
				tickets = append(tickets, ticket)
			} else {
				continue
			}
		default:
			tickets = append(tickets, ticket)
		}
	}
	return tickets
}