import {
  Stack,
  Text,
  Paper,
  Group,
  Badge,
  Button,
  ActionIcon,
  Tooltip,
  Title,
} from '@mantine/core';
import {
  IconFile,
  IconEdit,
  IconExternalLink,
  IconPlayerPlay,
} from '@tabler/icons-react';
import type { DocumentInfo } from '@/api/types';

const getDocumentTypeColor = (type: string | undefined) => {
  if (!type) return 'gray';
  switch (type.toLowerCase()) {
    case 'module': return 'blue';
    case 'solution': return 'green';
    case 'simulation': return 'orange';
    case 'scenario': return 'purple';
    default: return 'gray';
  }
};

interface DocumentSectionProps {
  title: string;
  icon: React.ReactNode;
  documents: DocumentInfo[];
  onEdit: (doc: DocumentInfo) => void;
  onOpenExternal: (doc: DocumentInfo) => void;
  onDocumentClick: (doc: DocumentInfo, e: React.MouseEvent) => void;
  onGenerate?: (doc: DocumentInfo) => void;
  showGenerateButton?: boolean;
  openingExternalPath: string | null;
}

export function DocumentSection({
  title,
  icon,
  documents,
  onEdit,
  onOpenExternal,
  onDocumentClick,
  onGenerate,
  showGenerateButton,
  openingExternalPath,
}: DocumentSectionProps) {
  if (documents.length === 0) return null;

  return (
    <Stack gap="xs">
      <Group gap="xs">
        {icon}
        <Title order={5}>{title}</Title>
        <Badge size="sm" variant="light">
          {documents.length}
        </Badge>
      </Group>
      {documents.map((doc, index) => (
        <Paper
          key={index}
          shadow="xs"
          p="md"
          withBorder
          style={{ cursor: 'pointer' }}
          onClick={(e) => onDocumentClick(doc, e)}
        >
          <Group justify="space-between" wrap="nowrap">
            <Group gap="md" style={{ flex: 1, minWidth: 0 }}>
              <IconFile size={20} style={{ flexShrink: 0 }} />
              <div style={{ flex: 1, minWidth: 0 }}>
                <Group gap="xs" mb={4}>
                  <Text fw={500}>{doc.name}</Text>
                  <Badge color={getDocumentTypeColor(doc.type)} size="sm">
                    {doc.type || 'unknown'}
                  </Badge>
                </Group>
                <Text
                  size="xs"
                  c="dimmed"
                  style={{ wordBreak: 'break-all', overflow: 'hidden' }}
                >
                  {doc.path}
                </Text>
              </div>
            </Group>

            <Group gap="xs" style={{ flexShrink: 0 }}>
              {showGenerateButton && onGenerate && (
                <Tooltip label="Generate code">
                  <Button
                    size="sm"
                    variant="light"
                    color="green"
                    leftSection={<IconPlayerPlay size={16} />}
                    onClick={() => onGenerate(doc)}
                  >
                    Generate
                  </Button>
                </Tooltip>
              )}

              <Tooltip label="Edit in browser">
                <ActionIcon
                  size="lg"
                  variant="light"
                  onClick={() => onEdit(doc)}
                >
                  <IconEdit size={18} />
                </ActionIcon>
              </Tooltip>

              <Tooltip label="Open in external editor">
                <ActionIcon
                  size="lg"
                  variant="light"
                  onClick={() => onOpenExternal(doc)}
                  loading={openingExternalPath === doc.path}
                >
                  <IconExternalLink size={18} />
                </ActionIcon>
              </Tooltip>
            </Group>
          </Group>
        </Paper>
      ))}
    </Stack>
  );
}
