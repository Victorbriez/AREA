import React, { useState } from "react";
import {
  StyleSheet,
  TouchableOpacity,
  TextInput,
  View,
  KeyboardAvoidingView,
  Platform,
} from "react-native";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import Ionicons from "@expo/vector-icons/Ionicons";
import { useRouter } from "expo-router";
import { ThemedSafeAreaView } from "@/components/ThemedSafeAreaView";
import * as SecureStore from "expo-secure-store";
import * as WebBrowser from "expo-web-browser";
import axios from "axios";

export default function AuthScreen() {
  const router = useRouter();
  const [isSignUp, setIsSignUp] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const handleOAuthLogin = async (provider: string) => {
    try {
      const authType = isSignUp ? "register" : "login";
      const API_URL = "http://192.168.1.40:8080";
      const { data } = await axios.get(
        `${API_URL}/v1/oauth/${provider}/url?type=${authType}`
      );
      const { oauth_url, state } = data;

      const result = await WebBrowser.openAuthSessionAsync(oauth_url);

      if (result.type === "success" && result.url) {
        const urlParams = new URL(result.url);
        const code = urlParams.searchParams.get("code");
        const returnedState = urlParams.searchParams.get("state");

        if (state !== returnedState) {
          throw new Error("Invalid state returned from OAuth provider.");
        }

        const response = await axios.get(
          `${API_URL}/v1/oauth/${provider}/callback?code=${code}&state=${state}`
        );
        const { token } = response.data;

        await SecureStore.setItemAsync("userToken", token);

        alert(isSignUp ? "Inscription réussie !" : "Connexion réussie !");
        alert(`Token d'accès pour ${provider}: ${token}`);
        router.back();
      }
    } catch (error) {
      console.error(`Erreur lors de l'authentification ${provider}:`, error);
      alert("Une erreur est survenue lors de l'authentification");
    }
  };

  const handleLoginOrSignUp = async () => {
    const authType = isSignUp ? "register" : "login";
    const API_URL = "http://192.168.1.40:8080";

    try {
      const data = isSignUp
        ? { email, name: email.split("@")[0], password }
        : { email, password };

      const response = await axios.post(
        `${API_URL}/v1/users/${authType}`,
        data
      );
      const { token } = response.data;

      await SecureStore.setItemAsync("userToken", token);

      router.back();
    } catch (error) {
      console.error("Erreur lors de l'authentification :", error);
      alert(
        "Erreur lors de l'authentification. Veuillez vérifier vos informations."
      );
    }
  };

  const handleGoBack = () => {
    router.back();
  };

  return (
    <ThemedSafeAreaView style={styles.safeContainer}>
      <KeyboardAvoidingView
        style={styles.safeContainer}
        behavior={Platform.OS === "ios" ? "padding" : undefined}
      >
        <ThemedView style={styles.container}>
          <TouchableOpacity style={styles.backButton} onPress={handleGoBack}>
            <Ionicons name="arrow-back-circle" size={40} color="#eb00f7" />
          </TouchableOpacity>

          <ThemedView style={styles.header}>
            <ThemedText type="title" style={styles.headerText}>
              {isSignUp ? "Inscription" : "Connexion"}
            </ThemedText>
            <ThemedText style={styles.headerSubText}>
              {isSignUp
                ? "Créez un compte pour commencer"
                : "Accédez à votre compte pour continuer"}
            </ThemedText>
          </ThemedView>

          <ThemedView style={styles.oauthContainer}>
            <TouchableOpacity
              style={[styles.oauthButton, { backgroundColor: "#DB4437" }]}
              onPress={() => handleOAuthLogin("google")}
            >
              <Ionicons size={24} name="logo-google" color="#fff" />
              <ThemedText style={styles.oauthButtonText}>
                {isSignUp
                  ? "S'inscrire avec Google"
                  : "Se connecter avec Google"}
              </ThemedText>
            </TouchableOpacity>

            <TouchableOpacity
              style={[styles.oauthButton, { backgroundColor: "#2F2F2F" }]}
              onPress={() => handleOAuthLogin("microsoft")}
            >
              <Ionicons size={24} name="logo-microsoft" color="#fff" />
              <ThemedText style={styles.oauthButtonText}>
                {isSignUp
                  ? "S'inscrire avec Microsoft"
                  : "Se connecter avec Microsoft"}
              </ThemedText>
            </TouchableOpacity>
          </ThemedView>

          <ThemedText style={styles.separatorText}>ou</ThemedText>

          <ThemedView style={styles.form}>
            <TextInput
              style={styles.input}
              placeholder="Adresse e-mail"
              keyboardType="email-address"
              placeholderTextColor="#666"
              value={email}
              onChangeText={setEmail}
            />
            <TextInput
              style={styles.input}
              placeholder="Mot de passe"
              secureTextEntry={true}
              placeholderTextColor="#666"
              value={password}
              onChangeText={setPassword}
            />

            {isSignUp && (
              <TextInput
                style={styles.input}
                placeholder="Confirmer le mot de passe"
                secureTextEntry={true}
                placeholderTextColor="#666"
                value={confirmPassword}
                onChangeText={setConfirmPassword}
              />
            )}

            <TouchableOpacity
              style={styles.button}
              onPress={handleLoginOrSignUp}
            >
              <ThemedText style={styles.buttonText}>
                {isSignUp ? "S'inscrire" : "Se connecter"}
              </ThemedText>
              <Ionicons
                size={24}
                name={isSignUp ? "person-add-outline" : "log-in-outline"}
                color="#fff"
              />
            </TouchableOpacity>
          </ThemedView>

          <View style={styles.registerContainer}>
            <ThemedText style={styles.registerText}>
              {isSignUp
                ? "Vous avez déjà un compte ? "
                : "Vous n'avez pas de compte ? "}
            </ThemedText>
            <TouchableOpacity onPress={() => setIsSignUp(!isSignUp)}>
              <ThemedText style={styles.registerLink}>
                {isSignUp ? "Connectez-vous" : "Inscrivez-vous"}
              </ThemedText>
            </TouchableOpacity>
          </View>
        </ThemedView>
      </KeyboardAvoidingView>
    </ThemedSafeAreaView>
  );
}

