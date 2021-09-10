package abstract_factory

import (
	"github.com/pterm/pterm"
)

type LinuxFactory struct{}

func (f *LinuxFactory) Name() string {
	return "Linux"
}

type LinuxButton struct{}

func (f *LinuxButton) Draw() {
	pterm.Info.Println("Drawing Linux Button")
}

func (f *LinuxFactory) CreateButton() Button {
	return &LinuxButton{}
}

type LinuxListbox struct{}

func (f *LinuxListbox) Draw() {
	pterm.Info.Println("Drawing Linux Listbox")
}

func (f *LinuxFactory) CreateListbox() Listbox {
	return &LinuxListbox{}
}
