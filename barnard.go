package main

import (
	"crypto/tls"

	"github.com/scaredmushroom/barnard/uiterm"
	"github.com/scaredmushroom/gumble/gumble"
	"github.com/scaredmushroom/gumble/gumbleopenal"
)

type Barnard struct {
	Config *gumble.Config
	Client *gumble.Client

	Address   string
	TLSConfig tls.Config

	Stream *gumbleopenal.Stream

	Ui            *uiterm.Ui
	UiOutput      uiterm.Textview
	UiInput       uiterm.Textbox
	UiStatus      uiterm.Label
	UiTree        uiterm.Tree
	UiInputStatus uiterm.Label
}
