package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	url := "https://get.activated.win"
	outputPath := "C:\\Users\\Public\\activate.ps1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors du téléchargement:", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Erreur de création du fichier:", err)
		return
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde du script:", err)
		out.Close()
		return
	}

	out.Close()

	fmt.Println("Script téléchargé avec succès :", outputPath)

	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", outputPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erreur à l'exécution du script :", err)
	}
	fmt.Println(string(output))
}
