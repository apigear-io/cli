import { Table, Text, Badge, Alert } from '@mantine/core';
import { IconInfoCircle } from '@tabler/icons-react';
import { useProxies } from '@/api/queries';

function formatBytes(bytes: number): string {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / Math.pow(k, i)).toFixed(2)} ${sizes[i] ?? 'TB'}`;
}

function formatRate(bytesPerSecond: number): string {
  if (bytesPerSecond === 0) return '0 B/s';
  const k = 1024;
  const sizes = ['B/s', 'KB/s', 'MB/s', 'GB/s'];
  const i = Math.floor(Math.log(bytesPerSecond) / Math.log(k));
  return `${(bytesPerSecond / Math.pow(k, i)).toFixed(2)} ${sizes[i] ?? 'TB/s'}`;
}

export function ProxyStatsTable() {
  const { data: proxies } = useProxies();

  if (proxies.length === 0) {
    return (
      <Alert icon={<IconInfoCircle />} color="blue">
        No proxies available. Create a proxy to see statistics here.
      </Alert>
    );
  }

  return (
    <Table striped highlightOnHover withTableBorder withColumnBorders>
      <Table.Thead>
        <Table.Tr>
          <Table.Th>PROXY</Table.Th>
          <Table.Th>CONNECTIONS</Table.Th>
          <Table.Th>MESSAGES IN</Table.Th>
          <Table.Th>MESSAGES OUT</Table.Th>
          <Table.Th>BYTES IN</Table.Th>
          <Table.Th>BYTES OUT</Table.Th>
          <Table.Th>RATE</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {proxies.map((proxy) => {
          const rate = (proxy.bytesReceived + proxy.bytesSent) / (proxy.uptime || 1);

          return (
            <Table.Tr key={proxy.name}>
              <Table.Td>
                <Text fw={500}>{proxy.name}</Text>
              </Table.Td>
              <Table.Td>
                <Badge color={proxy.activeConnections > 0 ? 'green' : 'gray'} size="sm">
                  {proxy.activeConnections}
                </Badge>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{proxy.messagesReceived.toLocaleString()}</Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{proxy.messagesSent.toLocaleString()}</Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{formatBytes(proxy.bytesReceived)}</Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{formatBytes(proxy.bytesSent)}</Text>
              </Table.Td>
              <Table.Td>
                <Text size="sm">{formatRate(rate)}</Text>
              </Table.Td>
            </Table.Tr>
          );
        })}
      </Table.Tbody>
    </Table>
  );
}
