# lapg syntax file

lang = "java"
prefix = "SAction"
package = "org.textway.lapg.parser.action"
positions = "offset"
endpositions = ""
lexerInput = "custom"


'{':	/\{/

_skip:	/'([^\n\\']|\\.)*'/
_skip:	/"([^\n\\"]|\\.)*"/
_skip:	/[^'"{}]+/

'}':	/\}/



%input javaaction no-eoi;

javaaction ::=
	'{' command_tokensopt '}' ;

command_tokens ::=
	command_tokens command_token | command_token ;

command_token ::=
	'{' command_tokensopt '}'
;
