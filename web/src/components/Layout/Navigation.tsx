import { NavLink as MantineNavLink, Stack } from '@mantine/core';
import { Link, useLocation } from 'react-router-dom';
import {
  IconDashboard,
  IconTemplate,
  IconFolder,
  IconCode,
  IconActivity,
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
