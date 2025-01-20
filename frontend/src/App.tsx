import { QueryClient, QueryClientProvider } from 'react-query';
import { Toaster } from "@/components/ui/toaster";
import { MainLayout } from './components/layout/MainLayout';
import { DeliveryPage } from './pages/DeliveryPage';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: 1,
    },
  },
});

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <MainLayout>
        <DeliveryPage />
      </MainLayout>
      <Toaster />
    </QueryClientProvider>
  );
}

export default App;