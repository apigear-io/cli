import { Suspense } from 'react';
import {
  Stack,
  Title,
  Group,
  Button,
  Text,
  Paper,
  Badge,
  SimpleGrid,
  Anchor,
} from '@mantine/core';
import {
  IconPlus,
  IconFolderOpen,
  IconTemplate,
  IconCircleCheck,
  IconAlertCircle,
} from '@tabler/icons-react';
import { Link, useNavigate } from 'react-router-dom';
import { useRecentProjects, useCachedTemplates, useHealth, useStatus } from '@/api/queries';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { ProjectCard } from '@/pages/Projects/components/ProjectCard';

function DashboardContent() {
  const navigate = useNavigate();
  const { data: projectsData } = useRecentProjects();
  const { data: templatesData } = useCachedTemplates();
  const { data: health } = useHealth();
  const { data: status } = useStatus();

  const recentProjects = projectsData.projects.slice(0, 5);
  const templates = templatesData.templates;
  const displayTemplates = templates.slice(0, 5);
  const updatesAvailable = templates.filter((t) => t.updateNeeded).length;

  return (
    <Stack gap="lg">
      <Title order={2}>Dashboard</Title>

      {/* Quick Actions */}
      <Group>
        <Button
          leftSection={<IconPlus size={16} />}
          onClick={() => navigate('/codegen/projects')}
        >
          Create Project
        </Button>
        <Button
          variant="default"
          leftSection={<IconFolderOpen size={16} />}
          onClick={() => navigate('/codegen/projects')}
        >
          Open Project
        </Button>
        <Button
          variant="default"
          leftSection={<IconTemplate size={16} />}
          onClick={() => navigate('/codegen/templates')}
        >
          Browse Templates
        </Button>
      </Group>

      {/* Recent Projects */}
      <Stack gap="xs">
        <Group justify="space-between">
          <Title order={4}>Recent Projects</Title>
          <Anchor component={Link} to="/codegen/projects" size="sm">
            View All
          </Anchor>
        </Group>
        {recentProjects.length > 0 ? (
          <SimpleGrid cols={{ base: 1, sm: 2, lg: 3 }}>
            {recentProjects.map((project) => (
              <ProjectCard key={project.path} project={project} />
            ))}
          </SimpleGrid>
        ) : (
          <Text size="sm" c="dimmed">
            No recent projects.{' '}
            <Anchor component={Link} to="/codegen/projects" size="sm">
              Create one
            </Anchor>{' '}
            to get started.
          </Text>
        )}
      </Stack>

      {/* Templates Overview */}
      <Stack gap="xs">
        <Group justify="space-between">
          <Group gap="sm">
            <Title order={4}>Templates</Title>
            <Badge variant="light" size="sm">
              {templates.length} installed
            </Badge>
            {updatesAvailable > 0 && (
              <Badge variant="light" color="yellow" size="sm">
                {updatesAvailable} update{updatesAvailable > 1 ? 's' : ''} available
              </Badge>
            )}
          </Group>
          <Anchor component={Link} to="/codegen/templates" size="sm">
            View All
          </Anchor>
        </Group>
        {displayTemplates.length > 0 ? (
          <Paper withBorder p="sm">
            <Stack gap="xs">
              {displayTemplates.map((template) => (
                <Group key={template.name} justify="space-between">
                  <Text size="sm">{template.name}</Text>
                  <Group gap="xs">
                    <Badge variant="outline" size="xs">
                      {template.version}
                    </Badge>
                    {template.updateNeeded && (
                      <Badge color="yellow" variant="light" size="xs">
                        Update available
                      </Badge>
                    )}
                  </Group>
                </Group>
              ))}
            </Stack>
          </Paper>
        ) : (
          <Text size="sm" c="dimmed">
            No templates installed.{' '}
            <Anchor component={Link} to="/codegen/templates" size="sm">
              Browse templates
            </Anchor>{' '}
            to install one.
          </Text>
        )}
      </Stack>

      {/* System Status */}
      <Paper withBorder p="sm">
        <Group gap="lg">
          <Group gap="xs">
            {health.status === 'ok' ? (
              <IconCircleCheck size={16} color="var(--mantine-color-green-6)" />
            ) : (
              <IconAlertCircle size={16} color="var(--mantine-color-red-6)" />
            )}
            <Badge color={health.status === 'ok' ? 'green' : 'red'} variant="light" size="sm">
              {health.status}
            </Badge>
          </Group>
          <Text size="sm" c="dimmed">
            Version: {status.version}
          </Text>
          <Text size="sm" c="dimmed">
            Uptime: {status.uptime}
          </Text>
          <Text size="sm" c="dimmed">
            Build: {status.buildDate}
          </Text>
        </Group>
      </Paper>
    </Stack>
  );
}

export function Dashboard() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading dashboard..." />}>
        <DashboardContent />
      </Suspense>
    </ErrorBoundary>
  );
}
