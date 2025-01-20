import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useToast } from "../hooks/use-toast";
import { Client } from '@/types/client';

const deliverySchema = z.object({
    name: z.string().min(3, 'Nome deve ter no mínimo 3 caracteres'),
    test: z.string().min(3, 'Teste técnico deve ter no mínimo 3 caracteres'),
    weight_kg: z.number().min(0.1, 'Peso deve ser maior que 0'),
    address: z.string().min(5, 'Endereço deve ter no mínimo 5 caracteres'),
});

interface DeliveryFormProps {
    onSubmit: (data: z.infer<typeof deliverySchema>) => Promise<void>;
    initialData?: Client;
}

export function DeliveryForm({ onSubmit, initialData }: DeliveryFormProps) {
    const { toast } = useToast();

    const form = useForm<z.infer<typeof deliverySchema>>({
        resolver: zodResolver(deliverySchema),
        defaultValues: initialData ? {
            name: initialData.name,
            test: initialData.test,
            weight_kg: initialData.weight_kg,
            address: initialData.address,
        } : undefined,
    });

    const handleSubmit = async (data: z.infer<typeof deliverySchema>) => {
        try {
            await onSubmit(data);
            toast({
                title: "Sucesso!",
                description: initialData
                    ? "Entrega atualizada com sucesso!"
                    : "Entrega cadastrada com sucesso!",
            });
            form.reset();
        } catch (error) {
            toast({
                title: "Erro",
                description: "Ocorreu um erro ao processar sua solicitação.",
                variant: "destructive",
            });
        }
    };

    return (
        <Card>
            <CardHeader>
                <CardTitle>
                    {initialData ? 'Editar entrega' : 'Novo registro'}
                </CardTitle>
            </CardHeader>
            <CardContent>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(handleSubmit)} className="space-y-4">

                        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 col-span-3 md:col-span-2">
                                <FormField
                                    control={form.control}
                                    name="name"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>Nome do cliente</FormLabel>
                                            <FormControl>
                                                <Input {...field} />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />

                                <FormField
                                    control={form.control}
                                    name="test"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>Teste técnico</FormLabel>
                                            <FormControl>
                                                <Input {...field} />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />
                            </div>

                            <div className="col-span-3 md:col-span-1">
                                <FormField
                                    control={form.control}
                                    name="weight_kg"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>Peso (kg)</FormLabel>
                                            <FormControl>
                                                <Input
                                                    type="number"
                                                    step="0.1"
                                                    {...field}
                                                    onChange={e => field.onChange(parseFloat(e.target.value))}

                                                />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />
                            </div>

                            <div className="col-span-3">
                                <FormField
                                    control={form.control}
                                    name="address"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>Endereço</FormLabel>
                                            <FormControl>
                                                <Input {...field} />
                                            </FormControl>
                                            <FormDescription>
                                                Digite o endereço completo para buscar as coordenadas
                                            </FormDescription>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />
                            </div>
                        </div>

                        <Button type="submit">
                            {(initialData ? 'Atualizar' : 'Cadastrar')}
                        </Button>
                    </form>
                </Form>
            </CardContent>
        </Card>
    );
}