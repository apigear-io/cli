import { Breadcrumbs as MantineBreadcrumbs, Anchor, Text } from '@mantine/core';
import { Link } from 'react-router-dom';
import { IconChevronRight } from '@tabler/icons-react';

export interface BreadcrumbItem {
  label: string;
  href?: string;
}

interface BreadcrumbsProps {
  items: BreadcrumbItem[];
}

export function Breadcrumbs({ items }: BreadcrumbsProps) {
  return (
    <MantineBreadcrumbs separator={<IconChevronRight size={14} />} mb="md">
      {items.map((item, index) => {
        const isLast = index === items.length - 1;

        if (isLast || !item.href) {
          return (
            <Text key={index} size="sm" c="dimmed">
              {item.label}
            </Text>
          );
        }

        return (
          <Anchor key={index} component={Link} to={item.href} size="sm">
            {item.label}
          </Anchor>
        );
      })}
    </MantineBreadcrumbs>
  );
}
