# Documentation Technique - AreaPage

La page **AreaPage** pour mobile est conçue pour permettre aux utilisateurs de gérer les configurations d'actions et de réactions (AREAs) directement depuis leur appareil. Cette interface mobile réinvente l’expérience de gestion des AREAs en adaptant les composants pour les rendre accessibles et intuitifs pour les petits écrans.

## Composants Principaux

### 1. **AreaList**

- **Description** : Liste principale affichant chaque service et ses AREAs configurés.
- **Fonctionnalités** :
  - Affiche les informations essentielles de chaque service (nom, statut, action, réaction).
  - Permet aux utilisateurs d'activer/désactiver les AREAs pour chaque service.
  - Utilise un design en carte pour chaque entrée pour une meilleure lisibilité.

### 2. **ActionReactionSelector**

- **Description** : Composant de sélection permettant de choisir les actions et réactions pour chaque service.
- **Fonctionnalités** :
  - Intègre des menus déroulants ou des éléments de sélection pour afficher les actions et réactions disponibles pour chaque service.
  - Permet d'enregistrer les configurations de l'utilisateur.

### 3. **StatusToggle**

- **Description** : Indicateur d’état et bouton d'activation/désactivation.
- **Fonctionnalités** :
  - Permet de visualiser le statut actif/inactif de chaque AREA.
  - Active ou désactive une AREA d’un simple tapotement.

### 4. **PaginationButtons**

- **Description** : Ensemble de boutons pour naviguer entre les pages de résultats lorsque la liste est longue.
- **Fonctionnalités** :
  - Facilité de navigation pour passer d'une page de services à une autre.
  - Présente les options de "Précédent" et "Suivant" avec des icônes pour une navigation intuitive.

## Fonctions

### **handleOAuthLink**

- **Description** : Récupère l'URL OAuth pour connecter un service à partir de l'application mobile.
- **Fonctionnalités** :
  - Envoie une requête au backend pour obtenir l’URL OAuth.
  - Utilise une méthode de redirection vers l’authentification dans le navigateur intégré de l'application pour gérer le processus de connexion.

### **Recherche et Filtres**

- **Filtrage** : Filtre les services disponibles en fonction du terme de recherche entré par l'utilisateur.
- **Pagination** : Divise les services en pages pour faciliter la gestion et la navigation lorsque le nombre de services est important.

## Interface Utilisateur

### **Liste des AREAs**

- **Présentation en Cartes** : Affiche chaque service dans un format de carte pour une navigation facile. Chaque carte comprend :
  - **Nom du Service** : Affiché avec une icône représentative du service.
  - **Action** : Action sélectionnée pour le service, avec un indicateur montrant si une action est active.
  - **Réaction** : Réaction choisie, indiquée sous forme de texte.
  - **Statut** : Indicateur d'état (Actif/Inactif) facilement visible.
  - **Actions Disponibles** : Boutons pour activer/désactiver l’AREA, changer l’action ou la réaction.

### **Pagination**

- Affichée en bas de la liste, avec des boutons clairs pour permettre à l'utilisateur de parcourir les pages de services.

### **Accessibilité**

- Les éléments sont adaptés aux tailles d’écran mobile et sont placés pour une utilisation à une main.
- Les icônes et les boutons sont suffisamment espacés pour éviter les erreurs de sélection.

## Améliorations Potentielles

- **Recherche Dynamique** : Ajouter un champ de recherche pour filtrer les services par nom, actualisé en temps réel.
- **Gestion d'État Global** : Intégrer un état centralisé pour gérer les connexions et statuts des AREAs pour chaque service.
- **Mise en Cache** : Enregistrer localement la configuration des AREAs pour une récupération rapide même hors ligne.

## Section API et Gestion des Données

La page mobile `AreaPage` requiert des appels API pour récupérer, mettre à jour et gérer les configurations des AREAs. Consultez la [section API et Gestion de Données](API_DataManagement.md) pour des informations détaillées sur les schémas de données et les endpoints utilisés.
