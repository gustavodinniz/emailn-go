package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, "Campaign X")
	assert.Equal(campaign.Content, "Body")
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNill(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}
