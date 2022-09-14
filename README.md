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