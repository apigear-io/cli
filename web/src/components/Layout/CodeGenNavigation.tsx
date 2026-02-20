import { NavLink as MantineNavLink, Stack } from '@mantine/core';
import { Link, useLocation } from 'react-router-dom';
import {
  IconDashboard,
  IconTemplate,
  IconFolder,
  IconCode,
  IconActivity,
} from '@tabler/icons-react';

interface CodeGenNavigationProps {
  onNavigate?: () => void;
}

export function CodeGenNavigation({ onNavigate }: CodeGenNavigationProps) {
  const location = useLocation();

  const links = [
    { to: '/codegen/dashboard', label: 'Dashboard', icon: IconDashboard },
    { to: '/codegen/templates', label: 'Templates', icon: IconTemplate },
    { to: '/codegen/projects', label: 'Projects', icon: IconFolder },
    { to: '/codegen/generate', label: 'Code Generation', icon: IconCode },
    { to: '/codegen/monitor', label: 'Monitor', icon: IconActivity },
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
