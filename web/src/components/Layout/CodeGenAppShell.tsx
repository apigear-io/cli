import { AppShell, Burger, Group, Title, Badge, SegmentedControl, Center } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import { IconCode, IconServer } from '@tabler/icons-react';
import { CodeGenNavigation } from './CodeGenNavigation';
import { useHealth } from '@/api/queries';

export function CodeGenAppShell() {
  const [opened, { toggle, close }] = useDisclosure();
  const { data: health, isLoading } = useHealth();
  const navigate = useNavigate();
  const location = useLocation();

  const healthStatus = health?.status === 'ok' ? 'success' : 'error';
  const healthColor = healthStatus === 'success' ? 'green' : 'red';

  // Determine current mode based on path
  const currentMode = location.pathname.startsWith('/stream') ? 'stream' : 'codegen';

  const handleModeChange = (value: string) => {
    close();
    if (value === 'stream') {
      navigate('/stream/dashboard');
    } else {
      navigate('/codegen/dashboard');
    }
  };

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

          <Group gap="sm">
            <SegmentedControl
              value={currentMode}
              onChange={handleModeChange}
              data={[
                {
                  value: 'codegen',
                  label: (
                    <Center style={{ gap: 8 }}>
                      <IconCode size={16} />
                      <span>CodeGen</span>
                    </Center>
                  ),
                },
                {
                  value: 'stream',
                  label: (
                    <Center style={{ gap: 8 }}>
                      <IconServer size={16} />
                      <span>Stream</span>
                    </Center>
                  ),
                },
              ]}
            />

            <Badge color={healthColor} variant="light">
              {isLoading ? 'Checking...' : health?.status || 'Unknown'}
            </Badge>
          </Group>
        </Group>
      </AppShell.Header>

      <AppShell.Navbar p="md">
        <CodeGenNavigation onNavigate={close} />
      </AppShell.Navbar>

      <AppShell.Main>
        <Outlet />
      </AppShell.Main>
    </AppShell>
  );
}
