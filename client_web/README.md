# AREA - Client Web

## Description

AREA (Action-Réaction) est une application permettant de connecter différents services entre eux via des actions et des réactions, similaire à Zapier. Cette partie du projet représente le **Web** où les utilisateurs peuvent :

- Activer ou désactiver des services.
- Visualiser et gérer les actions/réactions de chaque service.
- Rechercher des services dans la liste.

## Prérequis

Avant de commencer, assurez-vous d'avoir les éléments suivants installés sur votre machine :

- **Docker**
- **Docker Compose**
- **Node.js** (si vous souhaitez exécuter l'application en dehors de Docker)
- **npm** (si vous souhaitez exécuter l'application en dehors de Docker)

## Installation

1. Clonez le dépôt GitHub sur votre machine locale :

   ```bash
   git clone https://github.com/EpitechPromo2027/B-DEV-500-LIL-5-1-area-nicolas.pechart
   ```

2. Créez un fichier `.env` à la racine du projet et ajoutez les variables d'environnement suivantes :

   ```bash
   POSTGRES_HOST=...
   POSTGRES_PORT=...
   POSTGRES_USER=...
   POSTGRES_PASSWORD=...
   POSTGRES_DB=...
   GIN_MODE=...
   REDIS_HOST=...
   REDIS_PORT=...
   ```

3. Rendez-vous dans le dossier du projet :

   ```bash
   cd B-DEV-500-LIL-5-1-area-nicolas.pechart/client_web
   ```

4. Créez un fichier `.env` à la racine du dossier client_web et ajoutez les variables d'environnement suivantes :

   ```bash
   NEXT_PUBLIC_API_URL=...
   ```

5. Exécutez la commande suivante pour lancer l'application en mode développement :

   ```bash
   docker-compose up client_web
   ```

6. L'application est désormais accessible à l'adresse suivante : `http://localhost:8081`

### Pages disponibles

- **Home** : Page d'accueil de l'application.
- **Connexion** : Page de connexion à l'application via GitHub, Gmail ou par mail.
- **Services** : Liste des services disponibles (activation/désactivation).
- **AREA** : Liste des actions/réactions de chaque service.

### Technologies utilisées

- React : Bibliothèque JavaScript pour la création d'interfaces utilisateur.
- Next.js : Framework React pour le rendu côté serveur.
- NextUI : Bibliothèque de composants React pour Next.js.
- Tailwind CSS : Framework CSS pour la conception d'interfaces utilisateur.
- NextAuth.js : Bibliothèque d'authentification pour Next.js.
- Docker et Docker Compose : Outils de containerisation pour exécuter les différents services de l'application.
