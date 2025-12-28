language json(cc);

namespace = "json"
includeGuardPrefix = "EXAMPLES_JSON_"
tokenLineOffset = true
tokenColumn = true
filenamePrefix = "json_"
optimizeTables = true
eventBased = true
eventAST = true
genSelector = true
parseParams = ["int x", "bool y"]
debugParser = true
scanBytes = true

:: lexer

%s initial, foo;

'{': /\{/
'}': /\}/
'[': /\[/
']': /\]/
':': /:/
',': /,/

<foo> Foo: /\#/

space: /[\t\r\n ]+/ (space)

commentChars = /([^*]|\*+[^*\/])*\**/
MultiLineComment: /\/\*{commentChars}\*\// (space)

hex = /[0-9a-fA-F]/

# TODO
JSONString: /"([^"\\]|\\(["\/\\bfnrt]|u{hex}{4}))*"/
#JSONString: /"([^"\\\x00-\x1f]|\\(["\/\\bfnrt]|u{hex}{4}))*"/

fraction = /\.[0-9]+/
exp = /[eE][+-]?[0-9]+/
JSONNumber: /-?(0|[1-9][0-9]*){fraction}?{exp}?/

id: /[a-zA-Z][a-zA-Z0-9]*/ (class)

kw_null: /null/
'true': /true/
'false': /false/

'A': /A/
'A': /α/
'B': /B/

'A': /A!/ { /*some code */ }

error:
invalid_token:

:: parser

%input JSONText;

%inject MultiLineComment -> MultiLineComment/Bar,Foo;
%inject invalid_token -> InvalidToken;
%inject JSONString -> JSONString;

%generate Literals = set(first JSONValue<+A>);

%flag A;

JSONText -> JSONText :
    JSONValue<+A> ;

JSONValue<A> -> JSONValue :
    kw_null
  | 'true'
  | 'false'
  | [A] 'A'
  | [!A] 'B'
  | JSONObject
  | JSONEmptyObject
  | JSONArray
  | JSONString
  | JSONNumber
;

JSONEmptyObject -> JSONEmptyObject : (?= JSONEmptyObject) '{' '}' { @$.begin = @1.begin; } ;

JSONObject -> JSONObject/Foo :
    (?= !JSONEmptyObject) '{' JSONMemberList? '}' { @$.begin = @1.begin; } ;

JSONMember -> JSONMember/Foo :
    JSONString ':'[b] { LOG(INFO) << @b.begin; } JSONValue<~A>
  | error -> SyntaxProblem
;

JSONMemberList:
    JSONMember
  | JSONMemberList .foo ',' JSONMember
;

JSONArray -> JSONArray/Foo :
    .bar '[' JSONElementListopt ']' ;

JSONElementList :
    JSONValue<+A>
  | JSONElementList ',' JSONValue<+A>
;
