package main

import (
	"github.com/fatih/color"

	// on peut utiliser n'importe quoi comme import, on s'en fout
	// ==> voir le go.mod qui replace vers le bon répertoire
	// et c'est le nom de package qui compte (ou forcer le nom d'import)
	"nimportequoi/ole"

	// juste un package mais pas un module, pas besoin de trifouiller le path
	// on force un nom d'import parce que les 2 ont le meme package name turlututu
	titi "youpie/titi"
)

func main() {
	color.HiMagenta("Hello, world! dans le main\n")

	// le nom de module nest pas important
	// ce qui compte c'est le nom du package (ou d'import)
	turlututu.Test()

	// celui du package
	titi.Test()

	color.Cyan("\nHello, world! in main")
}
