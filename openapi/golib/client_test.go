package openapi

import (
	"context"
	"testing"
)

func TestChannelList(t *testing.T) {
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("x-api-key", "BACSKuAm6Uy9HzgokZWV_Q2CQpVnWsALw6pqjZL")
	c := NewAPIClient(cfg)
	data, response, err := c.ChannelApi.ChannelList(context.TODO())
	t.Logf("%v", response)
	if err != nil {
		t.Error(err)
		return
	}

	if data.Status.Code != 0 {
		t.Errorf("%v", data.Status)
		return
	}

	if data.Body[0].ChannelName != "Default Channel" {
		t.Error("Failed to find default channel")
	}
}
