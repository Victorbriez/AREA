# ServicesPage - Gestion des Services

La page **ServicesPage** permet aux utilisateurs de consulter, rechercher, filtrer, connecter et déconnecter différents services intégrés.

## Structure des Composants

- **ServiceCard** : Composant d'affichage pour chaque service.
  - Affiche le nom du service et un bouton pour connecter/déconnecter le service.
  - Gère l'authentification OAuth via un appel à l'API.
  
- **ServicesPage** : Composant principal de la page.
  - **useState** : 
    - `services` : État contenant la liste des services.
    - `searchTerm` : Terme de recherche.
    - `selectedCategory` : Catégorie sélectionnée (tous, connectés, disponibles).
    - `currentPage` : Page actuelle pour la pagination.
  - **useEffect** : Récupère la liste des services depuis l'API.

## Interface Utilisateur

- **Moteur de recherche** : Permet de filtrer les services par nom.
- **Onglets de catégories** : Filtre entre tous les services, les services connectés et ceux disponibles.
- **Pagination** : Navigue entre les pages de services.
  
## Styles et Layout
- Utilise `DefaultLayout` pour la structure générale.
- **Cards** : Les services sont présentés sous forme de cartes interactives.
- **Responsive** : Dispose les éléments en grille selon la taille de l'écran.

Consultez la [section API et Gestion de Données](API_DataManagement.md) pour plus de détails sur les appels API associés.
