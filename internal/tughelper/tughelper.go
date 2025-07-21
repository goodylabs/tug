package tughelper

import (
	"os"
	"runtime"
	"time"

	"log"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/utils"
)

func getDefaultTugRelease() *dto.TugReleaseDTO {
	var platform = runtime.GOOS
	if platform != dto.PlatformLinux && platform != dto.PlatformDarwin {
		log.Fatalln("unsupported platform: " + platform)
	}

	return &dto.TugReleaseDTO{
		Platform:             platform,
		CurrentVersion:       "N/A",
		LastVersionCheckDate: "N/A",
	}
}

func GetToday() string {
	return time.Now().Format("2006-01-02")
}

func GetTugRelease(tugReleaseFilePath string) (*dto.TugReleaseDTO, error) {
	var tugRelease dto.TugReleaseDTO
	err := utils.ReadJSON(tugReleaseFilePath, &tugRelease)
	return &tugRelease, err
}

func CreateDefaultTugRelease(tugReleaseFilePath string) error {
	tugRelease := getDefaultTugRelease()
	return utils.WriteJSON(tugReleaseFilePath, tugRelease)
}

// func CheckNewestFile

// func DownloadNewVersion(tugReleaseFilePath string) error {
// 	adapters.ShellExecutor.RemoteScriptExec("placeholder/v1/alpha")

// 	if tugRelease.LastVersionCheckDate == getToday() {
// 		fmt.Println("Downloading new version...")
// 		return nil
// 	}

//		return nil
//	}

// func UpdateVersionCheckDate(tugReleaseFilePath string) error {
// 	var tugRelease dto.TugReleaseDTO
// 	utils.ReadJSON(tugReleaseFilePath, &tugRelease)

// 	tugRelease.LastVersionCheckDate = GetToday()

// 	utils.WriteJSON(tugReleaseFilePath, &tugRelease)
// 	return nil
// }

func DownloadNewVersion() {
	rawFileUrl := `https://raw.githubusercontent.com/goodylabs/tug/refs/heads/main/scripts/example?token=GHSAT0AAAAAADHX3RHSOY3O4SHQCBRIMSGO2ECDD3A`
	downloadCmd := "curl '" + rawFileUrl + "' | " + os.Getenv("SHELL") + " -s"

	err := adapters.ShellExecutor.Exec(downloadCmd)
	if err != nil {
		log.Fatal("placeholder 2")
	}
}
