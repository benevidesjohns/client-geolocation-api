import { useEffect, useState } from 'react';
import { useToast } from "../hooks/use-toast";
import { clientService } from '../services/api';
import { DeliveryForm } from '../components/DeliveryForm';
import { DeliveryTable } from '../components/DeliveryTable';
import { SearchBar } from '../components/SearchBar';
import { Client } from '@/types/client';

export function DeliveryPage() {
    const [deliveries, setDeliveries] = useState<Client[]>([]);
    const [_, setEditingDelivery] = useState<Client | undefined>();
    const { toast } = useToast();

    const loadDeliveries = async () => {
        try {
            const data = await clientService.getAll();
            setDeliveries(data);
        } catch (error) {
            toast({
                title: "Erro",
                description: "Erro ao carregar entregas.",
                variant: "destructive",
            });
        }
    };

    useEffect(() => {
        loadDeliveries();
    }, []);

    const handleSearch = async (type: 'id' | 'city', value: string) => {
        try {
            if (type === 'id') {
                const delivery = await clientService.getById(value);
                setDeliveries(delivery ? [delivery] : []);
            } else {
                const deliveries = await clientService.getByCity(value);
                setDeliveries(deliveries);
            }
        } catch (error) {
            toast({
                title: "Erro",
                description: "Erro ao buscar entregas.",
                variant: "destructive",
            });
        }
    };

    const handleDelete = async (id: string) => {
        try {
            await clientService.delete(id);
            await loadDeliveries();
            toast({
                title: "Sucesso",
                description: "Entrega excluída com sucesso.",
            });
        } catch (error) {
            toast({
                title: "Erro",
                description: "Erro ao excluir entrega.",
                variant: "destructive"
            });
        }
    };

    return (
        <div className="container mx-auto py-6 space-y-6">
            <DeliveryForm
                onSubmit={async (data) => {
                    console.log(data)
                    try {
                        const clientData: Omit<Client, "id"> = {
                            name: data.name,
                            test: data.test,
                            weight_kg: data.weight_kg,
                            address: data.address,
                            number: '',
                            street: '',
                            neighborhood: '',
                            complement: '',
                            city: '',
                            state: '',
                            country: '',
                            latitude: 0,
                            longitude: 0
                        };

                        await clientService.create(clientData);
                        await loadDeliveries();
                    } catch (error) {
                        toast({
                            title: "Erro",
                            description: "Erro ao criar entrega",
                            variant: "destructive",
                        });
                    }
                }}
            />

            <SearchBar onSearch={handleSearch} />

            <DeliveryTable
                deliveries={deliveries}
                onEdit={setEditingDelivery}
                onDelete={handleDelete}
            />
        </div>
    );
}