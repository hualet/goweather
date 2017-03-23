package goweather

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrParseTemperature = errors.New("failed to parse temperature")
)

type WeatherData struct {
	Date           time.Time
	MinTemperature int
	MaxTemperature int
	Code           string
	Description    string
	CDescription   string
}

type WeatherInfo struct {
	City string
	// UpdateTime time.Time
	Weather []WeatherData
}

type Manager struct {
	httpClient *http.Client
}

func NewManager() *Manager {
	var netClient = &http.Client{
		Timeout: time.Second * 5,
	}
	return &Manager{httpClient: netClient}
}

func (m *Manager) Fetch(areaID int32) (info *WeatherInfo, err error) {
	url := fmt.Sprintf("http://m.weather.com.cn/mweather/%v.shtml", areaID)
	response, err := m.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	info, err = parseContent(response)
	if err != nil {
		return nil, err
	}

	return info, nil
}
