# Micro service exposant une API RESTFUL

Un site web est accessible à cette adresse : http://localhost:8080

Les différentes commandes ont été testées avec Postman et fonctionnent comme suit :
Création : POST "localhost:8080/post" avec de défini dans l'onglet Body puis raw une structure json
 (exemple :
  { 
	"Nom": "foo",
	"Description": "bar"
  }
 )
 Cela permet d'ajouter un nouveau document à la liste en mémoire.

Récupération : GET "localhost:8080/{Nom}" (exemple : localhost:8080/Fab) pour récupérer le document ayant pour nom "Fab"

Suppression : DELETE "localhost:8080/{Nom}" (exemple : localhost:8080/Fab) pour supprimer le document ayant pour nom "Fab"

Dans les documents, pour le moment, le champ ID est incrémenté à chaque nouvel ajout par rapport au nombre de documents présents dans la liste, il n'est pas mis à jour à cas de suppression d'un document.