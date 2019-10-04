[L'algorithme de Knuth-Morris-Pratt](https://fr.wikipedia.org/wiki/Algorithme_de_Knuth-Morris-Pratt) est un algorithme de [recherche de sous-chaîne](https://fr.wikipedia.org/wiki/Algorithme_de_recherche_de_sous-cha%C3%AEne).

Il fait un pre-traitement sur la sous-chaîne en creant un tableau avec les sauts a faire pour eviter de la recherche inutile.

L'algorithme ce decoupe en 2 partie, une phase de construction du tableau de saut et une phase de recherche.

## Construction de la table de saut

Pour "participate in parachute" (24), la table est "-1 0 0 0 0 0 0 -1 0 2 0 0 0 0 0 -1 0 0 3 0 0 0 0 0 0 (25)"

La premiere lettre est "p", donc a chaque "p", on inscrit -1 dans la table.

La seconde lettre est "a", donc a chaque "pa", on inscrit 2 dans la table, soit la taille du pattern.

La seconde lettre est "r", donc a chaque "par", on inscrit 3 dans la table, soit la taille du pattern.

...

On met 0 pour chaque combinaison finis par la longueur ou la premiere lettre comme:
* `-1 0 0 0 0 0 0 -1` == `particip`
* `-1 0 0 0 0 0 0 -1 0 2` == `participat` (le `2` signifie les 2 avant le `t`)

Pour `par` on inscrit `-1 0 0`, si le suivant est `t` alors `-1 0 0 0` sinon `-1 0 0 3` (a part le debut bien sûr)

Alors:
* `-1 0 0 0 0 0 0` == `partici`
* `-1 0 2` == `pat`
* `0 0 0 0 0` == `e in `
* `-1 0 0 3` == `para`
* `0 0 0 0 0` == `chute`
* `0` le dernier charactere sert un inscrit une combinaison si y'en a une a la fin, donc le tableau fait la taille de la chaine + 1