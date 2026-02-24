import { Drawer, Tabs, Stack, Title, Table, List, Alert } from '@mantine/core';
import { CodeHighlight } from '@mantine/code-highlight';
import { IconInfoCircle } from '@tabler/icons-react';

interface HelpTab {
  value: string;
  label: string;
  content: React.ReactNode;
}

interface HelpDrawerProps {
  opened: boolean;
  onClose: () => void;
  title: string;
  tabs: HelpTab[];
}

export function HelpDrawer({ opened, onClose, title, tabs }: HelpDrawerProps) {
  return (
    <Drawer
      opened={opened}
      onClose={onClose}
      title={<Title order={3}>{title}</Title>}
      position="right"
      size="lg"
      padding="xl"
    >
      <Tabs defaultValue={tabs[0]?.value}>
        <Tabs.List>
          {tabs.map((tab) => (
            <Tabs.Tab key={tab.value} value={tab.value}>
              {tab.label}
            </Tabs.Tab>
          ))}
        </Tabs.List>

        {tabs.map((tab) => (
          <Tabs.Panel key={tab.value} value={tab.value} pt="md">
            <Stack gap="md">{tab.content}</Stack>
          </Tabs.Panel>
        ))}
      </Tabs>
    </Drawer>
  );
}

// Reusable help content components
export function HelpSection({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <Stack gap="xs">
      <Title order={4}>{title}</Title>
      {children}
    </Stack>
  );
}

export function HelpCode({ code, language = 'javascript' }: { code: string; language?: string }) {
  return (
    <CodeHighlight
      code={code}
      language={language}
      copyLabel="Copy code"
      copiedLabel="Copied!"
      withCopyButton
    />
  );
}

export function HelpTable({
  headers,
  rows,
}: {
  headers: string[];
  rows: Array<React.ReactNode[]>;
}) {
  return (
    <Table striped withTableBorder withColumnBorders>
      <Table.Thead>
        <Table.Tr>
          {headers.map((header, i) => (
            <Table.Th key={i}>{header}</Table.Th>
          ))}
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {rows.map((row, i) => (
          <Table.Tr key={i}>
            {row.map((cell, j) => (
              <Table.Td key={j}>{cell}</Table.Td>
            ))}
          </Table.Tr>
        ))}
      </Table.Tbody>
    </Table>
  );
}

export function HelpAlert({ children }: { children: React.ReactNode }) {
  return (
    <Alert icon={<IconInfoCircle size={16} />} color="blue" variant="light">
      {children}
    </Alert>
  );
}

export function HelpList({ items }: { items: React.ReactNode[] }) {
  return (
    <List>
      {items.map((item, i) => (
        <List.Item key={i}>{item}</List.Item>
      ))}
    </List>
  );
}
