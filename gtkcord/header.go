package gtkcord

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/pkg/errors"
)

type Header struct {
	gtk.IWidget
	Bar  *gtk.HeaderBar
	Main *gtk.Box

	// Grid 1, on top of guilds
	Hamburger *HeaderMenu

	// Grid 2, on top of channels
	GuildName *gtk.Label
	// GuildButton TODO

	// Grid 3, on top of messages
	ChannelName *gtk.Label
	// Separator ---
	ChannelTopic *gtk.Label
}

func newHeader() (*Header, error) {
	h, err := gtk.HeaderBarNew()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create headerbar")
	}
	h.SetShowCloseButton(true)

	g, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grid")
	}

	h.PackStart(g)

	/*
	 * Grid 1
	 */

	hamburger, err := newHeaderMenu()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create hamburger")
	}
	hamseparator, err := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create ham separator")
	}
	g.PackStart(hamburger, false, false, 0)
	g.PackStart(hamseparator, false, false, 0)

	/*
	 * Grid 2
	 */

	label, err := gtk.LabelNew("test test")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create guild name label")
	}
	label.SetXAlign(0.0)
	label.SetMarginStart(20)
	label.SetSizeRequest(ChannelsWidth, -1)
	lblseparator, err := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create ham separator")
	}
	g.PackStart(label, false, false, 0)
	g.PackStart(lblseparator, false, false, 0)

	/*
	 * Grid 3
	 */

	chname, err := gtk.LabelNew("#idk")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create channel name")
	}
	chname.SetXAlign(0.0)
	chname.SetMarginStart(20)
	chtopic, err := gtk.LabelNew("topic")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create channel topic")
	}
	chtopic.SetXAlign(0.0)
	chtopic.SetMarginStart(20)

	g.PackStart(chname, false, false, 0)
	g.PackStart(chtopic, false, false, 0)

	return &Header{
		IWidget:      h,
		Bar:          h,
		Main:         g,
		Hamburger:    hamburger,
		GuildName:    label,
		ChannelName:  chname,
		ChannelTopic: chtopic,
	}, nil
}

// func (h *Header) hookGuild(a *Application) {}