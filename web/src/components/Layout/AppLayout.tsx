import { AppShell, Burger, Group, Title, Badge } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Outlet } from 'react-router-dom';
import { Navigation } from './Navigation';
import { useHealth } from '@/api/queries';

export function AppLayout() {
  const [opened, { toggle }] = useDisclosure();
  const { data: health, isLoading } = useHealth();

  const healthStatus = health?.status === 'ok' ? 'success' : 'error';
  const healthColor = healthStatus === 'success' ? 'green' : 'red';

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{
        width: 250,
        breakpoint: 'sm',
        collapsed: { mobile: !opened },
      }}
      padding="md"
    >
      <AppShell.Header>
        <Group h="100%" px="md" justify="space-between">
          <Group>
            <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
            <Title order={3}>ApiGear CLI</Title>
          </Group>
          <Badge color={healthColor} variant="light">
            {isLoading ? 'Checking...' : health?.status || 'Unknown'}
          </Badge>
        </Group>
      </AppShell.Header>

      <AppShell.Navbar p="md">
        <Navigation onNavigate={() => toggle()} />
      </AppShell.Navbar>

      <AppShell.Main>
        <Outlet />
      </AppShell.Main>
    </AppShell>
  );
}
