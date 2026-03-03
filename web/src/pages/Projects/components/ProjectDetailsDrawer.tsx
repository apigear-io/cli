import { Drawer, Stack, Text, Title, Badge, Group, Paper, ScrollArea } from '@mantine/core';
import { IconFile, IconFolder } from '@tabler/icons-react';
import type { ProjectInfo } from '@/api/types';

interface ProjectDetailsDrawerProps {
  project: ProjectInfo | null;
  onClose: () => void;
}

const getDocumentTypeColor = (type: string | undefined) => {
  if (!type) {
    return 'gray';
  }

  switch (type.toLowerCase()) {
    case 'module':
      return 'blue';
    case 'solution':
      return 'green';
    case 'simulation':
      return 'orange';
    case 'scenario':
      return 'purple';
    case 'unknown':
      return 'gray';
    default:
      return 'gray';
  }
};

export function ProjectDetailsDrawer({ project, onClose }: ProjectDetailsDrawerProps) {
  if (!project) {
    return null;
  }

  return (
    <Drawer
      opened={!!project}
      onClose={onClose}
      title="Project Details"
      position="right"
      size="lg"
    >
      <ScrollArea style={{ height: 'calc(100vh - 80px)' }}>
        <Stack gap="lg">
          {/* Project Info */}
          <Paper shadow="sm" p="md" withBorder>
            <Stack gap="xs">
              <Group gap="xs">
                <IconFolder size={20} />
                <Title order={4}>{project.name}</Title>
              </Group>
              <Text size="sm" c="dimmed">
                {project.path}
              </Text>
            </Stack>
          </Paper>

          {/* Documents List */}
          <Stack gap="sm">
            <Title order={5}>Documents ({project.documents.length})</Title>

            {project.documents.length === 0 ? (
              <Text c="dimmed" size="sm">
                No documents found in this project
              </Text>
            ) : (
              project.documents.map((doc, index) => (
                <Paper key={index} shadow="xs" p="sm" withBorder>
                  <Group justify="space-between">
                    <Group gap="xs">
                      <IconFile size={16} />
                      <Text size="sm" fw={500}>
                        {doc.name}
                      </Text>
                    </Group>
                    <Badge color={getDocumentTypeColor(doc.type)} size="sm">
                      {doc.type || 'unknown'}
                    </Badge>
                  </Group>
                  <Text size="xs" c="dimmed" mt="xs" style={{ wordBreak: 'break-all' }}>
                    {doc.path}
                  </Text>
                </Paper>
              ))
            )}
          </Stack>
        </Stack>
      </ScrollArea>
    </Drawer>
  );
}
