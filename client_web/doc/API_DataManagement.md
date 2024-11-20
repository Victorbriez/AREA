# API et Gestion de Données

## Appels API

Les pages `ServicesPage` et `AreaPage` utilisent des appels API pour récupérer et gérer les données des services et des AREAs.

- **URL de l'API** : Configurée via `process.env.NEXT_PUBLIC_API_URL`.
- **Authentification** : Les appels incluent un token d'authentification récupéré depuis les cookies.

### Exemples d'Appels API

- **Récupération des Services** :
  - Route : `/v1/providers/`
  - Paramètres : `page`, `pageSize` (pagination)
  - Utilisé dans : `ServicesPage`, `AreaPage`

- **URL OAuth** :
  - Route : `/v1/oauth/{provider}/url?type=link`
  - Utilisé pour la connexion OAuth dans `ServiceCard`.

## Gestion de l'État

Les états principaux sont gérés avec `useState` et `useEffect` :
- **`services`** : Contient la liste des services récupérés de l'API.
- **`areas`** : Contient la liste des AREAs avec leur statut.
- **Pagination** : Utilise `currentPage` et `itemsPerPage` pour gérer la pagination.

Les données sont filtrées et paginées avant d'être affichées à l'utilisateur.
