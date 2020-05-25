package teamspeak

import (
	"context"
	"os"
	"testing"

	"github.com/antihax/optional"
)

func BuildConfig() *Configuration {
	cfg := NewConfiguration()
	cfg.BasePath = os.Getenv("TEAMSPEAK_API_URL")
	cfg.AddDefaultHeader("x-api-key", os.Getenv("TEAMSPEAK_TOKEN"))
	return cfg
}

func BuildClient() *APIClient {
	c := NewAPIClient(BuildConfig())
	return c
}

func TestChannelInfo(t *testing.T) {
	c := BuildClient()
	data, _, err := c.ChannelApi.ChannelInfo(context.TODO(), "1")
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
		data, _, err := c.ChannelApi.ChannelCreate(context.TODO(), "Test Channel Name", &ChannelCreateOpts{ChannelFlagPermanent: optional.NewString("1")})
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

func TestChannelDelete(t *testing.T) {
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
