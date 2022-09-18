grammar ObjectApi;

documentRule: headerRule declarationsRule*;

headerRule: moduleRule importRule*;

moduleRule: 'module' name = IDENTIFIER version = VERSION;
importRule: 'import' name = IDENTIFIER version = VERSION;

declarationsRule: interfaceRule | structRule | enumRule;

interfaceRule:
	'interface' name = IDENTIFIER '{' interfaceMembersRule* '}';

interfaceMembersRule: propertyRule | operationRule | signalRule;

propertyRule: name = IDENTIFIER ':' schema = schemaRule;
operationRule:
	name = IDENTIFIER '(' params = operationParamRule* ')' operationReturnRule?;

operationReturnRule: ':' schema = schemaRule;
operationParamRule:
	name = IDENTIFIER ':' schema = schemaRule ','?;
signalRule:
	'signal' name = IDENTIFIER '(' params = operationParamRule* ')';

// structs
structRule: 'struct' name = IDENTIFIER '{' structFieldRule* '}';

structFieldRule: name = IDENTIFIER ':' schema = schemaRule;

// enums
enumRule: 'enum' name = IDENTIFIER '{' enumMemberRule* '}';

enumMemberRule: name = IDENTIFIER ('=' value = INTEGER)? ','?;

// a schame can be followed by "[]" to indicate an array
schemaRule: (primitiveSchema | symbolSchema) arrayRule?;

arrayRule: '[' ']';

primitiveSchema:
	name = 'bool'
	| name = 'int'
	| name = 'float'
	| name = 'string';

symbolSchema: name = IDENTIFIER;

WHITESPACE: [ \t\r\n]+ -> skip;
INTEGER: ('+' | '-')? '0' ..'9'+;
HEX: '0x' ('0' ..'9' | 'a' ..'f' | 'A' ..'F')+;
TYPE_IDENTIFIER: [A-Z_] [A-Z0-9_]* '[]'?;
IDENTIFIER: [_A-Za-z] [_0-9A-Za-z.]*;
VERSION: [0-9]'.' [0-9];