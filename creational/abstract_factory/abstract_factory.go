package abstract_factory

import (
	"github.com/pterm/pterm"
)

type Button interface {
	Draw()
}

type Listbox interface {
	Draw()
}

type GUIFactory interface {
	Name() string
	CreateButton() Button
	CreateListbox() Listbox
}

type Application struct {
	factory GUIFactory
}

func (a *Application) CreateUI() {
	btn := a.factory.CreateButton()
	btn.Draw()

	lbx := a.factory.CreateListbox()
	lbx.Draw()
}

func NewApp(factory GUIFactory) *Application {
	pterm.Println()
	pterm.Info.Println(factory.Name() + " Factory Started")

	return &Application{
		factory: factory,
	}
}
