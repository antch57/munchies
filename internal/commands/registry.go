package commands

import "github.com/antch57/munchies/models"

// Commands is the registry of all available commands.
var Registry = []models.Command{
	{Name: "help", Help: "Print this help", Run: PrintHelpCmd},
	{Name: "add", Help: "Record a snack that was eaten", Run: AddSnackCmd},
	{Name: "list", Help: "List snacks that were eaten", Run: ListSnackCmd},
}
