import { Client } from '@/types/client';
import axios from 'axios';

const api = axios.create({
  baseURL: "http://localhost:8080/api",
  headers: {
    'Content-Type': 'application/json',
  },
});

export const clientService = {
  create: async (client: Omit<Client, 'id'>) => {
    const response = await api.post<Client>('/deliveries', client);
    console.log(response.data)
    return response.data;
  },

  getAll: async () => {
    const response = await api.get<Client[]>('/deliveries');
    console.log(response.data)
    return response.data;
  },

  getById: async (id: string) => {
    const response = await api.get<Client>(`/deliveries?id=${id}`);
    console.log(response.data)
    return response.data;
  },

  getByCity: async (city: string) => {
    const response = await api.get<Client[]>(`/deliveries?city=${city}`);
    console.log(response.data)
    return response.data;
  },

  update: async (id: string, client: Omit<Client, 'id'>) => {
    const response = await api.put<Client>(`/deliveries/${id}`, client);
    console.log(response.data)
    return response.data;
  },

  delete: async (id: string) => {
    await api.delete(`/deliveries/${id}`);
  },

  deleteAll: async () => {
    await api.delete('/deliveries');
  }
};