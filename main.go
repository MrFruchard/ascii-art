package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LireLignesDepuisUneLigne lit les lignes à partir d'une ligne spécifiée et retourne les n lignes suivantes sous forme de slice de chaînes.
func LireLignesDepuisUneLigne(nomFichier string, ligneDebut int, nombreLignes int) ([]string, error) {
	f, err := os.Open(nomFichier)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ligneActuelle := 0
	var lignes []string

	// Lire le fichier ligne par ligne
	for scanner.Scan() {
		ligneActuelle++
		if ligneActuelle >= ligneDebut && len(lignes) < nombreLignes {
			lignes = append(lignes, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lignes, nil
}

// ConstruireAsciiArt génère l'art ASCII pour une chaîne donnée en utilisant le fichier spécifié.
func ConstruireAsciiArt(texte, filename string) (string, error) {
	var result strings.Builder
	const lignesParCaractere = 9

	// Diviser le texte en lignes
	lignesTexte := strings.Split(texte, "\n")

	// Pour chaque ligne de texte
	for indexLigne, ligne := range lignesTexte {
		// Initialiser un tableau pour stocker les lignes de chaque caractère
		caracteres := make([]string, lignesParCaractere)

		for i := 0; i < lignesParCaractere; i++ {
			caracteres[i] = ""
		}

		for _, c := range ligne {

			// Calculer la ligne de début dans le fichier pour le caractère courant
			nbline := (c-32)*lignesParCaractere + 2
			lignes, err := LireLignesDepuisUneLigne(filename, int(nbline), lignesParCaractere)
			if err != nil {
				return "", err
			}

			// Ajouter les lignes du caractère courant au tableau
			for i := 0; i < lignesParCaractere; i++ {
				if i < len(lignes) {
					caracteres[i] += lignes[i]
				}
			}
		}

		// Ajouter les lignes de l'ASCII art pour la ligne courante du texte
		for i := 0; i < lignesParCaractere-1; i++ {
			if len(ligne) > 0 {
				result.WriteString(caracteres[i] + "\n")
			}
		}

		// Ajouter un saut de ligne entre les sections d'ASCII art si ce n'est pas la dernière ligne
		if indexLigne+1 < len(lignesTexte)-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

func main() {
	var filename, texte, output string

	// Vérifier le nombre d'arguments
	switch len(os.Args) {
	case 2:
		filename = "standart.txt"
		texte = os.Args[1]
	case 3:
		texte = os.Args[1]
		filename = os.Args[2] + ".txt"
	case 4:
		output = os.Args[1]
		texte = os.Args[2]
		filename = os.Args[3] + ".txt"
	default:
		fmt.Println("Usage: go run . --output=<fileName.txt> something standard")
		return
	}

	// Convertir les séquences d'échappement en véritables sauts de ligne
	texte = strings.ReplaceAll(texte, "\\n", "\n")

	// Générer l'art ASCII
	asciiArt, err := ConstruireAsciiArt(texte, filename)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	// Vérifier si un fichier de sortie a été spécifié
	if output != "" {
		file, err := os.OpenFile(output[9:], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture du fichier:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(asciiArt)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
			return
		}
	} else {
		// Afficher l'art ASCII à la console
		fmt.Print(asciiArt)
	}
}
