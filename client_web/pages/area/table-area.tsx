import React, {useState, useEffect} from "react";
import {
    Button,
    Input,
    Table,
    TableHeader,
    TableCell,
    TableColumn,
    TableBody,
    TableRow,
    Pagination,
    Dropdown,
    DropdownTrigger,
    DropdownMenu,
    DropdownItem,
} from "@nextui-org/react";
import axios from "axios";
import Cookies from "js-cookie";

interface Area {
    id: number;
    provider_name: string;
    provider_slug: string;
    action?: string;
    reaction?: string;
    status?: string;
}

interface Action {
    id: number;
    name: string;
    description: string;
    method: string;
    type: string;
}

const AreaTable = () => {
    const API_URL = process.env.NEXT_PUBLIC_API_URL;
    const token = Cookies.get("token");
    const [areas, setAreas] = useState<Area[]>([]);
    const [actions, setActions] = useState<Action[]>([]);
    const [search, setSearch] = useState("");
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage = 4;

    useEffect(() => {
        const fetchAreas = async () => {
            try {
                const response = await axios.get(`${API_URL}/v1/providers/?page=${currentPage}&pageSize=${itemsPerPage}`);
                setAreas(response.data.data);
            } catch (error) {
                console.error("Erreur lors de la récupération des services", error);
            }
        };

        fetchAreas().then(r => r);
    }, [API_URL, token, currentPage, itemsPerPage]);

    useEffect(() => {
        const fetchActions = async () => {
            try {
                const response = await axios.get(`${API_URL}/v1/action/?page=${currentPage}&pageSize=${itemsPerPage}`);
                setActions(response.data.data);
            } catch (error) {
                console.error("Erreur lors de la récupération des actions", error);
            }
        };

        fetchActions().then(r => r);
    }, [API_URL, token, currentPage, itemsPerPage]);

    const filteredAreas = areas.filter((area) =>
      area.provider_name.toLowerCase().includes(search.toLowerCase())
    );

    const totalPages = Math.ceil(filteredAreas.length / itemsPerPage);
    const startIndex = (currentPage - 1) * itemsPerPage;
    const paginatedAreas = filteredAreas.slice(startIndex, startIndex + itemsPerPage);

    const Actions = actions.filter(action => action.type === "TRIGGER");
    const Reaction = actions.filter(action => action.type !== "TRIGGER");

    return (
      <div className="p-5 text-center">
          <Input
            isClearable
            className="mt-5 w-72 mb-5"
            placeholder="Rechercher des AREAs..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
          <Table
            aria-label="Liste des AREAs"
            className="h-auto min-w-full bg-gray-900"
          >
              <TableHeader>
                  <TableColumn>Service</TableColumn>
                  <TableColumn>Action</TableColumn>
                  <TableColumn>Réaction</TableColumn>
                  <TableColumn>Status</TableColumn>
                  <TableColumn>Actions</TableColumn>
              </TableHeader>
              <TableBody>
                  {paginatedAreas.map((area) => (
                    <TableRow key={area.id}>
                        <TableCell className="text-md font-semibold text-indigo-600">
                            {area.provider_name}
                        </TableCell>
                        <TableCell className="text-md font-semibold text-indigo-600">
                            <Dropdown>
                                <DropdownTrigger>
                                    <Button>{area.action ? area.action : "Select Action"}</Button>
                                </DropdownTrigger>
                                <DropdownMenu
                                  onAction={(key) => {
                                      const updatedAreas = areas.map(a =>
                                        a.id === area.id ? {...a, action: key.toString()} : a
                                      );
                                      setAreas(updatedAreas);
                                  }}
                                >
                                    {Actions.map(action => (
                                      <DropdownItem key={action.name}>{action.name}</DropdownItem>
                                    ))}
                                </DropdownMenu>
                            </Dropdown>
                        </TableCell>
                        <TableCell className="text-md font-semibold text-indigo-600">
                            <Dropdown>
                                <DropdownTrigger>
                                    <Button>{area.reaction ? area.reaction : "Select Reaction"}</Button>
                                </DropdownTrigger>
                                <DropdownMenu
                                  onAction={(key) => {
                                      const updatedAreas = areas.map(a =>
                                        a.id === area.id ? {...a, reaction: key.toString()} : a
                                      );
                                      setAreas(updatedAreas);
                                  }}
                                >
                                    {Reaction.map(action => (
                                      <DropdownItem key={action.name}>{action.name}</DropdownItem>
                                    ))}
                                </DropdownMenu>
                            </Dropdown>
                        </TableCell>
                        <TableCell className="text-md font-semibold text-indigo-600">
                            {area.status ? area.status : "inactive"}
                        </TableCell>
                        <TableCell>
                            <Button
                              color={area.status === "Active" ? "primary" : "warning"}
                              className="flex items-center justify-center w-24"
                            >
                                {area.status === "Active" ? "Désactiver" : "Activer"}
                            </Button>
                        </TableCell>
                    </TableRow>
                  ))}
              </TableBody>
          </Table>
          <div className="mt-4">
              <Pagination
                total={totalPages}
                initialPage={1}
                showControls
                isCompact
                onChange={(page) => setCurrentPage(page)}
                page={currentPage}
              />
          </div>
      </div>
    );
};

export default AreaTable;