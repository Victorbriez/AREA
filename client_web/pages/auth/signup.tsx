import { useState } from "react";
import {
  Card,
  CardBody,
  CardHeader,
  Button,
  Divider,
  Input,
} from "@nextui-org/react";
import { FcGoogle } from "react-icons/fc";
import { FaMicrosoft } from "react-icons/fa";
import axios from "axios";

const SignupPage = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const handleOAuthSignUp = (provider: string) => {
    setIsLoading(true);
    setErrorMessage("");
    const API_URL = process.env.NEXT_PUBLIC_API_URL;

    axios
      .get(`${API_URL}/v1/oauth/${provider}/url?type=register`)
      .then((response) => {
        const { oauth_url } = response.data;

        window.location.href = oauth_url;
      })
      .catch((error) => {
        console.error(
          "Erreur lors de la tentative d'inscription avec OAuth",
          error,
        );
        setErrorMessage(
          `Impossible de se connecter avec ${provider}. Veuillez réessayer.`,
        );
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  const handleSignup = async () => {
    setIsLoading(true);
    setErrorMessage("");
    setTimeout(() => {
      setIsLoading(false);
      setErrorMessage("Email ou mot de passe incorrect. Veuillez réessayer.");
    }, 2000);
  };

  return (
    <div className="flex items-center justify-center min-h-screen">
      <Card className="w-full max-w-md p-8 shadow-xl">
        <CardHeader className="flex flex-col items-center pb-6">
          <h1 className="text-3xl font-bold text-center mb-2">Inscription</h1>
          <p className="text-center text-gray-600">
            Créez un compte pour continuer
          </p>
        </CardHeader>
        <Divider className="flex-1" />
        <CardBody>
          {errorMessage && (
            <div
              className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4"
              role="alert"
            >
              <span className="block sm:inline">{errorMessage}</span>
            </div>
          )}
          <Button
            fullWidth
            className="mb-4 bg-white text-black border border-gray-300 hover:bg-gray-100"
            startContent={<FcGoogle className="text-xl" />}
            onClick={() => handleOAuthSignUp("google")}
          >
            S&apos;inscrire avec Google
          </Button>
          <Button
            fullWidth
            className="mb-4 bg-gray-800 text-white hover:bg-gray-700"
            startContent={<FaMicrosoft className="text-xl" />}
            onClick={() => handleOAuthSignUp("microsoft")}
          >
            S&apos;inscrire avec Microsoft
          </Button>

          <div className="flex items-center my-4">
            <Divider className="flex-1" />
            <span className="px-2 text-gray-500 text-sm">ou</span>
            <Divider className="flex-1" />
          </div>

          <Input
            fullWidth
            className="mb-4"
            placeholder="Nom"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <Input
            fullWidth
            className="mb-4"
            placeholder="Adresse e-mail"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <Input
            fullWidth
            className="mb-4"
            placeholder="Mot de passe"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button
            className="w-full mb-4"
            isLoading={isLoading}
            onClick={handleSignup}
          >
            S&apos;inscrire
          </Button>

          <Divider className="my-4" />
          <p className="text-center text-sm text-gray-600 mt-4">
            Vous avez déjà un compte?{" "}
            <a className="text-blue-600" href="/auth/login">
              Connectez-vous
            </a>
          </p>
        </CardBody>
      </Card>
    </div>
  );
};

export default SignupPage;
