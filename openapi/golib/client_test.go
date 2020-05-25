package openapi

import (
	"context"
	"os"
	"testing"

	"github.com/antihax/optional"
)

func BuildConfig() *Configuration {
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("x-api-key", os.Getenv("TEAMSPEAK_TOKEN"))
	return cfg
}

func BuildClient() *APIClient {
	c := NewAPIClient(BuildConfig())
	return c
}

func TestChannelList(t *testing.T) {
	c := BuildClient()
	data, _, err := c.ChannelApi.ChannelList(context.TODO())
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

func TestChannelCreate(t *testing.T) {
	c := BuildClient()
	var cid string

	{
		data, _, err := c.ChannelApi.ChannelCreate(context.TODO(), "Test Channel Name", &ChannelCreateOpts{})
		if err != nil {
			t.Error(err)
			return
		}

		if data.Status.Code != 0 {
			t.Errorf("%v", data.Status)
			return
		}

		cid = data.Body[0].Cid
	}

	{
		data, _, err := c.ChannelApi.ChannelDelete(context.TODO(), cid, &ChannelDeleteOpts{Force: optional.NewFloat32(1)})
		if err != nil {
			t.Error(err)
			return
		}

		if data.Status.Code != 0 {
			t.Errorf("%v", data.Status)
			return
		}
	}
}

func TestChannelEdit(t *testing.T) {
	c := BuildClient()
	var cid string

	{
		data, _, err := c.ChannelApi.ChannelCreate(context.TODO(), "Test Channel", &ChannelCreateOpts{})
		if err != nil {
			t.Error(err)
			return
		}

		if data.Status.Code != 0 {
			t.Errorf("%v", data.Status)
		}

		cid = data.Body[0].Cid
	}

	{
		data, _, err := c.ChannelApi.ChannelEdit(context.TODO(), cid, &ChannelEditOpts{ChannelDescription: optional.NewString("Test")})
		if err != nil {
			t.Error(err)
			return
		}

		if data.Status.Code != 0 {
			t.Errorf("%v", data.Status)
			return
		}
	}

	{
		data, _, err := c.ChannelApi.ChannelDelete(context.TODO(), cid, &ChannelDeleteOpts{Force: optional.NewFloat32(1)})
		if err != nil {
			t.Error(err)
			return
		}

		if data.Status.Code != 0 {
			t.Errorf("%v", data.Status)
			return
		}
	}
}

func TestChannelDeleete(t *testing.T) {
	c := BuildClient()
	var cid string

	{
		data, _, err := c.ChannelApi.ChannelCreate(context.TODO(), "Test", &ChannelCreateOpts{})
		if err != nil {
			t.Error(err)
			return
		}

		if data.Status.Code != 0 {
			t.Errorf("%v", data.Status)
			return
		}

		cid = data.Body[0].Cid
	}

	{
		data, _, err := c.ChannelApi.ChannelDelete(context.TODO(), cid, &ChannelDeleteOpts{Force: optional.NewFloat32(1.)})
		if err != nil {
			t.Error(err)
			return
		}

		if data.Status.Code != 0 {
			t.Errorf("%v", data.Status)
			return
		}
	}
}
