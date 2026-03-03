import Editor from '@monaco-editor/react';

interface ScriptEditorProps {
  code: string;
  onChange: (value: string) => void;
  height?: string;
}

export function ScriptEditor({ code, onChange, height = '400px' }: ScriptEditorProps) {
  return (
    <Editor
      height={height}
      defaultLanguage="javascript"
      value={code}
      onChange={(value) => onChange(value || '')}
      theme="vs-dark"
      options={{
        minimap: { enabled: true },
        fontSize: 14,
        automaticLayout: true,
        tabSize: 2,
        lineNumbers: 'on',
        scrollBeyondLastLine: false,
        wordWrap: 'on',
      }}
    />
  );
}