const styles = StyleSheet.create({
  safeContainer: {
    flex: 1,
  },
  container: {
    flex: 1,
    padding: 20,
    justifyContent: "center",
  },
  backButton: {
    position: "absolute",
    top: 40,
    left: 20,
    zIndex: 1,
  },
  header: {
    paddingBottom: 40,
    alignItems: "center",
  },
  headerText: {
    fontSize: 32,
    fontWeight: "bold",
  },
  headerSubText: {
    marginTop: 8,
    opacity: 0.8,
    fontSize: 16,
    textAlign: "center",
  },
  oauthContainer: {
    marginBottom: 20,
  },
  oauthButton: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    paddingVertical: 10,
    paddingHorizontal: 15,
    borderRadius: 8,
    marginBottom: 10,
  },
  oauthButtonText: {
    color: "#fff",
    fontWeight: "bold",
    marginLeft: 8,
  },
  separatorText: {
    textAlign: "center",
    marginBottom: 20,
    fontSize: 16,
    color: "#666",
  },
  form: {
    marginBottom: 40,
  },
  input: {
    backgroundColor: "#f5f5f5",
    borderRadius: 8,
    padding: 15,
    fontSize: 16,
    marginBottom: 20,
    color: "#333",
    borderColor: "#eb00f7",
    borderWidth: 1,
  },
  button: {
    backgroundColor: "#eb00f7",
    paddingVertical: 12,
    paddingHorizontal: 20,
    borderRadius: 8,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
  },
  buttonText: {
    color: "#fff",
    fontWeight: "bold",
    fontSize: 16,
    marginRight: 8,
  },
  registerContainer: {
    flexDirection: "row",
    justifyContent: "center",
  },
  registerText: {
    fontSize: 16,
    color: "#666",
  },
  registerLink: {
    color: "#eb00f7",
    fontWeight: "bold",
  },
});
