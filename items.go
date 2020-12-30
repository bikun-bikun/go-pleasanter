package pleasanter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ItemRequest struct {
	requestBase
	Offset int64 `json:"Offset,omitempty"`
	View   *View `json:"View,omitempty"`
}

type ItemResult struct {
	StatusCode     int           `json:"StatusCode,omitempty"`
	LimitPerDate   int           `json:"LimitPerDate,omitempty"`
	LimitRemaining int           `json:"LimitRemaining,omitempty"`
	Response       *ItemResponse `json:"Response,omitempty"`
}

func (c *Client) GetItems(t string) (*ItemResponse, error) {
	rb := ItemRequest{
		requestBase: c.requestBase,
		Offset:      0,
		View:        &View{},
	}
	s, err := json.Marshal(rb)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/api/items/%v/get", c.endpoint, t), strings.NewReader(string(s)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	var r ItemResponse
	if res.StatusCode != http.StatusOK {
		var er ErrorResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(er.Message)
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r, nil

}

type ItemResponse struct {
	StatusCode int              `json:"StatusCode,omitempty"`
	Response   ItemResponseBody `json:"Response,omitempty"`
}

type ItemResponseBody struct {
	Offset     int     `json:"Offset,omitempty"`
	PageSize   int     `json:"PageSize,omitempty"`
	TotalCount int     `json:"TotalCount,omitempty"`
	Data       []Datum `json:"Data,omitempty"`
}

type Datum struct {
	APIVersion         float64          `json:"ApiVersion,omitempty"`
	SiteID             int              `json:"SiteId,omitempty"`
	UpdatedTime        string           `json:"UpdatedTime,omitempty"`
	IssueID            int              `json:"IssueId,omitempty"`
	Ver                int              `json:"Ver,omitempty"`
	Title              string           `json:"Title,omitempty"`
	Body               string           `json:"Body,omitempty"`
	StartTime          string           `json:"StartTime,omitempty"`
	CompletionTime     string           `json:"CompletionTime,omitempty"`
	WorkValue          int              `json:"WorkValue,omitempty"`
	ProgressRate       int              `json:"ProgressRate,omitempty"`
	RemainingWorkValue int              `json:"RemainingWorkValue,omitempty"`
	Status             int              `json:"Status,omitempty"`
	Manager            int              `json:"Manager,omitempty"`
	Owner              int              `json:"Owner,omitempty"`
	Comments           string           `json:"Comments,omitempty"`
	Creator            int              `json:"Creator,omitempty"`
	Updator            int              `json:"Updator,omitempty"`
	CreatedTime        string           `json:"CreatedTime,omitempty"`
	ItemTitle          string           `json:"ItemTitle,omitempty"`
	DateHash           *DateHash        `json:"DateHash,omitempty"`
	DescriptionHash    *DescriptionHash `json:"DescriptionHash,omitempty"`
	NumHash            *NumHash         `json:"NumHash,omitempty"`
	CheckHash          *CheckHash       `json:"CheckHash,omitempty"`
	ClassHash          *ClassHash       `json:"ClassHash,omitempty"`
	AttachmentsHash    *AttachmentsHash `json:"AttachmentsHash,omitempty"`
}
