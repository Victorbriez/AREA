# Documentation Technique - HomePage

La page **HomePage** est le point d'entrée principal de l'application mobile Action-REAction, offrant une interface simple et intuitive pour se connecter, lier des services, et créer des actions-réactions. Elle assure également la gestion de la session utilisateur.

## Structure des Composants

### Composants Principaux

#### 1. **ThemedSafeAreaView**

- **Description** : Conteneur principal qui ajuste la vue selon les zones sûres de l'écran pour éviter tout chevauchement avec les barres de statut et de navigation.

#### 2. **Header**

- **Description** : En-tête de la page affichant le nom de l'application et une brève description.
- **Style** : Utilise des marges et un alignement central pour le texte, avec un texte en gras pour le titre.

#### 3. **Step Cards**

- **Description** : Ensemble de cartes qui présentent les étapes principales de l'utilisation de l'application.
  - **Étape 1 - Connexion/Déconnexion** : Indique si l'utilisateur est connecté, avec un bouton pour se connecter ou se déconnecter.
  - **Étape 2 - Lier un Service** : Permet de naviguer vers la page des services pour connecter des comptes tiers.
  - **Étape 3 - Créer une Action-Réaction** : Dirige l'utilisateur vers la création d'une nouvelle action-réaction.

### Fonctionnalités

#### **Vérification de l'authentification**

- Utilise **SecureStore** pour stocker de manière sécurisée le token utilisateur.
- **checkToken** : Vérifie l'existence du token de l'utilisateur dans SecureStore et met à jour l'état `isAuthenticated` en conséquence.

#### **Déconnexion**

- **handleLogout** : Supprime le token utilisateur de SecureStore, puis met à jour l'état `isAuthenticated` pour indiquer que l'utilisateur n'est plus connecté.

### Navigation et Gestion d'État

- **useRouter** : Utilisé pour naviguer entre les pages (authentification, services, création d'actions-réactions).
- **useFocusEffect** et **useEffect** : Utilisés pour vérifier l'authentification chaque fois que la page est affichée ou chargée pour s'assurer que l'utilisateur est connecté.

## Interface Utilisateur

### Header

- **Nom de l'application** : "Action-REAction"
- **Sous-titre** : "Automatisez vos tâches simplement"
- Centré en haut de l'écran pour une visibilité maximale.

### Cartes d'Étapes

- **Étape 1 - Connexion/Déconnexion** : Affiche une carte avec un bouton de connexion ou déconnexion selon l'état de l'utilisateur.
- **Étape 2 - Lier un Service** : Carte pour permettre la liaison d'un service tiers (ex : Google, Gmail).
- **Étape 3 - Créer une Action-Réaction** : Permet de créer une nouvelle action-réaction pour automatiser des tâches spécifiques.

### Style et Disposition

- **Couleurs de fond** : Utilise des couleurs vives pour les éléments principaux, avec des ombres pour accentuer les cartes.
- **Boutons** : Les boutons d'action (Connexion, Déconnexion, Connecter, Créer) sont bien espacés et dotés d'icônes pour indiquer l'action.

## Améliorations Potentielles

- **Ajouter un Feedback Visuel** : Indiquer lorsque l'application est en train de vérifier l'authentification.
- **Gestion d'état centralisée** : Utiliser un contexte global pour gérer l'état de l'utilisateur et éviter des appels répétitifs à SecureStore.
- **Accessibilité** : Ajouter des labels et descriptions pour une meilleure navigation via lecteur d'écran.
