package pleasanter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	getPath = "%v/api/items/%v/get"
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

func getItems(c *Client, t string, o int64, v *View) (*ItemResponse, error) {
	rb := ItemRequest{
		requestBase: c.requestBase,
		Offset:      o,
		View:        v,
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
	var r ItemResult
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

	return r.Response, nil
}

func (c *Client) UpdateItem(itemID string, itemData ItemData) (ItemUpdateResponse, error) {
	rb := ItemUpdateRequest{
		c.requestBase,
		itemData,
	}
	var iur ItemUpdateResponse
	s, err := json.Marshal(rb)
	if err != nil {
		return iur, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/api/items/%v/update", c.endpoint, itemID), strings.NewReader(string(s)))
	if err != nil {
		return iur, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return iur, err
	}

	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != http.StatusOK {
		var er ErrorResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(er.Message)
		}
	} else {
		err = decoder.Decode(&iur)
	}
	if err != nil {
		return iur, err
	}

	return iur, nil

}

func (c *Client) GetItemByID(itemID string) (*ItemData, error) {
	r, err := getItems(c, itemID, 0, nil)
	if err != nil {
		return nil, err
	}
	var i *ItemData
	for _, v := range r.Data {
		i = &v
	}
	return i, nil
}

func (c *Client) GetItems(tableID string, filter *View) ([]ItemData, error) {
	var itemData []ItemData
	var offset int64
	completed := false
	encountered := map[int]bool{}

	for completed == false {
		r, err := getItems(c, tableID, offset, filter)
		if err != nil {
			return nil, err
		}
		if itemData == nil {
			itemData = make([]ItemData, 0, r.TotalCount)
		}
		for _, d := range r.Data {
			if !encountered[d.IssueID] {
				encountered[d.IssueID] = true
				itemData = append(itemData, d)
			}
		}

		if r.TotalCount <= len(itemData) {
			completed = true
			continue
		}
		offset = int64(len(itemData) / r.PageSize)
	}
	return itemData, nil
}

type ItemResponse struct {
	Offset     int        `json:"Offset,omitempty"`
	PageSize   int        `json:"PageSize,omitempty"`
	TotalCount int        `json:"TotalCount,omitempty"`
	Data       []ItemData `json:"Data,omitempty"`
}

type ItemData struct {
	SiteID             int              `json:"SiteId,omitempty"`
	UpdatedTime        string           `json:"UpdatedTime,omitempty"`
	IssueID            int              `json:"IssueId,omitempty"`
	Ver                int              `json:"Ver,omitempty"`
	Title              string           `json:"Title,omitempty"`
	Body               string           `json:"Body,omitempty"`
	StartTime          string           `json:"StartTime,omitempty"`
	CompletionTime     string           `json:"CompletionTime,omitempty"`
	WorkValue          float64          `json:"WorkValue,omitempty"`
	ProgressRate       float64          `json:"ProgressRate,omitempty"`
	RemainingWorkValue float64          `json:"RemainingWorkValue,omitempty"`
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

type ItemUpdateRequest struct {
	requestBase
	ItemData
}

type ItemUpdateResponse struct {
	ID             int    `json:"Id"`
	StatusCode     int    `json:"StatusCode"`
	LimitPerDate   int    `json:"LimitPerDate"`
	LimitRemaining int    `json:"LimitRemaining"`
	Message        string `json:"Message"`
}
