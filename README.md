# ASCII Art Generator

Ce programme écrit en Go génère un art ASCII à partir d'une chaîne de caractères, en utilisant des modèles de caractères stockés dans un fichier texte.

## Fonctionnalités

- Lecture d'un fichier contenant les modèles ASCII pour chaque caractère.
- Conversion d'une chaîne de texte en art ASCII.
- Sauvegarde du résultat dans un fichier ou affichage dans la console.
- Gestion des textes avec plusieurs lignes et caractères spéciaux.

## Utilisation

Le programme peut être utilisé via la ligne de commande avec différents paramètres pour définir le texte à convertir, le fichier modèle ASCII et, si besoin, le fichier de sortie.

  - Cas 1 : Générer l'ASCII Art avec un fichier modèle par défaut (standart.txt)
    go run . "votre texte"
  
  - Cas 2 : Utiliser un fichier modèle ASCII spécifique
    go run . "votre texte" "nomFichierModele"

  - Cas 3 : Sauvegarder le résultat dans un fichier de sortie
    go run . --output=sortie.txt "votre texte" "nomFichierModele"

### Compilation

Pour compiler le programme, exécutez la commande suivante dans votre terminal :

```bash
go build
