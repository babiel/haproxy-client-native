package configuration

import (
	"strconv"
	"strings"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/haproxytech/models"
)

// GetTCPConnectionRules returns a struct with configuration version and an array of
// configured tcp connection rules in the specified frontend. Returns error on fail.
func (c *LBCTLConfigurationClient) GetTCPConnectionRules(frontend string, transactionID string) (*models.GetTCPConnectionRulesOKBody, error) {
	tcpRulesStr, err := c.executeLBCTL("l7-service-tcpreqconn-dump", transactionID, frontend)
	if err != nil {
		return nil, err
	}

	tcpRules := c.parseTCPConnectionRules(tcpRulesStr)

	v, err := c.GetVersion()
	if err != nil {
		return nil, err
	}

	return &models.GetTCPConnectionRulesOKBody{Version: v, Data: tcpRules}, nil
}

// GetTCPConnectionRule returns a struct with configuration version and a requested tcp connection rule
// in the specified frontend. Returns error on fail or if tcp connection rule does not exist.
func (c *LBCTLConfigurationClient) GetTCPConnectionRule(id int64, frontend string, transactionID string) (*models.GetTCPConnectionRuleOKBody, error) {
	tcpRuleStr, err := c.executeLBCTL("l7-service-tcpreqconn-show", transactionID, frontend, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	tcpRule := &models.TCPRule{ID: id}

	c.parseObject(tcpRuleStr, tcpRule)

	v, err := c.GetVersion()
	if err != nil {
		return nil, err
	}

	return &models.GetTCPConnectionRuleOKBody{Version: v, Data: tcpRule}, nil
}

// DeleteTCPConnectionRule deletes a tcp connection rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *LBCTLConfigurationClient) DeleteTCPConnectionRule(id int64, frontend string, transactionID string, version int64) error {
	return c.deleteObject(strconv.FormatInt(id, 10), "tcpreqconn", frontend, "service", transactionID, version)
}

// CreateTCPConnectionRule creates a tcp connection rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *LBCTLConfigurationClient) CreateTCPConnectionRule(frontend string, data *models.TCPRule, transactionID string, version int64) error {
	validationErr := data.Validate(strfmt.Default)
	if validationErr != nil {
		return NewConfError(ErrValidationError, validationErr.Error())
	}
	return c.createObject(strconv.FormatInt(data.ID, 10), "tcpreqconn", frontend, "service", data, nil, transactionID, version)
}

// EditTCPConnectionRule edits a tcp connection rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *LBCTLConfigurationClient) EditTCPConnectionRule(id int64, frontend string, data *models.TCPRule, transactionID string, version int64) error {
	validationErr := data.Validate(strfmt.Default)
	if validationErr != nil {
		return NewConfError(ErrValidationError, validationErr.Error())
	}
	ondiskBr, err := c.GetTCPConnectionRule(id, frontend, transactionID)
	if err != nil {
		return err
	}

	return c.editObject(strconv.FormatInt(data.ID, 10), "tcpreqconn", frontend, "service", data, ondiskBr, nil, transactionID, version)
}

func (c *LBCTLConfigurationClient) parseTCPConnectionRules(response string) models.TCPRules {
	tcpRules := make(models.TCPRules, 0, 1)
	for _, tcpRulesStr := range strings.Split(response, "\n\n") {
		if strings.TrimSpace(tcpRulesStr) == "" {
			continue
		}
		idStr, _ := splitHeaderLine(tcpRulesStr)
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			id = 0
		}

		tcpRulesObj := &models.TCPRule{ID: id}
		c.parseObject(tcpRulesStr, tcpRulesObj)
		tcpRules = append(tcpRules, tcpRulesObj)
	}
	return tcpRules
}