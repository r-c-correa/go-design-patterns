package abstract_factory

import (
	"github.com/pterm/pterm"
)

type WinFactory struct{}

func (f *WinFactory) Name() string {
	return "Windows"
}

type WinButton struct{}

func (f *WinButton) Draw() {
	pterm.Info.Println("Drawing Windows Button")
}

func (f *WinFactory) CreateButton() Button {
	return &WinButton{}
}

type WinListbox struct{}

func (f *WinListbox) Draw() {
	pterm.Info.Println("Drawing Windows Listbox")
}

func (f *WinFactory) CreateListbox() Listbox {
	return &WinListbox{}
}
