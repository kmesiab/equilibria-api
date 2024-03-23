package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	klogger "github.com/kmesiab/go-klogger"

	"github.com/kmesiab/equilibria-api/lib/nrclex"
)

type EmotionsController struct {
	MaxLimit  int
	MaxOffset int

	NrcLexService *nrclex.Service
}

func (ctrl *EmotionsController) NrcLex(c *fiber.Ctx) error {

	klogger.Logf("EmotionsController.Today: %s", c.Path()).Info()

	userID := c.Params("userid")
	start := c.Query("start")
	end := c.Query("end")
	offset := c.Query("offset")
	limit := c.Query("limit")

	params, err := ctrl.processParams(userID, start, end, limit, offset)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	klogger.Logf("processed params").Add("user_id", params.UserID).Info()

	nrcLexResults, err := ctrl.NrcLexService.FindRangeByUserID(
		params.UserID, params.Limit, params.Offset, params.Start, params.End,
	)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(nrcLexResults)
}

type GetUserRecordsPayload struct {
	UserID int64     `json:"userid"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}

// processParams validates and parses the input query parameters: start, end, limit, and offset.
// It ensures that:
// - The start and end parameters are provided and can be parsed into valid dates in RFC3339 format.
// - The end date is chronologically after the start date.
// - The limit and offset parameters are positive integers and do not exceed their respective maximum values set in the controller.
//
// Parameters:
// - start: The start date in RFC3339 format.
// - end: The end date in RFC3339 format.
// - limit: The maximum number of records to return.
// - offset: The number of records to skip before starting to return records.
//
// Returns:
// - An error if any parameter is invalid, including detailed messages for which parameter(s) and why.
// - The parsed start and end dates as *time.Time pointers if valid.
// - The parsed limit and offset as integers if valid.
//
// Note:
// This function is pivotal for ensuring the integrity of data retrieval by applying essential validations and parsing.
func (ctrl *EmotionsController) processParams(userID, start, end, limit, offset string) (*GetUserRecordsPayload, error) {

	var (
		err                            error
		userIDInt, offsetInt, limitInt int
		startDate, endDate             time.Time
	)

	// Convert the offset to an int
	if userIDInt, err = strconv.Atoi(userID); err != nil {

		return nil, fmt.Errorf("invalid user id parameter")
	}

	// Validate we have a start and end date
	if start == "" || end == "" {

		return nil, fmt.Errorf("start and end parameters are required")
	}

	// Validate and parse the start date
	if startDate, err = time.Parse(time.DateTime, start); err != nil {

		return nil, fmt.Errorf("Invalid start date: " + err.Error())
	}

	// Validate and parse the end date
	if endDate, err = time.Parse(time.DateTime, end); err != nil {

		return nil, fmt.Errorf("Invalid end date: " + err.Error())
	}

	// Make sure the end date is after the start date
	if !endDate.After(startDate) {

		return nil, fmt.Errorf("end date must be after start date")
	}

	// Convert the offset to an int
	if offsetInt, err = strconv.Atoi(offset); err != nil {

		return nil, fmt.Errorf("invalid offset parameter")
	}

	// And make sure it's within the range'
	if offsetInt > ctrl.MaxOffset {

		return nil, fmt.Errorf(
			"invalid offset parameter, must be less than %d", ctrl.MaxOffset)
	}

	// Then convert the limit to an int
	if limitInt, err = strconv.Atoi(limit); err != nil {

		return nil, fmt.Errorf("invalid limit parameter")
	}

	// And make sure it's within the range'
	if limitInt > ctrl.MaxLimit {

		return nil, fmt.Errorf(
			"invalid limit parameter, must be less than %d", ctrl.MaxLimit)
	}

	// Lastly, make sure the offset and limit are non-negative
	if offsetInt < 0 || limitInt < 0 {

		return nil, fmt.Errorf("offset and limit parameters must be non-negative")
	}

	return &GetUserRecordsPayload{
		UserID: int64(userIDInt),
		Start:  startDate,
		End:    endDate,
		Limit:  limitInt,
		Offset: offsetInt,
	}, nil
}
