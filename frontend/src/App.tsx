import { useQuery } from '@tanstack/react-query';
import { FC } from 'react';
import { listUsers } from './gen/proto/users/v1/users-UserService_connectquery';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { TransportProvider } from '@bufbuild/connect-query';
import { createConnectTransport } from '@bufbuild/connect-web';

const Example: FC = () => {
  const { data } = useQuery(listUsers.useQuery({}));
  return <div>{JSON.stringify(data)}</div>;
};

const queryClient = new QueryClient();

const App: FC = () => {
  const transport = createConnectTransport({
    baseUrl: 'http://0.0.0.0:8080',
  });
  return (
    <TransportProvider transport={transport}>
      <QueryClientProvider client={queryClient}>
        <Example />
      </QueryClientProvider>
    </TransportProvider>
  );
}

export default App;

