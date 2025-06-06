grammar ObjectApi;

documentRule: headerRule declarationsRule*;

headerRule: moduleRule importRule*;

moduleRule:
	metaRule* 'module' name = IDENTIFIER version = VERSION? SEMICOLON?;

importRule:
	'import' name = IDENTIFIER version = VERSION? SEMICOLON?;

declarationsRule:
	externRule
	| interfaceRule
	| structRule
	| enumRule;

externRule: metaRule* 'extern' name = IDENTIFIER SEMICOLON?;

interfaceRule:
	metaRule* 'interface' name = IDENTIFIER (
		'extends' extends = IDENTIFIER
	)? '{' interfaceMembersRule* '}';

interfaceMembersRule: propertyRule | operationRule | signalRule;

propertyRule:
	metaRule* readonly = 'readonly'? name = IDENTIFIER ':' schema = schemaRule SEMICOLON?;
operationRule:
	metaRule* name = IDENTIFIER '(' params = operationParamRule* ')' operationReturnRule? SEMICOLON?
		;

operationReturnRule: ':' schema = schemaRule;
operationParamRule:
	name = IDENTIFIER ':' schema = schemaRule ','?;
signalRule:
	metaRule* 'signal' name = IDENTIFIER '(' params = operationParamRule* ')' SEMICOLON?;

// structs
structRule:
	metaRule* 'struct' name = IDENTIFIER '{' structFieldRule* '}';

structFieldRule:
	metaRule* readonly = 'readonly'? name = IDENTIFIER ':' schema = schemaRule SEMICOLON?;

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
	| name = 'string'
	| name = 'bytes'
	| name = 'any'
	| name = 'void';

symbolSchema: name = IDENTIFIER;

metaRule: tagLine = TAGLINE | docLine = DOCLINE;

WHITESPACE: ([ \t\r\n])+ -> skip;
INTEGER: ('+' | '-')? DIGIT+;
HEX: '0x' [a-fA-F0-9]+;
IDENTIFIER: LETTER ( DIGIT | LETTER | DOT)*;
VERSION: DIGIT+ DOT DIGIT+;
DOCLINE: '//' (~[\r\n])*;
TAGLINE: '@' (~[\r\n])*;
COMMENT: '#' (~[\r\n])* -> skip;
DOT: '.';
LETTER: [a-zA-Z_];
DIGIT: [0-9];
UNDERSCORE: '_';
SEMICOLON: ';';