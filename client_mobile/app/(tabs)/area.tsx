import React, { useState } from "react";
import { ScrollView, StyleSheet, TouchableOpacity, Modal } from "react-native";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { ThemedSafeAreaView } from "@/components/ThemedSafeAreaView";
import Ionicons from "@expo/vector-icons/Ionicons";

interface Area {
  id: number;
  action: string;
  reaction: string;
}

interface ActionSelectorProps {
  value: string;
  options: string[];
  onSelect: (value: string) => void;
  placeholder: string;
}

const ActionSelector: React.FC<ActionSelectorProps> = ({
  value,
  options,
  onSelect,
  placeholder,
}) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <ThemedView>
      <TouchableOpacity
        style={styles.dropdownTrigger}
        onPress={() => setIsOpen(!isOpen)}
      >
        <ThemedText>{value || placeholder}</ThemedText>
        <Ionicons
          name={isOpen ? "chevron-up" : "chevron-down"}
          size={20}
          color="#666"
        />
      </TouchableOpacity>

      <Modal
        visible={isOpen}
        transparent={true}
        animationType="fade"
        onRequestClose={() => setIsOpen(false)}
      >
        <TouchableOpacity
          style={styles.modalOverlay}
          activeOpacity={1}
          onPress={() => setIsOpen(false)}
        >
          <ThemedView style={styles.dropdownMenuContainer}>
            <ScrollView>
              {options.map((option) => (
                <TouchableOpacity
                  key={option}
                  style={styles.dropdownItem}
                  onPress={() => {
                    onSelect(option);
                    setIsOpen(false);
                  }}
                >
                  <ThemedText>{option}</ThemedText>
                </TouchableOpacity>
              ))}
            </ScrollView>
          </ThemedView>
        </TouchableOpacity>
      </Modal>
    </ThemedView>
  );
};

const CreateAreaModal: React.FC<{
  visible: boolean;
  onClose: () => void;
  onCreate: (action: string, reaction: string) => void;
  initialAction?: string;
  initialReaction?: string;
}> = ({
  visible,
  onClose,
  onCreate,
  initialAction = "",
  initialReaction = "",
}) => {
  const [action, setAction] = useState(initialAction);
  const [reaction, setReaction] = useState(initialReaction);

  const actionOptions = ["Gmail", "Discord", "Github", "Spotify"];
  const reactionOptions = ["Spotify", "Slack", "Teams", "Discord"];

  const handleCreate = () => {
    if (action && reaction) {
      onCreate(action, reaction);
      setAction("");
      setReaction("");
      onClose();
    }
  };

  return (
    <Modal
      visible={visible}
      transparent={true}
      animationType="slide"
      onRequestClose={onClose}
    >
      <TouchableOpacity
        style={styles.modalOverlay}
        activeOpacity={1}
        onPress={onClose}
      >
        <ThemedView style={styles.modalContent}>
          <ThemedText style={styles.modalTitle}>Créer une AREA</ThemedText>

          <ThemedView style={styles.selectorWrapper}>
            <ThemedText style={styles.label}>Action</ThemedText>
            <ActionSelector
              value={action}
              options={actionOptions}
              onSelect={setAction}
              placeholder="Choisir une action"
            />
          </ThemedView>

          <ThemedView style={styles.selectorWrapper}>
            <ThemedText style={styles.label}>Réaction</ThemedText>
            <ActionSelector
              value={reaction}
              options={reactionOptions}
              onSelect={setReaction}
              placeholder="Choisir une réaction"
            />
          </ThemedView>

          <TouchableOpacity
            style={[
              styles.buttonStep,
              (!action || !reaction) && styles.buttonDisabled,
            ]}
            onPress={handleCreate}
            disabled={!action || !reaction}
          >
            <ThemedText style={styles.buttonStepText}>Créer</ThemedText>
            <Ionicons name="add-circle-outline" size={24} color="white" />
          </TouchableOpacity>
        </ThemedView>
      </TouchableOpacity>
    </Modal>
  );
};

const AreaCreator: React.FC<{ area: Area; onEdit: () => void }> = ({
  area,
  onEdit,
}) => (
  <ThemedView style={styles.areaContainer}>
    <TouchableOpacity onPress={onEdit} style={styles.areaHeader}>
      <ThemedView style={[styles.areaIcon, { backgroundColor: "#FF1CF7" }]}>
        <Ionicons name="git-branch-outline" size={24} color="white" />
      </ThemedView>
      <ThemedView style={styles.areaTextContainer}>
        <ThemedText type="subtitle">
          {area.action} ➜ {area.reaction}
        </ThemedText>
      </ThemedView>
      <Ionicons name="chevron-forward" size={24} color="#666" />
    </TouchableOpacity>
  </ThemedView>
);

