# Documentation Technique - ServicesScreen

La page **ServicesScreen** permet aux utilisateurs de connecter divers services intégrés, comme Google, GitHub, Discord, Twitter et Dropbox, afin de les utiliser dans des automatisations. Ce composant offre une interface intuitive pour consulter les services, avec des icônes et des boutons d'authentification.

## Structure des Composants

### 1. **ServiceButton**

- **Description** : Composant de bouton stylisé permettant de connecter un service spécifique.
- **Props** :
  - `title` : Titre du bouton (nom du service).
  - `onPress` : Fonction appelée lors du clic sur le bouton.
- **Rendu** :
  - Utilise un bouton tactile `TouchableOpacity`.
  - Affiche le texte du bouton et une icône flèche.
- **Style** : Bouton arrondi avec un fond coloré en `#eb00f7`, utilisé pour indiquer l’action de connexion.

### 2. **ServiceItem**

- **Description** : Carte de présentation d'un service, avec icône, titre, description, et bouton de connexion.
- **Props** :
  - `title` : Nom du service.
  - `icon` : Icône représentant le service (utilise Ionicons).
  - `description` : Brève description du service.
- **Rendu** :
  - Composant en deux parties : l'en-tête avec l'icône, le titre et la description, suivi du bouton de connexion.
- **Style** :
  - Icône circulaire avec fond coloré en `#FF1CF7`.
  - Le texte est stylisé pour une hiérarchie visuelle claire entre le titre et la description.
  - Effet d'ombre pour créer un effet de carte flottante.

### 3. **ServicesScreen**

- **Description** : Composant principal qui regroupe et affiche tous les services disponibles sous forme de cartes.
- **État** :
  - `services` : Liste statique des services, chacun étant un objet contenant le titre, l'icône et la description.
- **Rendu** :
  - Utilise un `ThemedSafeAreaView` pour l'accessibilité et l’adaptation aux thèmes de couleur.
  - En-tête avec titre et sous-titre pour introduire l’objectif de la page.
  - Liste de services affichée via un `ScrollView` pour permettre le défilement.
- **Style** :
  - Le composant est centré avec un en-tête contenant un texte de grande taille.
  - La grille des services est affichée en mode scrollable, avec des marges et un remplissage pour une lisibilité et un espacement optimisés.

## Interface Utilisateur

- **En-tête** : Affiche le titre de la page “Connexion aux services” avec une brève description.
- **Liste des Services** :
  - Chaque service est représenté par une carte contenant une icône, un titre, une description et un bouton de connexion.
  - Les cartes sont défilables pour permettre l’accès à tous les services, même sur les petits écrans.

## Styles et Layout

- Utilise un design `flex` pour assurer la réactivité.
- **Services en Cartes** : Les services sont présentés sous forme de cartes avec ombre portée et coins arrondis.
- **Boutons** : Les boutons de connexion sont colorés de manière vive pour inciter les utilisateurs à cliquer.

## Exemple d’Appels API

Chaque `ServiceButton` pourrait être relié à une fonction d'authentification OAuth, permettant d’appeler l'API backend pour gérer la connexion avec les différents services. Ces appels OAuth devront gérer des redirections et vérifier les permissions des utilisateurs.

## Améliorations Possibles

- **Filtrage et Recherche** : Ajouter un champ de recherche pour filtrer les services par nom.
- **Gestion d'État** : Migrer vers un état global pour gérer les connexions et déconnexions des services.
- **Pagination Dynamique** : Intégrer une pagination ou un chargement dynamique pour une liste de services extensible.

Ce composant `ServicesScreen` constitue une base pour la gestion de connexion des services et peut être étendu avec des fonctionnalités plus avancées, comme le filtrage et la gestion d’état dynamique.
