package main

import (
	"github.com/databr/api/database"
	"github.com/databr/bots/go_bot/parser"
	"github.com/databr/parliamentarian-bot/bot"
)

func main() {
	mongo := database.NewMongoDB()

	parser.Log.Debug("Running bot.SaveDeputiesFromTransparenciaBrasil{}")
	bot.SaveDeputiesFromTransparenciaBrasil{}.Run(mongo)

	parser.Log.Debug("Running bot.SaveDeputiesFromSearch{}")
	bot.SaveDeputiesFromSearch{}.Run(mongo)

	parser.Log.Debug("Running bot.SaveDeputiesFromXML{}")
	bot.SaveDeputiesFromXML{}.Run(mongo)

	parser.Log.Debug("Running bot.SaveDeputiesAbout{}")
	bot.SaveDeputiesAbout{}.Run(mongo)

	parser.Log.Debug("Running bot.SavePartiesFromTSE{}")
	bot.SavePartiesFromTSE{}.Run(mongo)

	//	bot.SaveDeputiesQuotas{}.Run(mongo)

	parser.Log.Debug("Running bot.SaveSenatorsFromIndex{}")
	bot.SaveSenatorsFromIndex{}.Run(mongo)
}
