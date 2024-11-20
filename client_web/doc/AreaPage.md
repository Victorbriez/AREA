# AreaPage - Gestion des AREAs

La page **AreaPage** permet de gérer les configurations d'actions et réactions (AREAs) pour chaque service intégré.

## Composants principaux

- **AreaTable** : Tableau principal pour l'affichage des services et leurs AREAs.
- **Dropdowns** : Sélecteurs pour les actions et réactions.
- **Statut et Actions** : Indicateur de statut et bouton pour activer/désactiver les services.

## Fonctions

- **handleOAuthLink** : Récupère l'URL OAuth pour connecter un service.
- **Recherche et Filtres** :
  - Filtre les services selon le terme de recherche saisi.
  - Utilise la pagination pour diviser les résultats en plusieurs pages.

## Interface Utilisateur

- **Tableau des AREAs** :
  - Affiche chaque service avec ses actions, réactions et statut.
  - Les colonnes incluent le nom du service, l’action, la réaction, et les actions possibles.
  
- **Pagination** : Permet de naviguer entre les pages de résultats.

Consultez la [section API et Gestion de Données](API_DataManagement.md) pour les détails des appels API et de la gestion de l'état.
