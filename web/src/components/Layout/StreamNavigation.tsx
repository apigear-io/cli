import { NavLink as MantineNavLink, Stack } from '@mantine/core';
import { Link, useLocation } from 'react-router-dom';
import {
  IconChartLine,
  IconServer,
  IconUsers,
  IconFileCode,
  IconFileText,
  IconEdit,
  IconPlayerPlay,
  IconSparkles,
  IconList,
  IconSettings,
} from '@tabler/icons-react';

interface StreamNavigationProps {
  onNavigate?: () => void;
}

export function StreamNavigation({ onNavigate }: StreamNavigationProps) {
  const location = useLocation();

  const links = [
    { to: '/stream/dashboard', label: 'Dashboard', icon: IconChartLine },
    { to: '/stream/proxies', label: 'Proxies', icon: IconServer },
    { to: '/stream/clients', label: 'Clients', icon: IconUsers },
    { to: '/stream/scripting', label: 'Scripting', icon: IconFileCode },
    { to: '/stream/traces', label: 'Traces', icon: IconFileText },
    { to: '/stream/editor', label: 'Stream Editor', icon: IconEdit },
    { to: '/stream/player', label: 'Stream Player', icon: IconPlayerPlay },
    { to: '/stream/generator', label: 'Generator', icon: IconSparkles },
    { to: '/stream/logs', label: 'Logs', icon: IconList },
    { to: '/stream/settings', label: 'Settings', icon: IconSettings },
  ];

  return (
    <Stack gap="xs">
      {links.map((link) => {
        const Icon = link.icon;
        const isActive = location.pathname === link.to;

        return (
          <MantineNavLink
            key={link.to}
            component={Link}
            to={link.to}
            label={link.label}
            leftSection={<Icon size={20} />}
            active={isActive}
            onClick={onNavigate}
          />
        );
      })}
    </Stack>
  );
}
