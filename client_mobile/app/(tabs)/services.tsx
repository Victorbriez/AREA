import React, { useState, useEffect } from "react";
import Ionicons from "@expo/vector-icons/Ionicons";
import {
  ScrollView,
  StyleSheet,
  TouchableOpacity,
  View,
  TextInput,
  ActivityIndicator,
} from "react-native";
import axios from "axios";
import * as SecureStore from "expo-secure-store";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { ThemedSafeAreaView } from "@/components/ThemedSafeAreaView";

interface Service {
  id: number;
  provider_name: string;
  provider_slug: string;
  status?: "connected" | "disconnected";
}

interface ServiceItemProps {
  service: Service;
  onPress: () => void;
}

const ServiceItem: React.FC<ServiceItemProps> = ({ service, onPress }) => (
  <ThemedView style={styles.serviceContainer}>
    <ThemedView style={styles.serviceHeader}>
      <ThemedView style={[styles.serviceIcon, { backgroundColor: "#FF1CF7" }]}>
        <Ionicons
          name={
            service.provider_slug === "microsoft"
              ? ("logo-microsoft" as any)
              : service.provider_slug === "google"
              ? ("logo-google" as any)
              : service.provider_slug === "github"
              ? ("logo-github" as any)
              : service.provider_slug === "discord"
              ? ("logo-discord" as any)
              : service.provider_slug === "spotify"
              ? ("musical-notes" as any)
              : service.provider_slug === "twitch"
              ? ("logo-twitch" as any)
              : ("link" as any)
          }
          size={24}
          color="white"
        />
      </ThemedView>
      <View style={styles.serviceTextContainer}>
        <ThemedText type="subtitle" style={styles.serviceTitle}>
          {service.provider_name}
        </ThemedText>
      </View>
    </ThemedView>
    <TouchableOpacity
      style={[
        styles.buttonStep,
        {
          backgroundColor: service.status === "connected" ? "#666" : "#eb00f7",
        },
      ]}
      onPress={onPress}
    >
      <ThemedText style={styles.buttonStepText}>
        {service.status === "connected" ? "Déconnecter" : "Connecter"}
      </ThemedText>
      <Ionicons size={24} name="arrow-forward-outline" color="white" />
    </TouchableOpacity>
  </ThemedView>
);

const FilterButton: React.FC<{
  title: string;
  selected: boolean;
  onPress: () => void;
}> = ({ title, selected, onPress }) => (
  <TouchableOpacity
    style={[styles.filterButton, selected && styles.filterButtonSelected]}
    onPress={onPress}
  >
    <ThemedText
      style={[
        styles.filterButtonText,
        selected && styles.filterButtonTextSelected,
      ]}
    >
      {title}
    </ThemedText>
  </TouchableOpacity>
);

