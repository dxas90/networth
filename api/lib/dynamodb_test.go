package nwlib

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/networth-app/networth/api/lib/dotenv"
)

func TestSetToken(t *testing.T) {
	db := NewDynamoDBClient()
	username := "test_set@networth.app"
	institutionID := "ins_1"

	tokensList := []*Token{}
	token1 := &Token{
		AccessToken: "1",
	}
	token2 := &Token{
		AccessToken: "2",
	}
	tokensList = append(tokensList, token1)
	tokensList = append(tokensList, token2)
	tokens := &Tokens{
		Tokens: tokensList,
	}

	// set for all ins
	if err := db.SetToken(username, institutionID, tokens); err != nil {
		t.Errorf("Cannot set token %v", err)
	}
}

func TestGetToken(t *testing.T) {
	db := NewDynamoDBClient()
	username := "test@networth.app"
	institutionID := "ins_1"
	invalidInstitutionID := "ins_1_invalid"

	// get without ins_id
	tokens := db.GetToken(username, "")
	assert.Equal(t, tokens.Tokens[0].InstitutionID, institutionID)

	// get using ins_id
	tokens = db.GetToken(username, institutionID)
	assert.Equal(t, tokens.Tokens[0].InstitutionID, institutionID)

	// get using invalid ins_id
	tokens = db.GetToken(username, invalidInstitutionID)
	assert.Equal(t, len(tokens.Tokens) == 0, true)
}