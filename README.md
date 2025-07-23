# Interprete Monkey in Go

Questo repository contiene un'implementazione completa del linguaggio di programmazione "Monkey", scritta in Go, basata sul fantastico libro di Thorsten Ball, **"Writing An Interpreter In Go"**.

L'obiettivo di questo progetto è esplorare e comprendere i meccanismi interni di un interprete, costruendo ogni componente da zero, senza l'ausilio di librerie o tool di terze parti.

## Il Linguaggio Monkey

Monkey è un linguaggio didattico con una sintassi C-like, progettato per essere semplice ma abbastanza potente da includere funzionalità avanzate. Tra le sue caratteristiche principali troviamo:
- Sintassi simile al C
- Variabili (bindings) con `let`
- Tipi di dati: Interi, Booleani, Stringhe, Array e Hash
- Espressioni aritmetiche e logiche
- Funzioni di prima classe e di ordine superiore (Higher-Order Functions)
- Chiusure (Closures)
- Un sistema di funzioni predefinite (built-in)

## Come Funziona: L'Architettura dell'Interprete

L'interprete è costruito seguendo un'architettura classica a "tree-walking", suddivisa in tre fasi principali:

```
+----------------+      +----------+      +---------+      +-----------+      +----------+
| Codice Sorgente|----->|  Lexer   |----->|  Tokens |----->|  Parser   |----->|    AST   |
+----------------+      +----------+      +---------+      +-----------+      +----------+
                                                                                  |
                                                                                  |
                                                                                  v
                                                                            +----------+
                                                                            | Risultato|
                                                                            +----------+
                                                                                  ^
                                                                                  |
                                                                                  |
                                                                            +-----------+
                                                                            | Valutatore|
                                                                            +-----------+
```

### 1. Lexer (Analisi Lessicale)
Il **Lexer** è il primo componente che analizza il codice. Legge il testo sorgente carattere per carattere e lo trasforma in una sequenza di "token". Un token è un'unità logica minima del linguaggio, come una parola chiave (`let`, `fn`), un identificatore (`x`), un numero (`123`) o un operatore (`+`).

`let five = 5;`  =>  `[LET, IDENT("five"), ASSIGN, INT("5"), SEMICOLON]`

### 2. Parser (Analisi Sintattica)
Il **Parser** riceve la sequenza di token dal Lexer e verifica che la loro disposizione segua le regole grammaticali del linguaggio. Se la sintassi è corretta, costruisce una struttura dati ad albero chiamata **Abstract Syntax Tree (AST)**. L'AST rappresenta la struttura gerarchica e logica del programma, ignorando dettagli come spazi bianchi o punti e virgola.

Per gestire espressioni complesse e la precedenza degli operatori (es. `*` prima di `+`), il parser utilizza l'elegante algoritmo **Pratt Parsing**.

![AST](ast.png)





### 3. Valutatore (Evaluation)
Il **Valutatore** è il cuore dell'interprete. "Cammina" sull'AST (tree-walking) nodo per nodo e dà un significato (semantica) al programma. Utilizza una funzione ricorsiva, `Eval`, per eseguire le azioni corrispondenti a ogni nodo:
-   **Calcoli**: Esegue operazioni aritmetiche e logiche.
-   **Variabili**: Salva e recupera i valori delle variabili usando una struttura chiamata **Environment**, che agisce come una "memoria" per gli scope.
-   **Controllo di Flusso**: Gestisce le condizioni `if/else` e le istruzioni `return`.
-   **Funzioni**: Crea oggetti funzione, gestisce le chiamate e, grazie all'Environment, supporta le chiusure.

Il risultato finale della valutazione è un "oggetto" interno che rappresenta il valore calcolato.

## Struttura del Repository

Il codice è organizzato in pacchetti, ognuno con una responsabilità specifica:
-   `main.go`: Il punto di ingresso del programma che avvia il REPL.
-   `/ast`: Contiene le definizioni delle strutture dati per i nodi dell'Abstract Syntax Tree.
-   `/lexer`: Il tokenizzatore che trasforma il codice sorgente in token.
-   `/parser`: L'analizzatore sintattico che costruisce l'AST a partire dai token.
-   `/evaluator`: Il valutatore che esegue il codice camminando sull'AST.
-   `/object`: Definisce il sistema di oggetti interni per rappresentare i valori (interi, booleani, funzioni, ecc.) durante la valutazione.
-   `/token`: Definisce i tipi di token usati dal Lexer e dal Parser.
-   `/repl`: Implementa il ciclo Read-Eval-Print Loop, l'interfaccia interattiva da riga di comando.

## Come Eseguirlo

Per avviare l'interprete in modalità interattiva (REPL), è sufficiente avere Go installato ed eseguire:
```sh
go run main.go
```
Apparirà un prompt `>>` dove potrai scrivere codice Monkey.

## Esempi di Codice Monkey

Ecco alcuni esempi di ciò che il linguaggio Monkey può fare:

#### Variabili e Operazioni
```monkey
>> let x = 10;
>> let y = x * 2;
>> y / 4;
5
```

#### Funzioni e Chiusure (Closures)
```monkey
>> let newAdder = fn(x) { fn(y) { x + y }; };
>> let addTwo = newAdder(2);
>> addTwo(3);
5
```

#### Funzioni di Ordine Superiore
```monkey
>> let twice = fn(f, x) { f(f(x)); };
>> let addThree = fn(x) { x + 3; };
>> twice(addThree, 2);
8
```


## Riconoscimenti

Questo progetto è un'implementazione diretta del lavoro descritto nel libro **"Writing An Interpreter In Go"** di **Thorsten Ball**. Tutti i concetti e il design sono merito suo.
