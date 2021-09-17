/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package serverstats

import "time"

type StatsByDay struct {
	Date        time.Time             `json:"date"`
	HourlyStats StatsByHourCollection `json:"hourlyStats"`
}

type StatsByDayCollection []*StatsByDay

func NewStatsByDay(date time.Time) *StatsByDay {
	return &StatsByDay{
		Date:        date,
		HourlyStats: make(StatsByHourCollection, 0),
	}
}
