import { SimpleGrid, Card, Stack, Text } from '@mantine/core';
import {
  IconChartLine,
  IconEdit,
  IconPlayerPlay,
  IconFileCode,
  IconServer,
  IconSettings,
} from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';

interface ActionCardProps {
  icon: React.ReactNode;
  title: string;
  description: string;
  path: string;
  color: string;
}

function ActionCard({ icon, title, description, path, color }: ActionCardProps) {
  const navigate = useNavigate();

  return (
    <Card
      shadow="sm"
      padding="lg"
      radius="md"
      withBorder
      style={{ cursor: 'pointer', transition: 'transform 0.2s' }}
      onClick={() => navigate(path)}
      onMouseEnter={(e) => {
        e.currentTarget.style.transform = 'translateY(-2px)';
      }}
      onMouseLeave={(e) => {
        e.currentTarget.style.transform = 'translateY(0)';
      }}
    >
      <Stack align="center" gap="md">
        <div
          style={{
            width: 60,
            height: 60,
            borderRadius: '50%',
            backgroundColor: `var(--mantine-color-${color}-1)`,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
          }}
        >
          {icon}
        </div>
        <Stack gap={4} align="center">
          <Text fw={600} size="sm">
            {title}
          </Text>
          <Text size="xs" c="dimmed" ta="center">
            {description}
          </Text>
        </Stack>
      </Stack>
    </Card>
  );
}

export function QuickActions() {
  const actions = [
    {
      icon: <IconChartLine size={28} color="var(--mantine-color-blue-6)" />,
      title: 'Proxy Stream',
      description: 'View live messages',
      path: '/stream/proxies',
      color: 'blue',
    },
    {
      icon: <IconEdit size={28} color="var(--mantine-color-green-6)" />,
      title: 'Stream Editor',
      description: 'Edit and filter traces',
      path: '/stream/editor',
      color: 'green',
    },
    {
      icon: <IconPlayerPlay size={28} color="var(--mantine-color-violet-6)" />,
      title: 'Stream Player',
      description: 'Replay trace files',
      path: '/stream/traces',
      color: 'violet',
    },
    {
      icon: <IconFileCode size={28} color="var(--mantine-color-orange-6)" />,
      title: 'Scripting',
      description: 'Run client scripts',
      path: '/stream/scripting',
      color: 'orange',
    },
    {
      icon: <IconServer size={28} color="var(--mantine-color-cyan-6)" />,
      title: 'Proxies',
      description: 'Manage proxy servers',
      path: '/stream/proxies',
      color: 'cyan',
    },
    {
      icon: <IconSettings size={28} color="var(--mantine-color-gray-6)" />,
      title: 'Settings',
      description: 'Configure options',
      path: '/stream/traces',
      color: 'gray',
    },
  ];

  return (
    <SimpleGrid cols={{ base: 2, sm: 3, md: 6 }} spacing="md">
      {actions.map((action) => (
        <ActionCard key={action.title} {...action} />
      ))}
    </SimpleGrid>
  );
}
