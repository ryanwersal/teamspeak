package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	ts "github.com/ryanwersal/teamspeak"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChannelCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,

		Schema: map[string]*schema.Schema{
			"channel_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"channel_topic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"channel_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"permanent": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceChannelRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*ts.APIClient)
	data, _, err := c.ChannelApi.ChannelInfo(context.TODO(), d.Id())
	if err != nil {
		return err
	}

	// Channel does not exist
	if data.Status.Code == 768 {
		d.SetId("")
		return nil
	}

	if data.Status.Code != 0 {
		msg := fmt.Sprintf("Read failed with error (%v): %v", data.Status.Code, data.Status.Message)
		return errors.New(msg)
	}

	body := data.Body[0]
	permanent := body.ChannelFlagPermanent == "1"

	d.Set("channel_name", body.ChannelName)
	d.Set("channel_topic", body.ChannelTopic)
	d.Set("channel_description", body.ChannelDescription)
	d.Set("permanent", permanent)

	return nil
}

func resourceChannelCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("channel_name").(string)
	topic := d.Get("channel_topic").(string)
	description := d.Get("channel_description").(string)
	permanent := d.Get("permanent").(bool)
	permanentStr := "0"
	if permanent {
		permanentStr = "1"
	}

	c := m.(*ts.APIClient)
	data, _, err := c.ChannelApi.ChannelCreate(
		context.TODO(),
		name,
		&ts.ChannelCreateOpts{
			ChannelTopic:         optional.NewString(topic),
			ChannelDescription:   optional.NewString(description),
			ChannelFlagPermanent: optional.NewString(permanentStr),
		},
	)
	if err != nil {
		return err
	}

	if data.Status.Code != 0 {
		msg := fmt.Sprintf("Create failed with error (%v): %v", data.Status.Code, data.Status.Message)
		return errors.New(msg)
	}

	cid := data.Body[0].Cid
	d.SetId(cid)

	return resourceChannelRead(d, m)
}

func resourceChannelUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("channel_name").(string)
	topic := d.Get("channel_topic").(string)
	description := d.Get("channel_description").(string)
	permanent := d.Get("permanent").(bool)
	permanentStr := "0"
	if permanent {
		permanentStr = "1"
	}

	c := m.(*ts.APIClient)
	data, _, err := c.ChannelApi.ChannelEdit(
		context.TODO(),
		d.Id(),
		&ts.ChannelEditOpts{
			ChannelName:          optional.NewString(name),
			ChannelTopic:         optional.NewString(topic),
			ChannelDescription:   optional.NewString(description),
			ChannelFlagPermanent: optional.NewString(permanentStr),
		},
	)
	if err != nil {
		return err
	}

	if data.Status.Code != 0 {
		msg := fmt.Sprintf("Update failed with error (%v): %v", data.Status.Code, data.Status.Message)
		return errors.New(msg)
	}

	return resourceChannelRead(d, m)
}

func resourceChannelDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*ts.APIClient)
	data, _, err := c.ChannelApi.ChannelDelete(
		context.TODO(),
		d.Id(),
		&ts.ChannelDeleteOpts{Force: optional.NewFloat32(1)},
	)
	if err != nil {
		return err
	}

	// Guess it was deleted already
	if data.Status.Code == 768 {
		return nil
	}

	if data.Status.Code != 0 {
		msg := fmt.Sprintf("Delete failed with error (%v): %v", data.Status.Code, data.Status.Message)
		return errors.New(msg)
	}

	return nil
}
