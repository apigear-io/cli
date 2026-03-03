import { Card, Group, Text, Badge, ActionIcon, Menu, Tooltip } from '@mantine/core';
import { IconFolder, IconTrash, IconFiles, IconDots, IconChevronRight } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';
import type { ProjectInfo } from '@/api/types';

interface ProjectCardProps {
  project: ProjectInfo;
  onDelete?: (path: string) => void;
}

export function ProjectCard({ project, onDelete }: ProjectCardProps) {
  const navigate = useNavigate();

  const handleCardClick = () => {
    const encodedPath = encodeURIComponent(project.path);
    navigate(`/codegen/projects/${encodedPath}`);
  };

  const handleMenuClick = (e: React.MouseEvent) => {
    // Prevent card click when clicking menu
    e.stopPropagation();
  };

  return (
    <Card
      shadow="sm"
      padding="lg"
      radius="md"
      withBorder
      style={{ cursor: 'pointer' }}
      onClick={handleCardClick}
    >
      <Group justify="space-between" mb="xs">
        <Group gap="xs">
          <IconFolder size={20} />
          <Text fw={500} size="lg">
            {project.name}
          </Text>
        </Group>

        <Group gap="xs">
          {onDelete && (
            <Menu shadow="md" width={200}>
              <Menu.Target>
                <ActionIcon variant="subtle" color="gray" onClick={handleMenuClick}>
                  <IconDots size={16} />
                </ActionIcon>
              </Menu.Target>

              <Menu.Dropdown onClick={handleMenuClick}>
                <Menu.Item
                  leftSection={<IconTrash size={16} />}
                  color="red"
                  onClick={() => onDelete(project.path)}
                >
                  Delete
                </Menu.Item>
              </Menu.Dropdown>
            </Menu>
          )}
          <IconChevronRight size={16} style={{ opacity: 0.5 }} />
        </Group>
      </Group>

      <Tooltip label={project.path} multiline w={300}>
        <Text size="sm" c="dimmed" truncate>
          {project.path}
        </Text>
      </Tooltip>

      <Group mt="md">
        <Badge leftSection={<IconFiles size={12} />} variant="light">
          {project.documents.length} {project.documents.length === 1 ? 'document' : 'documents'}
        </Badge>
      </Group>
    </Card>
  );
}