const AreaManagerNative: React.FC = () => {
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [areas, setAreas] = useState<Area[]>([
    { id: 1, action: "Gmail", reaction: "Spotify" },
    { id: 2, action: "Discord", reaction: "Slack" },
    { id: 3, action: "Github", reaction: "Teams" },
  ]);
  const [selectedArea, setSelectedArea] = useState<Area | null>(null);

  const handleCreateArea = (action: string, reaction: string) => {
    if (selectedArea) {
      setAreas(
        areas.map((area) =>
          area.id === selectedArea.id ? { ...area, action, reaction } : area
        )
      );
      setSelectedArea(null);
    } else {
      const newArea = {
        id: areas.length + 1,
        action,
        reaction,
      };
      setAreas([...areas, newArea]);
    }
  };

  return (
    <ThemedSafeAreaView style={styles.safeContainer}>
      <ThemedView style={styles.container}>
        <ThemedView style={styles.header}>
          <ThemedText type="title">Gérez vos AREAs</ThemedText>
          <ThemedText>Créez et gérez vos automatisations</ThemedText>
        </ThemedView>

        <TouchableOpacity
          style={styles.addButton}
          onPress={() => {
            setSelectedArea(null);
            setIsModalVisible(true);
          }}
        >
          <ThemedText style={styles.addButtonText}>Nouvelle AREA</ThemedText>
          <Ionicons name="add-circle-outline" size={24} color="#eb00f7" />
        </TouchableOpacity>

        <ScrollView
          contentContainerStyle={styles.scrollViewContent}
          showsVerticalScrollIndicator={false}
        >
          {areas.map((area) => (
            <AreaCreator
              key={area.id}
              area={area}
              onEdit={() => {
                setSelectedArea(area);
                setIsModalVisible(true);
              }}
            />
          ))}
        </ScrollView>

        <CreateAreaModal
          visible={isModalVisible}
          onClose={() => {
            setIsModalVisible(false);
            setSelectedArea(null);
          }}
          onCreate={handleCreateArea}
          initialAction={selectedArea?.action}
          initialReaction={selectedArea?.reaction}
        />
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
  addButton: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    padding: 15,
    marginBottom: 20,
    borderWidth: 2,
    borderColor: "#eb00f7",
    borderRadius: 10,
    borderStyle: "dashed",
  },
  addButtonText: {
    color: "#eb00f7",
    marginRight: 8,
    fontWeight: "bold",
  },
  scrollViewContent: {
    paddingVertical: 10,
  },
  areaContainer: {
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
  areaHeader: {
    flexDirection: "row",
    alignItems: "center",
  },
  areaIcon: {
    width: 45,
    height: 45,
    borderRadius: 23,
    justifyContent: "center",
    alignItems: "center",
    marginRight: 15,
  },
  areaTextContainer: {
    flex: 1,
  },
  modalOverlay: {
    flex: 1,
    backgroundColor: "rgba(0, 0, 0, 0.5)",
    justifyContent: "center",
    alignItems: "center",
  },
  modalContent: {
    width: "90%",
    padding: 20,
    borderRadius: 15,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5,
  },
  modalTitle: {
    fontSize: 20,
    fontWeight: "bold",
    marginBottom: 20,
    textAlign: "center",
  },
  dropdownTrigger: {
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    padding: 12,
    borderWidth: 1,
    borderColor: "#ccc",
    borderRadius: 8,
  },
  dropdownMenuContainer: {
    width: "80%",
    maxHeight: "50%",
    borderRadius: 8,
    padding: 10,
  },
  dropdownItem: {
    padding: 15,
    borderBottomWidth: 1,
    borderBottomColor: "#eee",
  },
  selectorWrapper: {
    marginBottom: 15,
  },
  label: {
    marginBottom: 8,
  },
  buttonStep: {
    paddingVertical: 12,
    paddingHorizontal: 20,
    borderRadius: 8,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    backgroundColor: "#eb00f7",
  },
  buttonStepText: {
    color: "white",
    fontWeight: "bold",
    marginRight: 8,
  },
  buttonDisabled: {
    opacity: 0.5,
  },
});

export default AreaManagerNative;
