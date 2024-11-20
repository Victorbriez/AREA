import React, { useState, useEffect } from "react";
import {
  Button,
  Card,
  CardHeader,
  Tabs,
  Tab,
  Input,
  Pagination,
} from "@nextui-org/react";
import { Search } from "lucide-react";
import axios from "axios";
import Cookies from "js-cookie";

import DefaultLayout from "@/layouts/default";
import { title } from "@/components/primitives";

interface Service {
  id: number;
  provider_name: string;
  provider_slug: string;
  status?: "connected" | "disconnected";
}

const ServiceCard: React.FC<{ service: Service }> = ({ service }) => {
  const handleOAuthLink = async (provider: string) => {
    const API_URL = process.env.NEXT_PUBLIC_API_URL;
    const token = Cookies.get("token");

    if (!token) {
      console.error("Aucun token d'authentification trouvé.");

      return;
    }

    try {
      const response = await axios.get(
        `${API_URL}/v1/oauth/${provider}/url?type=link`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        },
      );

      const { oauth_url } = response.data;

      window.location.href = oauth_url;
    } catch (error) {
      console.error(
        "Erreur lors de la récupération de l'URL de connexion",
        error,
      );
    }
  };

  return (
    <Card className="w-full transform transition-all duration-300 hover:scale-105 hover:shadow-lg hover:border-violet-500 border-transparent border-2">
      <CardHeader className="flex flex-col sm:flex-row justify-between items-center gap-4 p-4">
        <div className="flex items-center gap-3">
          <div className="text-center sm:text-left">
            <p className="text-md font-semibold text-indigo-600">
              {service.provider_name}
            </p>
          </div>
        </div>
        <Button
          aria-label="Connecter ou déconnecter le service"
          className="w-full sm:w-auto"
          color="secondary"
          size="sm"
          variant={service.status === "connected" ? "faded" : "solid"}
          onClick={() => handleOAuthLink(service.provider_slug)}
        >
          {service.status === "connected" ? "Déconnecter" : "Connecter"}
        </Button>
      </CardHeader>
    </Card>
  );
};

const ServicesPage: React.FC = () => {
  const [services, setServices] = useState<Service[]>([]);
  const [searchTerm, setSearchTerm] = useState("");
  const [selectedCategory, setSelectedCategory] = useState("all");
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 6;

  useEffect(() => {
    const fetchServices = async () => {
      const API_URL = process.env.NEXT_PUBLIC_API_URL;

      try {
        const response = await axios.get(
          `${API_URL}/v1/providers/?page=${currentPage}&pageSize=${itemsPerPage}`,
        );

        setServices(response.data.data);
      } catch (error) {
        console.error("Erreur lors de la récupération des services", error);
      }
    };

    fetchServices().then((r) => r);
  }, [currentPage, itemsPerPage]);

  const filteredServices = services.filter(
    (service) =>
      service.provider_name.toLowerCase().includes(searchTerm.toLowerCase()) &&
      (selectedCategory === "all" ||
        (selectedCategory === "connected" && service.status === "connected") ||
        (selectedCategory === "available" &&
          service.status === "disconnected")),
  );

  const totalPages = Math.ceil(filteredServices.length / itemsPerPage);

  return (
    <DefaultLayout>
      <section className="flex flex-col items-center justify-center gap-8 py-8 md:py-10 px-4 max-w-7xl mx-auto">
        <div className="text-center mb-6">
          <h1 className={title({ color: "violet" })}>Gérez vos services</h1>
          <p className="text-large text-default-500 mt-2">
            Connectez et gérez vos services intégrés
          </p>
        </div>

        <div className="w-full flex flex-col sm:flex-row justify-between items-center gap-4 mb-6">
          <Input
            className="w-full sm:max-w-xs shadow-sm hover:shadow-md"
            placeholder="Rechercher des services..."
            startContent={<Search className="text-default-400" size={20} />}
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
          <Tabs
            aria-label="Catégories de services"
            className="flex"
            color="secondary"
            selectedKey={selectedCategory}
            onSelectionChange={(key) => setSelectedCategory(key as string)}
          >
            <Tab key="all" title="Tous" />
            <Tab key="connected" title="Connectés" />
            <Tab key="available" title="Disponibles" />
          </Tabs>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 w-full mt-4">
          {filteredServices.map((service) => (
            <ServiceCard key={service.id} service={service} />
          ))}
        </div>

        {filteredServices.length === 0 && (
          <p className="text-center text-default-500">
            Aucun service ne correspond à vos critères.
          </p>
        )}

        <div className="mt-4">
          <Pagination
            isCompact
            showControls
            initialPage={1}
            page={currentPage}
            total={totalPages}
            onChange={(page) => setCurrentPage(page)}
          />
        </div>
      </section>
    </DefaultLayout>
  );
};

export default ServicesPage;
