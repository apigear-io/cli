import { NavLink as MantineNavLink, Stack } from '@mantine/core';
import { Link, useLocation } from 'react-router-dom';
import {
  IconDashboard,
  IconTemplate,
  IconFolder,
  IconCode,
  IconActivity,
  IconServer,
  IconChartLine,
  IconUsers,
  IconFileCode,
  IconFileText,
  IconEdit,
  IconPlayerPlay,
  IconList,
  IconSparkles,
} from '@tabler/icons-react';

interface NavigationProps {
  onNavigate?: () => void;
}

export function Navigation({ onNavigate }: NavigationProps) {
  const location = useLocation();

  const links = [
    { to: '/dashboard', label: 'Dashboard', icon: IconDashboard },
    { to: '/templates', label: 'Templates', icon: IconTemplate },
    { to: '/projects', label: 'Projects', icon: IconFolder },
    { to: '/codegen', label: 'Code Generation', icon: IconCode },
    { to: '/monitor', label: 'Monitor', icon: IconActivity },
  ];

  const streamLinks = [
    { to: '/stream/dashboard', label: 'Dashboard', icon: IconChartLine },
    { to: '/stream/proxies', label: 'Proxies', icon: IconServer },
    { to: '/stream/clients', label: 'Clients', icon: IconUsers },
    { to: '/stream/scripting', label: 'Scripting', icon: IconFileCode },
    { to: '/stream/traces', label: 'Traces', icon: IconFileText },
    { to: '/stream/editor', label: 'Stream Editor', icon: IconEdit },
    { to: '/stream/player', label: 'Stream Player', icon: IconPlayerPlay },
    { to: '/stream/generator', label: 'Generator', icon: IconSparkles },
    { to: '/stream/logs', label: 'Logs', icon: IconList },
  ];

  const isStreamActive = location.pathname.startsWith('/stream');

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

      <MantineNavLink
        label="Stream"
        leftSection={<IconServer size={20} />}
        active={isStreamActive}
        defaultOpened={isStreamActive}
      >
        {streamLinks.map((link) => {
          const Icon = link.icon;
          const isActive = location.pathname === link.to;

          return (
            <MantineNavLink
              key={link.to}
              component={Link}
              to={link.to}
              label={link.label}
              leftSection={<Icon size={16} />}
              active={isActive}
              onClick={onNavigate}
            />
          );
        })}
      </MantineNavLink>
    </Stack>
  );
}
