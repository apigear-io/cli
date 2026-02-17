import { useState } from 'react';
import { Card, Stack, Group, Text, Badge, Button, Progress } from '@mantine/core';
import { notifications } from '@mantine/notifications';
import { IconCheck, IconAlertCircle } from '@tabler/icons-react';
import { useInstallTemplate } from '@/api/queries';
import type { TemplateInfo, InstallProgressEvent } from '@/api/types';

interface TemplateCardProps {
  template: TemplateInfo;
}

export function TemplateCard({ template }: TemplateCardProps) {
  const [installing, setInstalling] = useState(false);
  const [progress, setProgress] = useState(0);
  const [progressMessage, setProgressMessage] = useState('');
  const installMutation = useInstallTemplate();

  const handleInstall = async () => {
    setInstalling(true);
    setProgress(0);

    try {
      await installMutation.mutateAsync({
        id: template.name,
        onProgress: (event: InstallProgressEvent) => {
          setProgress(event.progress);
          setProgressMessage(event.message);
        },
      });

      notifications.show({
        title: 'Success',
        message: `Template ${template.name} installed successfully`,
        color: 'green',
        icon: <IconCheck size={18} />,
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Installation failed',
        color: 'red',
        icon: <IconAlertCircle size={18} />,
      });
    } finally {
      setInstalling(false);
      setProgress(0);
      setProgressMessage('');
    }
  };

  // Check if template is up to date
  // Version can be empty for newly installed templates, treat as up to date if in cache
  const hasVersion = template.version && template.version.trim() !== '';
  const hasLatest = template.latest && template.latest.trim() !== '';
  const isUpToDate = template.inCache && (!hasLatest || !hasVersion || template.version === template.latest);
  const hasUpdate = template.inCache && hasVersion && hasLatest && template.version !== template.latest;

  return (
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Stack gap="md">
        <Group justify="space-between" align="flex-start">
          <div style={{ flex: 1, minWidth: 0 }}>
            <Text fw={600} size="sm" truncate>
              {template.name}
            </Text>
          </div>
          <Group gap="xs">
            {template.inCache && (
              <Badge color="green" size="sm" variant="light">
                Installed
              </Badge>
            )}
            {hasUpdate && (
              <Badge color="orange" size="sm" variant="light">
                Update Available
              </Badge>
            )}
          </Group>
        </Group>

        <Text size="sm" c="dimmed" lineClamp={2} style={{ minHeight: '2.5rem' }}>
          {template.description || 'No description available'}
        </Text>

        <Group gap="xs" wrap="nowrap">
          <Text size="xs" c="dimmed" truncate>
            Latest: {template.latest || 'N/A'}
          </Text>
          {template.inCache && template.version && (
            <Text size="xs" c="dimmed" truncate>
              • Installed: {template.version}
            </Text>
          )}
        </Group>

        {installing ? (
          <Stack gap="xs">
            <Progress value={progress} animated />
            <Text size="xs" c="dimmed" ta="center">
              {progressMessage}
            </Text>
          </Stack>
        ) : (
          <Button
            fullWidth
            onClick={handleInstall}
            disabled={isUpToDate}
            variant={hasUpdate ? 'filled' : 'light'}
            loading={installMutation.isPending}
          >
            {isUpToDate ? 'Up to Date' : template.inCache ? 'Update' : 'Install'}
          </Button>
        )}
      </Stack>
    </Card>
  );
}
