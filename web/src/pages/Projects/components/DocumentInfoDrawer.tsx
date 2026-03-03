import { Drawer, Stack, Text, Group, Badge, Button, Divider } from '@mantine/core';
import {
  IconFile,
  IconEdit,
  IconExternalLink,
  IconPlayerPlay,
} from '@tabler/icons-react';
import type { DocumentInfo } from '@/api/types';

interface DocumentInfoDrawerProps {
  document: DocumentInfo | null;
  onClose: () => void;
  onEdit: (doc: DocumentInfo) => void;
  onOpenExternal: (doc: DocumentInfo) => void;
  onGenerate?: (doc: DocumentInfo) => void;
  isOpeningExternal: boolean;
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

const isSolutionFile = (doc: DocumentInfo) => {
  return doc.name.endsWith('.solution.yaml');
};

export function DocumentInfoDrawer({
  document,
  onClose,
  onEdit,
  onOpenExternal,
  onGenerate,
  isOpeningExternal,
}: DocumentInfoDrawerProps) {
  if (!document) {
    return null;
  }

  return (
    <Drawer
      opened={!!document}
      onClose={onClose}
      title="Document Details"
      position="right"
      size="lg"
    >
      <Stack gap="lg">
        {/* Document Header */}
        <Group gap="xs">
          <IconFile size={24} />
          <div style={{ flex: 1 }}>
            <Text fw={500} size="lg">
              {document.name}
            </Text>
            <Badge color={getDocumentTypeColor(document.type)} size="sm" mt={4}>
              {document.type || 'unknown'}
            </Badge>
          </div>
        </Group>

        <Divider />

        {/* Document Info */}
        <Stack gap="xs">
          <Text size="sm" fw={500}>
            File Path
          </Text>
          <Text size="sm" c="dimmed" style={{ wordBreak: 'break-all' }}>
            {document.path}
          </Text>
        </Stack>

        <Divider />

        {/* Actions */}
        <Stack gap="md">
          <Text size="sm" fw={500}>
            Actions
          </Text>

          {isSolutionFile(document) && onGenerate && (
            <Button
              fullWidth
              variant="light"
              color="green"
              leftSection={<IconPlayerPlay size={18} />}
              onClick={() => onGenerate(document)}
            >
              Generate Code
            </Button>
          )}

          <Button
            fullWidth
            variant="light"
            leftSection={<IconEdit size={18} />}
            onClick={() => onEdit(document)}
          >
            Edit in Browser
          </Button>

          <Button
            fullWidth
            variant="light"
            leftSection={<IconExternalLink size={18} />}
            onClick={() => onOpenExternal(document)}
            loading={isOpeningExternal}
          >
            Open in External Editor
          </Button>
        </Stack>
      </Stack>
    </Drawer>
  );
}
