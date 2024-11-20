# Composants Utilisés

Voici les principaux composants utilisés dans le client web :

## ServiceCard

Composant d'affichage individuel pour chaque service, utilisé dans `ServicesPage`.
- **Props** : `service` (objet contenant les détails du service)
- **Fonctionnalité** : Affiche le nom du service et permet la connexion/déconnexion avec un bouton.

## AreaTable

Tableau pour gérer les AREAs sur `AreaPage`.
- **Props** : Aucun, utilise un état local pour les services.
- **Fonctionnalité** : Affiche chaque service avec des menus de sélection pour les actions et réactions.

## Pagination

Utilisé dans plusieurs pages pour naviguer entre les services et les AREAs.
- **Fonctionnalité** : Permet de naviguer entre les pages.

Consultez chaque page spécifique ([ServicesPage](ServicesPage.md) et [AreaPage](AreaPage.md)) pour voir comment ces composants sont intégrés.