const ServicesScreen: React.FC = () => {
  const [services, setServices] = useState<Service[]>([]);
  const [searchTerm, setSearchTerm] = useState("");
  const [selectedCategory, setSelectedCategory] = useState("all");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchServices = async () => {
      const API_URL = "http://192.168.1.40:8080";
      try {
        const token = await SecureStore.getItemAsync("userToken");
        const response = await axios.get(`${API_URL}/v1/providers`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setServices(response.data.data);
      } catch (error) {
        console.error("Erreur lors de la récupération des services", error);
      } finally {
        setLoading(false);
      }
    };

    fetchServices();
  }, []);

  const handleOAuthLink = async (provider: string) => {
    const API_URL = "http://192.168.1.40:8080";
    try {
      const token = await SecureStore.getItemAsync("userToken");
      const response = await axios.get(
        `${API_URL}/v1/oauth/${provider}/url?type=link`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      const { oauth_url } = response.data;
      alert(oauth_url);
    } catch (error) {
      console.error(
        "Erreur lors de la récupération de l'URL de connexion",
        error
      );
    }
  };

  const filteredServices = services.filter(
    (service) =>
      service.provider_name.toLowerCase().includes(searchTerm.toLowerCase()) &&
      (selectedCategory === "all" ||
        (selectedCategory === "connected" && service.status === "connected") ||
        (selectedCategory === "available" && service.status === "disconnected"))
  );

  return (
    <ThemedSafeAreaView style={styles.safeContainer}>
      <ThemedView style={styles.container}>
        <ThemedView style={styles.header}>
          <ThemedText type="title" style={styles.headerText}>
            Gérez vos services
          </ThemedText>
          <ThemedText style={styles.headerSubText}>
            Connectez et gérez vos services intégrés
          </ThemedText>
        </ThemedView>

        <ThemedView style={styles.searchContainer}>
          <Ionicons
            name="search"
            size={20}
            color="#666"
            style={styles.searchIcon}
          />
          <TextInput
            style={styles.searchInput}
            placeholder="Rechercher des services..."
            value={searchTerm}
            onChangeText={setSearchTerm}
            placeholderTextColor="#666"
          />
        </ThemedView>

        {loading ? (
          <ActivityIndicator
            size="large"
            color="#eb00f7"
            style={styles.loader}
          />
        ) : (
          <ScrollView
            contentContainerStyle={styles.scrollViewContent}
            showsVerticalScrollIndicator={false}
          >
            {filteredServices.map((service) => (
              <ServiceItem
                key={service.id}
                service={service}
                onPress={() => handleOAuthLink(service.provider_slug)}
              />
            ))}
            {filteredServices.length === 0 && (
              <ThemedText style={styles.noResults}>
                Aucun service ne correspond à vos critères.
              </ThemedText>
            )}
          </ScrollView>
        )}
      </ThemedView>
    </ThemedSafeAreaView>
  );
};

const styles = StyleSheet.create({
  safeContainer: {
    flex: 1,
  },
  container: {
    flex: 1,
    paddingHorizontal: 20,
  },
  header: {
    paddingTop: 20,
    paddingBottom: 20,
    alignItems: "center",
  },
  headerText: {
    fontSize: 32,
    fontWeight: "bold",
    textAlign: "center",
  },
  headerSubText: {
    textAlign: "center",
    marginTop: 8,
    opacity: 0.8,
    fontSize: 16,
  },
  searchContainer: {
    flexDirection: "row",
    alignItems: "center",
    backgroundColor: "#f5f5f5",
    borderRadius: 10,
    paddingHorizontal: 15,
    marginBottom: 20,
  },
  searchIcon: {
    marginRight: 10,
  },
  searchInput: {
    flex: 1,
    height: 45,
    fontSize: 16,
  },
  filterContainer: {
    flexDirection: "row",
    justifyContent: "space-around",
    marginBottom: 20,
  },
  filterButton: {
    paddingVertical: 8,
    paddingHorizontal: 16,
    borderRadius: 20,
    backgroundColor: "#f5f5f5",
  },
  filterButtonSelected: {
    backgroundColor: "#eb00f7",
  },
  filterButtonText: {
    fontSize: 14,
    color: "#666",
  },
  filterButtonTextSelected: {
    color: "white",
  },
  scrollViewContent: {
    paddingVertical: 10,
  },
  serviceContainer: {
    padding: 20,
    marginBottom: 20,
    borderRadius: 15,
    shadowColor: "#eb00f7",
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.2,
    shadowRadius: 5,
    elevation: 5,
  },
  serviceHeader: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: 15,
  },
  serviceIcon: {
    width: 45,
    height: 45,
    borderRadius: 23,
    justifyContent: "center",
    alignItems: "center",
    marginRight: 15,
  },
  serviceTextContainer: {
    flex: 1,
  },
  serviceTitle: {
    fontSize: 22,
    fontWeight: "600",
  },
  buttonStep: {
    paddingVertical: 12,
    paddingHorizontal: 20,
    borderRadius: 8,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    marginTop: 10,
  },
  buttonStepText: {
    color: "white",
    fontWeight: "bold",
    marginRight: 8,
  },
  loader: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  noResults: {
    textAlign: "center",
    fontSize: 16,
    color: "#666",
    marginTop: 20,
  },
});

export default ServicesScreen;
