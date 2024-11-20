import React, { useEffect, useState } from "react";
import { ScrollView, StyleSheet, TouchableOpacity, View } from "react-native";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { ThemedSafeAreaView } from "@/components/ThemedSafeAreaView";
import { useRouter, useFocusEffect } from "expo-router";
import Ionicons from "@expo/vector-icons/Ionicons";
import * as SecureStore from "expo-secure-store";
import axios from "axios";

export default function HomeScreen() {
  const router = useRouter();
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const checkToken = async () => {
    const token = await SecureStore.getItemAsync("userToken");
    setIsAuthenticated(!!token);
  };

  useFocusEffect(
    React.useCallback(() => {
      checkToken();
    }, [])
  );

  useEffect(() => {
    checkToken();
  }, []);

  const handleLogout = async () => {
      await SecureStore.deleteItemAsync("userToken");
      setIsAuthenticated(false);
  };

  const steps = [
    {
      number: "1",
      title: isAuthenticated ? "Déconnexion" : "Connexion",
      description: isAuthenticated
        ? "Déconnectez-vous pour quitter l'application."
        : "Connectez-vous pour accéder à toutes les fonctionnalités.",
      icon: "log-in-outline",
      color: "#FF1CF7",
      action: isAuthenticated ? handleLogout : () => router.push("/auth"),
      buttonText: isAuthenticated ? "Déconnexion" : "Commencer",
    },
    {
      number: "2",
      title: "Link un service",
      description:
        "Connecte ton compte à un service tiers comme Google ou Twitch. Cela te permettra de créer des actions-réactions.",
      icon: "link-outline",
      color: "#d433f8",
      action: () => router.push("/services"),
      buttonText: "Connecter",
    },
    {
      number: "3",
      title: "Crée une action-réaction",
      description:
        "Crée une action-réaction pour automatiser une tâche. Par exemple, envoie un mail à ta mère chaque semaine.",
      icon: "git-compare-outline",
      color: "#b249f8",
      action: () => router.push("/area"),
      buttonText: "Créer",
    },
  ];

  return (
    <ThemedSafeAreaView style={styles.safeContainer}>
      <ThemedView style={styles.container}>
        <ThemedView style={styles.header}>
          <ThemedText type="title">Action-REAction</ThemedText>
          <ThemedText style={styles.headerSubText}>
            Automatisez vos tâches simplement
          </ThemedText>
        </ThemedView>

        <ScrollView
          contentContainerStyle={styles.scrollViewContent}
          showsVerticalScrollIndicator={false}
        >
          {steps.map((step, index) => (
            <ThemedView key={index} style={styles.stepContainer}>
              <ThemedView style={styles.stepHeader}>
                <ThemedView style={styles.stepIconContainer}>
                  <ThemedView
                    style={[styles.stepIcon, { backgroundColor: step.color }]}
                  >
                    <Ionicons name={step.icon as any} size={24} color="white" />
                  </ThemedView>
                  <ThemedView
                    style={[styles.stepNumber, { backgroundColor: step.color }]}
                  >
                    <ThemedText>{step.number}</ThemedText>
                  </ThemedView>
                </ThemedView>
                <View style={styles.stepTextContainer}>
                  <ThemedText type="subtitle">{step.title}</ThemedText>
                  <ThemedText>{step.description}</ThemedText>
                </View>
              </ThemedView>

              <TouchableOpacity
                style={[styles.buttonStep, { backgroundColor: step.color }]}
                onPress={step.action}
              >
                <ThemedText style={styles.buttonStepText}>
                  {step.buttonText}
                </ThemedText>
                <Ionicons
                  size={24}
                  name="arrow-forward-outline"
                  color="white"
                />
              </TouchableOpacity>
            </ThemedView>
          ))}
        </ScrollView>
      </ThemedView>
    </ThemedSafeAreaView>
  );
}

const styles = StyleSheet.create({
  safeContainer: {
    flex: 1,
  },
  container: {
    flex: 1,
    paddingHorizontal: 20,
  },
  header: {
    paddingVertical: 20,
    alignItems: "center",
  },
  headerSubText: {
    textAlign: "center",
    marginTop: 8,
    opacity: 0.8,
    fontSize: 16,
  },
  scrollViewContent: {
    paddingVertical: 10,
  },
  stepContainer: {
    padding: 20,
    marginBottom: 20,
    borderRadius: 15,
  },
  stepHeader: {
    flexDirection: "row",
    alignItems: "flex-start",
    marginBottom: 15,
  },
  stepIconContainer: {
    alignItems: "center",
    marginRight: 15,
  },
  stepIcon: {
    width: 45,
    height: 45,
    borderRadius: 23,
    justifyContent: "center",
    alignItems: "center",
    marginBottom: 8,
  },
  stepNumber: {
    width: 24,
    height: 24,
    borderRadius: 12,
    justifyContent: "center",
    alignItems: "center",
  },
  stepTextContainer: {
    flex: 1,
  },
  buttonStep: {
    paddingVertical: 12,
    paddingHorizontal: 20,
    borderRadius: 8,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    marginTop: 15,
  },
  buttonStepText: {
    marginRight: 8,
  },
});
