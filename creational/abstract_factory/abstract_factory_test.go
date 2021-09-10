package abstract_factory

import (
	"testing"

	"github.com/kr/pretty"
)

func TestFactoryName(t *testing.T) {
	t.Parallel()

	rows := map[string]struct {
		Input    GUIFactory
		Expected string
	}{
		"linux": {
			&LinuxFactory{},
			"Linux",
		},
		"windows": {
			&WinFactory{},
			"Windows",
		},
	}

	for scenario, testData := range rows {
		scenario := scenario
		testData := testData

		t.Run(scenario, func(t *testing.T) {
			t.Parallel()

			if testData.Input.Name() != testData.Expected {
				t.FailNow()
			}
		})
	}
}

func TestFactoryCreateButton(t *testing.T) {
	t.Parallel()

	rows := map[string]struct {
		Input    GUIFactory
		Expected Button
	}{
		"linux": {
			&LinuxFactory{},
			&LinuxButton{},
		},
		"windows": {
			&WinFactory{},
			&WinButton{},
		},
	}

	for scenario, testData := range rows {
		scenario := scenario
		testData := testData

		t.Run(scenario, func(t *testing.T) {
			t.Parallel()

			if diff := pretty.Diff(testData.Expected, testData.Input.CreateButton()); diff != nil {
				t.FailNow()
			}
		})
	}
}

func TestFactoryCreateListbox(t *testing.T) {
	t.Parallel()

	rows := map[string]struct {
		Input    GUIFactory
		Expected Listbox
	}{
		"linux": {
			&LinuxFactory{},
			&LinuxListbox{},
		},
		"windows": {
			&WinFactory{},
			&WinListbox{},
		},
	}

	for scenario, testData := range rows {
		scenario := scenario
		testData := testData

		t.Run(scenario, func(t *testing.T) {
			t.Parallel()

			if diff := pretty.Diff(testData.Expected, testData.Input.CreateListbox()); diff != nil {
				t.FailNow()
			}
		})
	}
}

func TestFactoryButtonDraw(t *testing.T) {
	t.Parallel()

	rows := map[string]struct {
		Input Button
	}{
		"linux": {
			&LinuxButton{},
		},
		"windows": {
			&WinButton{},
		},
	}

	for scenario, testData := range rows {
		scenario := scenario
		testData := testData

		t.Run(scenario, func(t *testing.T) {
			t.Parallel()

			testData.Input.Draw()
		})
	}
}

func TestFactoryListboxDraw(t *testing.T) {
	t.Parallel()

	rows := map[string]struct {
		Input Listbox
	}{
		"linux": {
			&LinuxListbox{},
		},
		"windows": {
			&WinListbox{},
		},
	}

	for scenario, testData := range rows {
		scenario := scenario
		testData := testData

		t.Run(scenario, func(t *testing.T) {
			t.Parallel()

			testData.Input.Draw()
		})
	}
}

func TestNewApp(t *testing.T) {
	t.Parallel()

	rows := map[string]struct {
		Input    GUIFactory
		Expected *Application
	}{
		"linux": {
			&LinuxFactory{},
			&Application{
				factory: &LinuxFactory{},
			},
		},
		"windows": {
			&WinFactory{},
			&Application{
				factory: &WinFactory{},
			},
		},
	}

	for scenario, testData := range rows {
		scenario := scenario
		testData := testData

		t.Run(scenario, func(t *testing.T) {
			t.Parallel()

			if diff := pretty.Diff(testData.Expected, NewApp(testData.Input)); diff != nil {
				t.FailNow()
			}
		})
	}
}

func TestCreateUI(t *testing.T) {
	t.Parallel()

	rows := map[string]struct {
		Input *Application
	}{
		"linux": {
			NewApp(&LinuxFactory{}),
		},
		"windows": {
			NewApp(&WinFactory{}),
		},
	}

	for scenario, testData := range rows {
		scenario := scenario
		testData := testData

		t.Run(scenario, func(t *testing.T) {
			t.Parallel()

			testData.Input.CreateUI()
		})
	}
}
