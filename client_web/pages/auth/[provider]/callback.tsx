import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import axios from "axios";
import { CircularProgress, Button } from "@nextui-org/react";
import Cookies from "js-cookie";

import { title, subtitle } from "@/components/primitives";

export default function OAuthCallback() {
  const router = useRouter();
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const { code, provider, state } = router.query;
    const API_URL = process.env.NEXT_PUBLIC_API_URL;

    if (code && provider && state) {
      axios
        .get(
          `${API_URL}/v1/oauth/${provider}/callback?state=${state}&code=${code}`,
        )
        .then((response) => {
          if (response.data.error) {
            throw new Error(response.data.error);
          }
          const { token } = response.data;

          Cookies.set("token", token, {
            secure: true,
            sameSite: "strict",
            expires: 1,
          });
          router.push("/");
        })
        .catch((error) => {
          setError(
            error.response?.data?.error ||
              error.message ||
              "Une erreur s'est produite lors de l'authentification. Veuillez réessayer.",
          );
        })
        .finally(() => {
          setIsLoading(false);
        });
    } else if (router.isReady) {
      setError("Paramètres d'authentification manquants.");
      setIsLoading(false);
    }
  }, [router.query, router.isReady, router]);

  const handleRetry = () => {
    router.push("/auth/login");
  };

  return (
    <section className="flex items-center justify-center min-h-screen px-4">
      <div className="flex flex-col items-center justify-center space-y-4 text-center max-w-md w-full">
        <h1
          className={title({
            color: "violet",
          })}
        >
          {isLoading
            ? "Authentification en cours..."
            : error
              ? "Erreur d'authentification"
              : "Authentification réussie"}
        </h1>
        <p className={subtitle({ class: "text-center" })}>
          {isLoading
            ? "Veuillez patienter pendant que nous sécurisons votre connexion."
            : error
              ? error
              : "Vous allez être redirigé vers la page d'accueil..."}
        </p>
        {isLoading && (
          <CircularProgress className="mt-4" color="secondary" size="lg" />
        )}
        {error && (
          <Button
            aria-label="Réessayer la connexion"
            className="mt-4"
            color="secondary"
            radius="full"
            variant="shadow"
            onClick={handleRetry}
          >
            Réessayer la connexion
          </Button>
        )}
      </div>
    </section>
  );
}
