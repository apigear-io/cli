import { Stack, Paper, Group, Text, Button, Center, Loader, ActionIcon, Tooltip } from '@mantine/core';
import { modals } from '@mantine/modals';
import { notifications } from '@mantine/notifications';
import { IconMoodEmpty, IconCheck, IconAlertCircle, IconTrash, IconBrandGithub } from '@tabler/icons-react';
import { useRemoveTemplate } from '@/api/queries';
import type { TemplateInfo } from '@/api/types';

interface CachedTemplateListProps {
  templates: TemplateInfo[];
  isLoading: boolean;
}

export function CachedTemplateList({ templates, isLoading }: CachedTemplateListProps) {
  const removeMutation = useRemoveTemplate();

  const handleRemove = (template: TemplateInfo) => {
    modals.openConfirmModal({
      title: 'Remove Template',
      children: (
        <Text size="sm">
          Are you sure you want to remove <strong>{template.name}</strong>? This action cannot be
          undone.
        </Text>
      ),
      labels: { confirm: 'Remove', cancel: 'Cancel' },
      confirmProps: { color: 'red' },
      onConfirm: async () => {
        try {
          await removeMutation.mutateAsync(template.name);
          notifications.show({
            title: 'Success',
            message: `Template ${template.name} removed successfully`,
            color: 'green',
            icon: <IconCheck size={18} />,
          });
        } catch (error) {
          notifications.show({
            title: 'Error',
            message: error instanceof Error ? error.message : 'Failed to remove template',
            color: 'red',
            icon: <IconAlertCircle size={18} />,
          });
        }
      },
    });
  };

  if (isLoading) {
    return (
      <Center py="xl">
        <Stack align="center" gap="md">
          <Loader size="lg" />
          <Text c="dimmed">Loading installed templates...</Text>
        </Stack>
      </Center>
    );
  }

  if (templates.length === 0) {
    return (
      <Center py="xl">
        <Stack align="center" gap="md">
          <IconMoodEmpty size={48} stroke={1.5} opacity={0.5} />
          <Text c="dimmed">No templates installed</Text>
          <Text size="sm" c="dimmed">
            Install templates from the Registry tab to get started
          </Text>
        </Stack>
      </Center>
    );
  }

  return (
    <Stack gap="xs">
      {templates.map((template) => {
        const githubUrl = template.git ? template.git.replace(/\.git$/, '') : null;

        return (
          <Paper key={template.name} p="md" withBorder>
            <Group justify="space-between" wrap="nowrap">
              <Stack gap={4} style={{ flex: 1, minWidth: 0 }}>
                <Group gap="xs">
                  <Text fw={500} truncate>
                    {template.name}
                  </Text>
                  {githubUrl && (
                    <Tooltip label="View on GitHub">
                      <ActionIcon
                        component="a"
                        href={githubUrl}
                        target="_blank"
                        rel="noopener noreferrer"
                        variant="subtle"
                        color="gray"
                        size="sm"
                      >
                        <IconBrandGithub size={16} />
                      </ActionIcon>
                    </Tooltip>
                  )}
                </Group>
                <Group gap="xs">
                  <Text size="sm" c="dimmed">
                    v{template.version || 'unknown'}
                  </Text>
                  {template.description && (
                    <>
                      <Text size="sm" c="dimmed">
                        •
                      </Text>
                      <Text size="sm" c="dimmed" truncate>
                        {template.description}
                      </Text>
                    </>
                  )}
                </Group>
              </Stack>
              <Button
                color="red"
                variant="light"
                leftSection={<IconTrash size={16} />}
                onClick={() => handleRemove(template)}
                loading={removeMutation.isPending}
              >
                Remove
              </Button>
            </Group>
          </Paper>
        );
      })}
    </Stack>
  );
}
