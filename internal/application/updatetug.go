package application

import (
	"fmt"
	"log"

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
			log.Fatal("placeholder")
		}
	}

	if tugRelease.LastVersionCheckDate == tughelper.GetToday() {
		return
	}

	rawFileUrl := `https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/scripts/download.sh?token=GHSAT0AAAAAADHX3RHT3VVLZGAJPGUZT4L22ECC64Q`
	downloadCmd := fmt.Sprintf("curl %s | `which $(echo $SHELL)` -s ", rawFileUrl)
	err = adapters.ShellExecutor.Exec(downloadCmd)
	if err != nil {
		log.Fatal("placeholder")
	}

}
