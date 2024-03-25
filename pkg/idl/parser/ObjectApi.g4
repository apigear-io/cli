grammar ObjectApi;

documentRule: headerRule declarationsRule*;

headerRule: moduleRule importRule*;

moduleRule:
	metaRule* 'module' name = IDENTIFIER version = VERSION?;

importRule: 'import' name = IDENTIFIER version = VERSION?;

declarationsRule: interfaceRule | structRule | enumRule;

interfaceRule:
	metaRule* 'interface' name = IDENTIFIER (
		'extends' extends = IDENTIFIER
	)? '{' interfaceMembersRule* '}';

interfaceMembersRule: propertyRule | operationRule | signalRule;

propertyRule:
	metaRule* readonly = 'readonly'? name = IDENTIFIER ':' schema = schemaRule;
operationRule:
	metaRule* name = IDENTIFIER '(' params = operationParamRule* ')' operationReturnRule?;

operationReturnRule: ':' schema = schemaRule;
operationParamRule:
	name = IDENTIFIER ':' schema = schemaRule ','?;
signalRule:
	metaRule* 'signal' name = IDENTIFIER '(' params = operationParamRule* ')';

// structs
structRule:
	metaRule* 'struct' name = IDENTIFIER '{' structFieldRule* '}';

structFieldRule:
	metaRule* readonly = 'readonly'? name = IDENTIFIER ':' schema = schemaRule;

// enums
enumRule:
	metaRule* 'enum' name = IDENTIFIER '{' (enumMemberRule)* '}';

enumMemberRule:
	metaRule* name = IDENTIFIER ('=' value = INTEGER)? ','?;

// a schame can be followed by "[]" to indicate an array
schemaRule: (primitiveSchema | symbolSchema) (arrayRule)?;

arrayRule: '[' ']';

primitiveSchema:
	name = 'bool'
	| name = 'int'
	| name = 'int32'
	| name = 'int64'
	| name = 'float'
	| name = 'float32'
	| name = 'float64'
	| name = 'string';

symbolSchema: name = IDENTIFIER;

metaRule: tagLine = TAGLINE | docLine = DOCLINE;

WHITESPACE: ([ \t\r\n])+ -> skip;
INTEGER: ('+' | '-')? ('0' ..'9')+;
HEX: '0x' ('0' ..'9' | 'a' ..'f' | 'A' ..'F')+;
TYPE_IDENTIFIER: [A-Z_] ([A-Z0-9_])* ('[]')?;
IDENTIFIER: [_A-Za-z] ([_0-9A-Za-z.])*;
VERSION: [0-9]'.' [0-9];
DOCLINE: '//' (~[\r\n])*;
TAGLINE: '@' (~[\r\n])*;
COMMENT: '#' (~[\r\n])* -> skip;