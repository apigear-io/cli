grammar ObjectApi;

documentRule: headerRule declarationsRule*;

headerRule: moduleRule importRule*;

moduleRule: 'module' name = IDENTIFIER version = VERSION;
importRule: 'import' name = IDENTIFIER version = VERSION;

declarationsRule: interfaceRule | structRule | enumRule;

interfaceRule:
	'interface' name = IDENTIFIER extends = IDENTIFIER? '{' interfaceMembersRule* '}';

interfaceMembersRule: propertyRule | methodRule | signalRule;

propertyRule: name = IDENTIFIER ':' schema = schemaRule;
methodRule:
	name = IDENTIFIER '(' inputs = inputRule* ')' ':' (
		schema = schemaRule
	)?;
inputRule: name = IDENTIFIER ':' schema = schemaRule ','?;
signalRule:
	'signal' name = IDENTIFIER '(' inputs = inputRule* ')';

// structs
structRule: 'struct' name = IDENTIFIER '{' structFieldRule* '}';

structFieldRule: name = IDENTIFIER ':' schema = schemaRule;

// enums
enumRule: 'enum' name = IDENTIFIER '{' enumMemberRule* '}';

enumMemberRule: name = IDENTIFIER ('=' value = INTEGER)? ','?;

schemaRule: primitiveSchema | referenceSchema | arraySchema;

primitiveSchema:
	name = 'bool'
	| name = 'int'
	| name = 'float'
	| name = 'string';
referenceSchema: name = IDENTIFIER;
arraySchema: (primitiveSchema | referenceSchema) '[' ']';

WHITESPACE: [ \t\r\n]+ -> skip;
INTEGER: ('+' | '-')? '0' ..'9'+;
HEX: '0x' ('0' ..'9' | 'a' ..'f' | 'A' ..'F')+;
IDENTIFIER: [_A-Za-z] [_0-9A-Za-z]*;
VERSION: [0-9]'.' [0-9];