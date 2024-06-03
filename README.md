# Fuzzer di Username

Progetto svolto per il PCTO 2023/2024 presso [Ethical Security SRL](https://www.ethsec.com/) da Ali Husnain e Alessio Ghidini

## Descrizione del Progetto

Semplice implementazione in linguaggio Go-lang di un tool CLI, progettato per generare possibili pattern di username a partire da una lista di nomi e cognomi.

### Funzionalità principali

- **Generazione di Username**: A partire da un file contenente nomi e cognomi nel formato "nome.cognome", l'applicazione genera una varietà di pattern di username, ad esempio "nome + prima lettera del cognome", "cognome + prima lettera del nome", "nome-cognome", "cognome.nome", etc.
- **Permutazione di nomi**: A partire da una wordlist di nomi/cognomi comuni italiani (`data/italy.txt`) o globali (`data/world.txt`), l'applicazione genera ogni permutazione tra ogni nome/cognome, applicando poi le varie trasformazioni.
- **Output su File**: Possibilità di specificare un file di output dove salvare tutti gli username generati.

## Requisiti

- Go 1.15 o versione successiva.

## Installazione

1. Clonare la repository:
    ```bash
    git clone https://github.com/Lobsterge/username_fuzzer.git
    ```

2. Compilare l'applicazione:
    ```bash
    cd username_fuzzer
    go build .
    ```

## Utilizzo
```
    ./username_fuzzer <flags>
    
    Usage of ./username_fuzzer:
    -c string
    	-command (shorthand)
    -case
    	Make the usernames case-sensitive, if this flag is not checked they will be all lowercase
    -command string
    	file -> generates usernames from input file
    	italy -> generates usernames from the most common names in Italy
    	world -> generates usernames from the most common names globally
    -cs
    	-case (shorthand)
    -h	-help (shorthand)
    -help
    	Shows the various command line options
    -i string
    	-input (shorthand)
    -input string
    	Path of the input file in the format (name.surname)
    -o string
    	-output (shorthand) (default "output.txt")
    -output string
    	Path of the output file (default "output.txt")
    -p	-permutation (shorthand)
    -permutation
    	Applies a permutation on the list provided by -input

```