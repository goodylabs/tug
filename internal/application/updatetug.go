package application

import (
	"fmt"
	"log"
	"os"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/tughelper"
)

// sprawdza czy jest plik z relase.json
//   jak nie to go tworzy
// sprawdzenie czy dzisiaj był check
//   jak tak to return
// patrzy na najnowszą wersję (curl)
// 	 jak wersja taka sama jak nasza to return
// pyta czy chce nową wersję
// 	 jak nie to return
// jak tak to curl + bash na skrypt instalacyjny

func UpdateTugUseCase(tugReleasePath string) {
	tugRelease, err := tughelper.GetTugRelease(tugReleasePath)
	if err != nil {
		err := tughelper.CreateDefaultTugRelease(tugReleasePath)
		if err != nil {
			log.Fatal("placeholder 1", err)
		}
	}

	if tugRelease.LastVersionCheckDate == tughelper.GetToday() {
		return
	}

	rawFileUrl := `https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/scripts/example?token=GHSAT0AAAAAADHX3RHSOY3O4SHQCBRIMSGO2ECDD3A`
	downloadCmd := "curl '" + rawFileUrl + "' | " + os.Getenv("SHELL") + " -s"

	fmt.Println(downloadCmd)
	err = adapters.ShellExecutor.Exec(downloadCmd)
	if err != nil {
		log.Fatal("placeholder 2")
	}

}
